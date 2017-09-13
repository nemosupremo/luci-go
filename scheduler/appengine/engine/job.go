// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"bytes"
	"context"
	"hash/fnv"
	"strings"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/scheduler/appengine/acl"
	"go.chromium.org/luci/scheduler/appengine/catalog"
	"go.chromium.org/luci/scheduler/appengine/schedule"
)

// Job stores the last known definition of a scheduler job, as well as its
// current state. Root entity, its kind is "Job".
type Job struct {
	_kind  string                `gae:"$kind,Job"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// cachedSchedule and cachedScheduleErr are used by ParseSchedule().
	cachedSchedule    *schedule.Schedule `gae:"-"`
	cachedScheduleErr error              `gae:"-"`

	// JobID is '<ProjectID>/<JobName>' string. JobName is unique with a project,
	// but not globally. JobID is unique globally.
	JobID string `gae:"$id"`

	// ProjectID exists for indexing. It matches <projectID> portion of JobID.
	ProjectID string

	// Flavor describes what category of jobs this is, see the enum.
	Flavor catalog.JobFlavor `gae:",noindex"`

	// Enabled is false if the job was disabled or removed from config.
	//
	// Disabled jobs do not show up in UI at all (they are still kept in the
	// datastore though, for audit purposes).
	Enabled bool

	// Paused is true if job's schedule is ignored and job can only be started
	// manually via "Run now" button.
	Paused bool `gae:",noindex"`

	// Revision is last seen job definition revision.
	Revision string `gae:",noindex"`

	// RevisionURL is URL to human readable page with config file at
	// an appropriate revision.
	RevisionURL string `gae:",noindex"`

	// Schedule is the job's schedule in regular cron expression format.
	Schedule string `gae:",noindex"`

	// Task is the job's payload in serialized form. Opaque from the point of view
	// of the engine. See Catalog.UnmarshalTask().
	Task []byte `gae:",noindex"`

	// TriggeredJobIDs is a list of jobIDs of jobs which this job triggers.
	// The list is sorted and without duplicates.
	TriggeredJobIDs []string `gae:",noindex"`

	// ACLs are the latest ACLs applied to Job and all its invocations.
	Acls acl.GrantsByRole `gae:",noindex"`

	// State is the job's state machine state, see StateMachine.
	State JobState
}

// JobName returns name of this Job as defined its project's config.
//
// This is "<name>"" part extracted from "<project>/<name>" job ID.
func (e *Job) JobName() string {
	chunks := strings.Split(e.JobID, "/")
	return chunks[1]
}

// EffectiveSchedule returns schedule string to use for the job, considering its
// Paused field.
//
// Paused jobs always use "triggered" schedule.
func (e *Job) EffectiveSchedule() string {
	if e.Paused {
		return "triggered"
	}
	return e.Schedule
}

// ParseSchedule returns *Schedule object, parsing e.Schedule field.
//
// If job is paused e.Schedule field is ignored and "triggered" schedule is
// returned instead.
func (e *Job) ParseSchedule() (*schedule.Schedule, error) {
	if e.cachedSchedule == nil && e.cachedScheduleErr == nil {
		hash := fnv.New64()
		hash.Write([]byte(e.JobID))
		seed := hash.Sum64()
		e.cachedSchedule, e.cachedScheduleErr = schedule.Parse(e.EffectiveSchedule(), seed)
		if e.cachedSchedule == nil && e.cachedScheduleErr == nil {
			panic("no schedule and no error")
		}
	}
	return e.cachedSchedule, e.cachedScheduleErr
}

// IsEqual returns true iff 'e' is equal to 'other'.
func (e *Job) IsEqual(other *Job) bool {
	return e == other || (e.JobID == other.JobID &&
		e.ProjectID == other.ProjectID &&
		e.Flavor == other.Flavor &&
		e.Enabled == other.Enabled &&
		e.Paused == other.Paused &&
		e.Revision == other.Revision &&
		e.RevisionURL == other.RevisionURL &&
		e.Schedule == other.Schedule &&
		e.Acls.Equal(&other.Acls) &&
		bytes.Equal(e.Task, other.Task) &&
		equalSortedLists(e.TriggeredJobIDs, other.TriggeredJobIDs) &&
		e.State.Equal(&other.State))
}

// MatchesDefinition returns true if job definition in the entity matches the
// one specified by catalog.Definition struct.
func (e *Job) MatchesDefinition(def catalog.Definition) bool {
	return e.JobID == def.JobID &&
		e.Flavor == def.Flavor &&
		e.Schedule == def.Schedule &&
		e.Acls.Equal(&def.Acls) &&
		bytes.Equal(e.Task, def.Task) &&
		equalSortedLists(e.TriggeredJobIDs, def.TriggeredJobIDs)
}

// IsVisible checks if current identity has READER access to this job.
//
// Returns only transient errors.
func (e *Job) IsVisible(c context.Context) (bool, error) {
	return e.Acls.IsReader(logging.SetField(c, "JobID", e.JobID))
}

// IsOwned checks if current identity has OWNER access to this job.
//
// Returns only transient errors.
func (e *Job) IsOwned(c context.Context) (bool, error) {
	return e.Acls.IsOwner(logging.SetField(c, "JobID", e.JobID))
}