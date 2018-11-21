// Copyright 2018 The LUCI Authors.
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

package backend

import (
	"context"
	"testing"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/luci/appengine/tq"
	"go.chromium.org/luci/appengine/tq/tqtesting"

	"go.chromium.org/luci/gce/api/config/v1"
	"go.chromium.org/luci/gce/api/tasks/v1"
	rpc "go.chromium.org/luci/gce/appengine/rpc/memory"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestQueues(t *testing.T) {
	t.Parallel()

	Convey("expand", t, func() {
		dsp := &tq.Dispatcher{}
		registerTasks(dsp)
		srv := &rpc.Config{}
		c := withServer(withDispatcher(memory.Use(context.Background()), dsp), srv)
		tqt := tqtesting.GetTestable(c, dsp)
		tqt.CreateQueues()

		Convey("invalid", func() {
			Convey("nil", func() {
				err := expand(c, nil)
				So(err, ShouldErrLike, "unexpected payload")
				So(tqt.GetScheduledTasks(), ShouldBeEmpty)
			})

			Convey("empty", func() {
				err := expand(c, &tasks.Expand{})
				So(err, ShouldErrLike, "ID is required")
				So(tqt.GetScheduledTasks(), ShouldBeEmpty)
			})
		})

		Convey("valid", func() {
			Convey("missing", func() {
				err := expand(c, &tasks.Expand{Id: "id"})
				So(err, ShouldErrLike, "failed to get VMs block")
				So(tqt.GetScheduledTasks(), ShouldBeEmpty)
			})

			Convey("zero", func() {
				srv.EnsureVMs(c, &config.EnsureVMsRequest{
					Id:  "id",
					Vms: &config.Block{},
				})
				err := expand(c, &tasks.Expand{Id: "id"})
				So(err, ShouldBeNil)
				So(tqt.GetScheduledTasks(), ShouldBeEmpty)
			})

			Convey("one", func() {
				srv.EnsureVMs(c, &config.EnsureVMsRequest{
					Id: "id",
					Vms: &config.Block{
						Amount: 1,
					},
				})
				err := expand(c, &tasks.Expand{Id: "id"})
				So(err, ShouldBeNil)
				So(tqt.GetScheduledTasks(), ShouldHaveLength, 1)
			})
		})
	})
}
