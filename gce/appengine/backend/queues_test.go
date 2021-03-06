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
	"net/http"
	"testing"

	"google.golang.org/api/compute/v1"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/appengine/tq"
	"go.chromium.org/luci/appengine/tq/tqtesting"

	"go.chromium.org/luci/gce/api/config/v1"
	"go.chromium.org/luci/gce/api/tasks/v1"
	"go.chromium.org/luci/gce/appengine/model"
	rpc "go.chromium.org/luci/gce/appengine/rpc/memory"
	"go.chromium.org/luci/gce/appengine/testing/roundtripper"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestQueues(t *testing.T) {
	t.Parallel()

	Convey("queues", t, func() {
		dsp := &tq.Dispatcher{}
		registerTasks(dsp)
		srv := &rpc.Config{}
		rt := &roundtripper.JSONRoundTripper{}
		gce, err := compute.New(&http.Client{Transport: rt})
		So(err, ShouldBeNil)
		c := withCompute(withConfig(withDispatcher(memory.Use(context.Background()), dsp), srv), gce)
		datastore.GetTestable(c).AutoIndex(true)
		datastore.GetTestable(c).Consistent(true)
		tqt := tqtesting.GetTestable(c, dsp)
		tqt.CreateQueues()

		Convey("createVM", func() {
			Convey("invalid", func() {
				Convey("nil", func() {
					err := createVM(c, nil)
					So(err, ShouldErrLike, "unexpected payload")
				})

				Convey("empty", func() {
					err := createVM(c, &tasks.CreateVM{})
					So(err, ShouldErrLike, "config is required")
				})
			})

			Convey("valid", func() {
				Convey("nil", func() {
					err := createVM(c, &tasks.CreateVM{
						Index:  2,
						Config: "id",
					})
					So(err, ShouldBeNil)
					err = datastore.Get(c, &model.VM{
						ID: "id-2",
					})
					So(err, ShouldBeNil)
				})

				Convey("empty", func() {
					err := createVM(c, &tasks.CreateVM{
						Attributes: &config.VM{},
						Index:      2,
						Config:     "id",
					})
					So(err, ShouldBeNil)
					err = datastore.Get(c, &model.VM{
						ID:      "id-2",
						Drained: false,
					})
					So(err, ShouldBeNil)
				})

				Convey("non-empty", func() {
					err := createVM(c, &tasks.CreateVM{
						Attributes: &config.VM{
							Disk: []*config.Disk{
								{
									Image: "image",
								},
							},
						},
						Index:  2,
						Config: "id",
					})
					So(err, ShouldBeNil)
					v := &model.VM{
						ID: "id-2",
					}
					err = datastore.Get(c, v)
					So(err, ShouldBeNil)
					So(v, ShouldResemble, &model.VM{
						ID: "id-2",
						Attributes: config.VM{
							Disk: []*config.Disk{
								{
									Image: "image",
								},
							},
						},
						Config:  "id",
						Drained: false,
						Index:   2,
					})
				})

				Convey("not updated", func() {
					datastore.Put(c, &model.VM{
						ID: "id-2",
						Attributes: config.VM{
							Zone: "zone",
						},
						Drained: true,
					})
					err := createVM(c, &tasks.CreateVM{
						Attributes: &config.VM{
							Project: "project",
						},
						Index:  2,
						Config: "id",
					})
					So(err, ShouldBeNil)
					v := &model.VM{
						ID: "id-2",
					}
					err = datastore.Get(c, v)
					So(err, ShouldBeNil)
					So(v, ShouldResemble, &model.VM{
						ID: "id-2",
						Attributes: config.VM{
							Zone: "zone",
						},
						Drained: true,
					})
				})
			})
		})

		Convey("drainVM", func() {
			Convey("invalid", func() {
				Convey("nil", func() {
					err := drainVM(c, nil)
					So(err, ShouldErrLike, "unexpected payload")
				})

				Convey("empty", func() {
					err := drainVM(c, &tasks.DrainVM{})
					So(err, ShouldErrLike, "ID is required")
				})

				Convey("config", func() {
					datastore.Put(c, &model.VM{
						ID: "id",
					})
					err := drainVM(c, &tasks.DrainVM{
						Id: "id",
					})
					So(err, ShouldErrLike, "failed to fetch config")
				})
			})

			Convey("valid", func() {
				Convey("config", func() {
					Convey("deleted", func() {
						datastore.Put(c, &model.VM{
							ID:     "id",
							Config: "config",
						})
						err := drainVM(c, &tasks.DrainVM{
							Id: "id",
						})
						So(err, ShouldBeNil)
						v := &model.VM{
							ID: "id",
						}
						datastore.Get(c, v)
						So(v.Drained, ShouldBeTrue)
					})

					Convey("amount", func() {
						Convey("unspecified", func() {
							datastore.Put(c, &model.Config{
								ID: "config",
							})
							datastore.Put(c, &model.VM{
								ID:     "id",
								Config: "config",
							})
							err := drainVM(c, &tasks.DrainVM{
								Id: "id",
							})
							So(err, ShouldBeNil)
							v := &model.VM{
								ID: "id",
							}
							datastore.Get(c, v)
							So(v.Drained, ShouldBeTrue)
						})

						Convey("matches", func() {
							datastore.Put(c, &model.Config{
								ID: "config",
								Config: config.Config{
									Amount: &config.Amount{
										Default: 2,
									},
								},
							})
							datastore.Put(c, &model.VM{
								ID:     "id",
								Config: "config",
								Index:  2,
							})
							err := drainVM(c, &tasks.DrainVM{
								Id: "id",
							})
							So(err, ShouldBeNil)
							v := &model.VM{
								ID: "id",
							}
							datastore.Get(c, v)
							So(v.Drained, ShouldBeTrue)
						})

						Convey("exceeds", func() {
							datastore.Put(c, &model.Config{
								ID: "config",
								Config: config.Config{
									Amount: &config.Amount{
										Default: 1,
									},
								},
							})
							datastore.Put(c, &model.VM{
								ID:     "id",
								Config: "config",
								Index:  2,
							})
							err := drainVM(c, &tasks.DrainVM{
								Id: "id",
							})
							So(err, ShouldBeNil)
							v := &model.VM{
								ID: "id",
							}
							datastore.Get(c, v)
							So(v.Drained, ShouldBeTrue)
						})

						Convey("active", func() {
							datastore.Put(c, &model.Config{
								ID: "config",
								Config: config.Config{
									Amount: &config.Amount{
										Default: 3,
									},
								},
							})
							datastore.Put(c, &model.VM{
								ID:     "id",
								Config: "config",
								Index:  2,
							})
							err := drainVM(c, &tasks.DrainVM{
								Id: "id",
							})
							So(err, ShouldBeNil)
							v := &model.VM{
								ID: "id",
							}
							datastore.Get(c, v)
							So(v.Drained, ShouldBeFalse)
						})
					})
				})

				Convey("deleted", func() {
					err := drainVM(c, &tasks.DrainVM{
						Id: "id",
					})
					So(err, ShouldBeNil)
					err = datastore.Get(c, &model.VM{
						ID: "id",
					})
					So(err, ShouldEqual, datastore.ErrNoSuchEntity)
				})
			})
		})

		Convey("expandConfig", func() {
			Convey("invalid", func() {
				Convey("nil", func() {
					err := expandConfig(c, nil)
					So(err, ShouldErrLike, "unexpected payload")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})

				Convey("empty", func() {
					err := expandConfig(c, &tasks.ExpandConfig{})
					So(err, ShouldErrLike, "ID is required")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})

				Convey("missing", func() {
					err := expandConfig(c, &tasks.ExpandConfig{
						Id: "id",
					})
					So(err, ShouldErrLike, "failed to fetch config")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})
			})

			Convey("valid", func() {
				Convey("none", func() {
					srv.Ensure(c, &config.EnsureRequest{
						Id:     "id",
						Config: &config.Config{},
					})
					err := expandConfig(c, &tasks.ExpandConfig{
						Id: "id",
					})
					So(err, ShouldBeNil)
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})

				Convey("create", func() {
					srv.Ensure(c, &config.EnsureRequest{
						Id: "id",
						Config: &config.Config{
							Amount: &config.Amount{
								Default: 3,
							},
						},
					})
					err := expandConfig(c, &tasks.ExpandConfig{
						Id: "id",
					})
					So(err, ShouldBeNil)
					So(tqt.GetScheduledTasks(), ShouldHaveLength, 3)
				})
			})
		})

		Convey("reportQuota", func() {
			Convey("invalid", func() {
				Convey("nil", func() {
					err := reportQuota(c, nil)
					So(err, ShouldErrLike, "unexpected payload")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})

				Convey("empty", func() {
					err := reportQuota(c, &tasks.ReportQuota{})
					So(err, ShouldErrLike, "ID is required")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})

				Convey("missing", func() {
					err := reportQuota(c, &tasks.ReportQuota{
						Id: "id",
					})
					So(err, ShouldErrLike, "failed to fetch project")
					So(tqt.GetScheduledTasks(), ShouldBeEmpty)
				})
			})

			Convey("valid", func() {
				rt.Handler = func(req interface{}) (int, interface{}) {
					return http.StatusOK, &compute.RegionList{
						Items: []*compute.Region{
							{
								Name: "ignore",
							},
							{
								Name: "region",
								Quotas: []*compute.Quota{
									{
										Limit:  100.0,
										Metric: "ignore",
										Usage:  0.0,
									},
									{
										Limit:  100.0,
										Metric: "metric",
										Usage:  50.0,
									},
								},
							},
						},
					}
				}
				datastore.Put(c, &model.Project{
					ID:      "id",
					Metrics: []string{"metric"},
					Project: "project",
					Regions: []string{"region"},
				})
				err := reportQuota(c, &tasks.ReportQuota{
					Id: "id",
				})
				So(err, ShouldBeNil)
			})
		})
	})
}
