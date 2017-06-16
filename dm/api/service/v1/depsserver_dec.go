// Code generated by svcdec; DO NOT EDIT

package dm

import (
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
)

type DecoratedDeps struct {
	// Service is the service to decorate.
	Service DepsServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedDeps) EnsureGraphData(c context.Context, req *EnsureGraphDataReq) (rsp *EnsureGraphDataRsp, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "EnsureGraphData", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.EnsureGraphData(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "EnsureGraphData", rsp, err)
	}
	return
}

func (s *DecoratedDeps) ActivateExecution(c context.Context, req *ActivateExecutionReq) (rsp *google_protobuf2.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ActivateExecution", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ActivateExecution(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ActivateExecution", rsp, err)
	}
	return
}

func (s *DecoratedDeps) FinishAttempt(c context.Context, req *FinishAttemptReq) (rsp *google_protobuf2.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "FinishAttempt", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.FinishAttempt(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "FinishAttempt", rsp, err)
	}
	return
}

func (s *DecoratedDeps) WalkGraph(c context.Context, req *WalkGraphReq) (rsp *GraphData, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "WalkGraph", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.WalkGraph(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "WalkGraph", rsp, err)
	}
	return
}
