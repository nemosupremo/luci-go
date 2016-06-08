// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package deps

import (
	"testing"
	"time"

	"github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/appengine/cmd/dm/model"
	"github.com/luci/luci-go/appengine/tumble"
	"github.com/luci/luci-go/common/api/dm/service/v1"
	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddDeps(t *testing.T) {
	t.Parallel()

	Convey("EnsureGraphData (Adding deps)", t, func() {
		ttest := &tumble.Testing{}
		c := ttest.Context()
		ds := datastore.Get(c)
		s := newDecoratedDeps()
		zt := time.Time{}

		a := &model.Attempt{ID: *dm.NewAttemptID("quest", 1)}
		a.CurExecution = 1
		a.State = dm.Attempt_EXECUTING
		ak := ds.KeyForObj(a)

		e := &model.Execution{
			ID: 1, Attempt: ak, Token: []byte("key"),
			State: dm.Execution_RUNNING}

		toQuestDesc := &dm.Quest_Desc{
			DistributorConfigName: "foof",
			JsonPayload:           `{"data":"yes"}`,
		}
		toQuest, err := model.NewQuest(c, toQuestDesc)
		So(err, ShouldBeNil)
		to := &model.Attempt{ID: *dm.NewAttemptID(toQuest.ID, 1)}
		fwd := &model.FwdDep{Depender: ak, Dependee: to.ID}

		req := &dm.EnsureGraphDataReq{
			ForExecution: &dm.Execution_Auth{
				Id:    dm.NewExecutionID(a.ID.Quest, a.ID.Id, 1),
				Token: []byte("key"),
			},
			Attempts: dm.NewAttemptList(map[string][]uint32{
				to.ID.Quest: {to.ID.Id},
			}),
		}

		Convey("Bad", func() {
			Convey("No such originating attempt", func() {
				_, err := s.EnsureGraphData(c, req)
				So(err, ShouldBeRPCUnauthenticated)
			})

			Convey("No such destination quest", func() {
				So(ds.Put(a, e), ShouldBeNil)

				_, err := s.EnsureGraphData(c, req)
				So(err, ShouldBeRPCInvalidArgument, `cannot create attempts for absent quest "Q9SgH-f5kraxP_om80CdR9EmAvgmnUws_s5fvRmZiuc"`)
			})
		})

		Convey("Good", func() {
			So(ds.Put(a, e, toQuest), ShouldBeNil)

			Convey("deps already exist", func() {
				So(ds.Put(fwd, to), ShouldBeNil)

				rsp, err := s.EnsureGraphData(c, req)
				So(err, ShouldBeNil)
				purgeTimestamps(rsp.Result)
				So(rsp, ShouldResemble, &dm.EnsureGraphDataRsp{
					Accepted: true,
					Result: &dm.GraphData{Quests: map[string]*dm.Quest{
						toQuest.ID: {
							Data: &dm.Quest_Data{
								Desc:    toQuestDesc,
								BuiltBy: []*dm.Quest_TemplateSpec{},
							},
							Attempts: map[uint32]*dm.Attempt{1: dm.NewAttemptNeedsExecution(zt)},
						},
					}},
				})
			})

			Convey("deps already done", func() {
				to.State = dm.Attempt_FINISHED
				So(ds.Put(to), ShouldBeNil)

				rsp, err := s.EnsureGraphData(c, req)
				So(err, ShouldBeNil)
				purgeTimestamps(rsp.Result)
				So(rsp, ShouldResemble, &dm.EnsureGraphDataRsp{
					Accepted: true,
					Result: &dm.GraphData{Quests: map[string]*dm.Quest{
						toQuest.ID: {
							Data: &dm.Quest_Data{
								Desc:    toQuestDesc,
								BuiltBy: []*dm.Quest_TemplateSpec{},
							},
							Attempts: map[uint32]*dm.Attempt{1: dm.NewAttemptFinished(zt, 0, "")},
						},
					}},
				})

				So(ds.Get(fwd), ShouldBeNil)
			})

			Convey("adding new deps", func() {
				So(ds.Put(&model.Quest{ID: "to"}), ShouldBeNil)

				rsp, err := s.EnsureGraphData(c, req)
				So(err, ShouldBeNil)
				So(rsp, ShouldResemble, &dm.EnsureGraphDataRsp{ShouldHalt: true})

				So(ds.Get(fwd), ShouldBeNil)
				So(ds.Get(a), ShouldBeNil)
				So(a.State, ShouldEqual, dm.Attempt_ADDING_DEPS)
			})

		})
	})
}
