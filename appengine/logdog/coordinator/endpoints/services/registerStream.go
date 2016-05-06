// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package services

import (
	"errors"
	"time"

	ds "github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/appengine/logdog/coordinator"
	"github.com/luci/luci-go/appengine/logdog/coordinator/hierarchy"
	"github.com/luci/luci-go/appengine/logdog/coordinator/mutations"
	"github.com/luci/luci-go/appengine/tumble"
	"github.com/luci/luci-go/common/api/logdog_coordinator/services/v1"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/config"
	"github.com/luci/luci-go/common/grpcutil"
	"github.com/luci/luci-go/common/logdog/types"
	log "github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/proto/logdog/logpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func loadLogStreamState(project config.ProjectName, ls *coordinator.LogStream) *logdog.LogStreamState {
	st := logdog.LogStreamState{
		Project:       string(project),
		Path:          string(ls.Path()),
		ProtoVersion:  ls.ProtoVersion,
		TerminalIndex: ls.TerminalIndex,
		Archived:      ls.Archived(),
		Purged:        ls.Purged,
	}
	if !ls.Terminated() {
		st.TerminalIndex = -1
	}
	return &st
}

// RegisterStream is an idempotent stream state register operation.
func (s *server) RegisterStream(c context.Context, req *logdog.RegisterStreamRequest) (*logdog.RegisterStreamResponse, error) {
	log.Fields{
		"project": req.Project,
		"path":    req.Path,
	}.Infof(c, "Registration request for log stream.")

	path := types.StreamPath(req.Path)
	if err := path.Validate(); err != nil {
		return nil, grpcutil.Errf(codes.InvalidArgument, "Invalid path (%s): %s", path, err)
	}

	switch {
	case req.ProtoVersion == "":
		return nil, grpcutil.Errf(codes.InvalidArgument, "No protobuf version supplied.")
	case req.ProtoVersion != logpb.Version:
		return nil, grpcutil.Errf(codes.InvalidArgument, "Unrecognized protobuf version.")
	case req.Desc == nil:
		return nil, grpcutil.Errf(codes.InvalidArgument, "Missing log stream descriptor.")
	}

	prefix, name := path.Split()
	if err := req.Desc.Validate(true); err != nil {
		return nil, grpcutil.Errf(codes.InvalidArgument, "Invalid log stream descriptor: %s", err)
	}
	switch {
	case req.Desc.Prefix != string(prefix):
		return nil, grpcutil.Errf(codes.InvalidArgument, "Descriptor prefix does not match path (%s != %s)",
			req.Desc.Prefix, prefix)
	case req.Desc.Name != string(name):
		return nil, grpcutil.Errf(codes.InvalidArgument, "Descriptor name does not match path (%s != %s)",
			req.Desc.Name, name)
	}

	// Load our config and archive expiration.
	_, cfg, err := coordinator.GetServices(c).Config(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load configuration.")
		return nil, grpcutil.Internal
	}

	archiveDelayMax := cfg.Coordinator.ArchiveDelayMax.Duration()
	if archiveDelayMax < 0 {
		log.Fields{
			"archiveDelayMax": archiveDelayMax,
		}.Errorf(c, "Must have positive maximum archive delay.")
		return nil, grpcutil.Internal
	}

	// Register our Prefix.
	//
	// This will also verify that our request secret matches the registered one,
	// if one is registered.
	//
	// Note: This step will not be necessary once a "register prefix" RPC call is
	// implemented.
	lsp := logStreamPrefix{
		prefix: string(prefix),
		secret: req.Secret,
	}
	pfx, err := registerPrefix(c, &lsp)
	if err != nil {
		log.Errorf(c, "Failed to register/validate log stream prefix.")
		return nil, err
	}
	log.Fields{
		"prefix":        pfx.Prefix,
		"prefixCreated": pfx.Created,
	}.Debugf(c, "Loaded log stream prefix.")

	ls := coordinator.LogStreamFromPath(path)

	// Check for registration (non-transactional).
	di := ds.Get(c)
	switch err := checkRegisterStream(di, req, ls); err {
	case nil:
		// The stream is already compatibly registered.
		break

	case ds.ErrNoSuchEntity:
		// The stream is not registered. Perform a transactional registration via
		// mutation.
		//
		// Determine which hierarchy components we need to add.
		comps := hierarchy.Components(path)
		if comps, err = hierarchy.Missing(di, comps); err != nil {
			log.WithError(err).Warningf(c, "Failed to probe for missing hierarchy components.")
		}

		// Before we go into transaction, try and put these entries. This can
		// encounter datastore contention, so we'll schedule a Tumble mutation if
		// this doesn't work.
		if err := hierarchy.PutMulti(di, comps); err != nil {
			log.WithError(err).Infof(c, "Failed to add missing hierarchy components.")
			return nil, grpcutil.Internal
		}

		// The stream does not exist. Proceed with transactional registration.
		err = tumble.RunMutation(c, &registerStreamMutation{
			LogStream:    ls,
			req:          req,
			pfx:          pfx,
			archiveDelay: archiveDelayMax,
		})
		if err != nil {
			log.Fields{
				log.ErrorKey: err,
			}.Errorf(c, "Failed to register LogStream.")
			return nil, err
		}

	default:
		log.WithError(err).Errorf(c, "Failed to check for log stream.")
		return nil, err
	}

	return &logdog.RegisterStreamResponse{
		State:  loadLogStreamState(coordinator.Project(c), ls),
		Secret: lsp.secret,
	}, nil
}

func checkRegisterStream(di ds.Interface, req *logdog.RegisterStreamRequest, ls *coordinator.LogStream) error {
	// Load any existing entity.
	if err := di.Get(ls); err != nil {
		return err
	}

	// A log stream is already present. We will either error if it is incompatible
	// with our request, or return nil if it compatible (idempotent).
	if err := matchesLogStream(req, ls); err != nil {
		return grpcutil.Errf(codes.AlreadyExists, "Log stream is already registered: %v", err)
	}
	return nil
}

func matchesLogStream(r *logdog.RegisterStreamRequest, ls *coordinator.LogStream) error {
	if r.Path != string(ls.Path()) {
		return errors.New("paths do not match")
	}

	if r.ProtoVersion != ls.ProtoVersion {
		return errors.New("protobuf version does not match")
	}

	dv, err := ls.DescriptorValue()
	if err != nil {
		return errors.New("log stream has invalid descriptor value")
	}
	if !dv.Equal(r.Desc) {
		return errors.New("descriptor protobufs do not match")
	}

	return nil
}

type registerStreamMutation struct {
	*coordinator.LogStream

	req          *logdog.RegisterStreamRequest
	pfx          *coordinator.LogPrefix
	archiveDelay time.Duration
}

func (m *registerStreamMutation) RollForward(c context.Context) ([]tumble.Mutation, error) {
	di := ds.Get(c)

	// Check if our stream is registered (transactional).
	switch err := checkRegisterStream(di, m.req, m.LogStream); err {
	case nil:
		// The stream is compatibly registered, so this is idempotent.
		return nil, nil

	case ds.ErrNoSuchEntity:
		// The stream is not registered. We will proceed to register.
		break

	default:
		log.WithError(err).Errorf(c, "Failed to check for stream registration (transactional).")
		return nil, err
	}

	log.Infof(c, "Registering new log stream'")

	// The stream is not yet registered.
	if err := m.LoadDescriptor(m.req.Desc); err != nil {
		log.Fields{
			log.ErrorKey: err,
		}.Errorf(c, "Failed to load descriptor into LogStream.")
		return nil, grpcutil.Errf(codes.InvalidArgument, "Failed to load descriptor.")
	}

	now := clock.Now(c).UTC()

	m.Secret = m.pfx.Secret // Copy Prefix Secret to reduce number of Get needed.
	m.ProtoVersion = m.req.ProtoVersion
	m.State = coordinator.LSStreaming
	m.Created = now
	m.TerminalIndex = -1

	if err := di.Put(m.LogStream); err != nil {
		log.Fields{
			log.ErrorKey: err,
		}.Errorf(c, "Failed to Put() LogStream.")
		return nil, grpcutil.Internal
	}

	// Add a named delayed mutation to archive this stream if it's not archived
	// yet. We will cancel this in terminateStream once we dispatch an immediate
	// archival task.
	archiveExpiredMutation := mutations.CreateArchiveTask{
		Path:       m.Path(),
		Expiration: now.Add(m.archiveDelay),
	}
	log.Fields{
		"archiveDelay": m.archiveDelay,
		"deadline":     archiveExpiredMutation.Expiration,
	}.Infof(c, "Scheduling expiration deadline mutation.")
	aeParent, aeName := archiveExpiredMutation.TaskName(di)
	err := tumble.PutNamedMutations(c, aeParent, map[string]tumble.Mutation{
		aeName: &archiveExpiredMutation,
	})
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load named mutations.")
		return nil, grpcutil.Internal
	}

	return nil, nil
}

func (m *registerStreamMutation) Root(c context.Context) *ds.Key {
	return ds.Get(c).KeyForObj(m.LogStream)
}
