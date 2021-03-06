// Copyright 2016 The LUCI Authors.
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

package buildbucket

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/gae/service/datastore"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gitpb "go.chromium.org/luci/common/proto/git"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/milo/api/config"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/milo/common/model"
	"go.chromium.org/luci/milo/frontend/ui"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/router"
)

var ErrNotFound = errors.Reason("Build not found").Tag(common.CodeNotFound).Err()

// BuildAddress constructs the build address of a buildbucketpb.Build.
// This is used as the key for the BuildSummary entity.
func BuildAddress(build *buildbucketpb.Build) string {
	if build == nil {
		return ""
	}
	num := strconv.FormatInt(build.Id, 10)
	if build.Number != 0 {
		num = strconv.FormatInt(int64(build.Number), 10)
	}
	b := build.Builder
	return fmt.Sprintf("luci.%s.%s/%s/%s", b.Project, b.Bucket, b.Builder, num)
}

// simplisticBlamelist returns a slice of ui.Blame for a build.
//
// HACK(iannucci) - Getting the frontend to render a proper blamelist will
// require some significant refactoring. To do this properly, we'll need:
//   * The frontend to get BuildSummary from the backend.
//   * BuildSummary to have a .PreviousBuild() API.
//   * The frontend to obtain the annotation streams itself (so it could see
//     the SourceManifest objects inside of them). Currently getRespBuild defers
//     to swarming's implementation of buildsource.ID.Get(), which only returns
//     the resp object.
func simplisticBlamelist(c context.Context, build *model.BuildSummary) ([]*ui.Commit, error) {
	bs := build.GitilesCommit()
	if bs == nil {
		return nil, nil
	}

	builds, commits, err := build.PreviousByGitilesCommit(c)
	switch {
	case err == nil || err == model.ErrUnknownPreviousBuild:
	case status.Code(err) == codes.PermissionDenied:
		return nil, common.CodeUnauthorized.Tag().Apply(err)
	default:
		return nil, err
	}

	repoURL := bs.RepoURL()
	result := make([]*ui.Commit, 0, len(commits)+1)
	for _, commit := range commits {
		result = append(result, uiCommit(commit, repoURL))
	}
	logging.Infof(c, "Fetched %d commit blamelist from Gitiles", len(result))

	// this means that there were more than 100 commits in-between.
	if len(builds) == 0 && len(commits) > 0 {
		result = append(result, &ui.Commit{
			Description: "<blame list capped at 100 commits>",
			Revision:    &ui.Link{},
			AuthorName:  "<blame list capped at 100 commits>",
		})
	}

	return result, nil
}

func uiCommit(commit *gitpb.Commit, repoURL string) *ui.Commit {
	res := &ui.Commit{
		AuthorName:  commit.Author.Name,
		AuthorEmail: commit.Author.Email,
		Repo:        repoURL,
		Description: commit.Message,

		// TODO(iannucci): this use of links is very sloppy; the frontend should
		// know how to render a Commit without having Links embedded in it.
		Revision: ui.NewLink(
			commit.Id,
			repoURL+"/+/"+commit.Id, fmt.Sprintf("commit by %s", commit.Author.Email)),
	}
	res.CommitTime, _ = ptypes.Timestamp(commit.Committer.Time)
	res.File = make([]string, 0, len(commit.TreeDiff))
	for _, td := range commit.TreeDiff {
		// If file is moved, there is both new and old path,
		// from which we take only new path.
		// If a file is deleted, its new path is /dev/null.
		// In that case, we're only interested in the old path.
		switch {
		case td.NewPath != "" && td.NewPath != "/dev/null":
			res.File = append(res.File, td.NewPath)
		case td.OldPath != "":
			res.File = append(res.File, td.OldPath)
		}
	}
	return res
}

// GetBuildSummary fetches a build summary where the Context URI matches the
// given address.
func GetBuildSummary(c context.Context, id int64) (*model.BuildSummary, error) {
	// The host is set to prod because buildbot is hardcoded to talk to prod.
	uri := fmt.Sprintf("buildbucket://cr-buildbucket.appspot.com/build/%d", id)
	bs := make([]*model.BuildSummary, 0, 1)
	q := datastore.NewQuery("BuildSummary").Eq("ContextURI", uri).Limit(1)
	switch err := datastore.GetAll(c, q, &bs); {
	case err != nil:
		return nil, common.ReplaceNSEWith(err.(errors.MultiError), ErrNotFound)
	case len(bs) == 0:
		return nil, ErrNotFound
	default:
		return bs[0], nil
	}
}

// getBlame fetches blame information from Gitiles.
// This requires the BuildSummary to be indexed in Milo.
func getBlame(c context.Context, host string, b *buildbucketpb.Build) ([]*ui.Commit, error) {
	commit := b.GetInput().GetGitilesCommit()
	// No commit? No blamelist.
	if commit == nil {
		return nil, nil
	}
	// TODO(hinoka): This converts a buildbucketpb.Commit into a string
	// and back into a buildbucketpb.Commit.  That's a bit silly.
	return simplisticBlamelist(c, &model.BuildSummary{
		BuildKey:  MakeBuildKey(c, host, BuildAddress(b)),
		BuildSet:  []string{commit.BuildSetString()},
		BuilderID: BuilderID{BuilderID: *b.Builder}.String(),
	})
}

func buildbucketClient(c context.Context, host string) (buildbucketpb.BuildsClient, error) {
	t, err := auth.GetRPCTransport(c, auth.AsUser)
	if err != nil {
		return nil, err
	}
	return buildbucketpb.NewBuildsPRPCClient(&prpc.Client{
		C:    &http.Client{Transport: t},
		Host: host,
	}), nil
}

// GetBuild returns the buildbucketpb.Build from a build request.
func GetBuild(c context.Context, host string, bid buildbucketpb.GetBuildRequest) (*buildbucketpb.Build, error) {
	client, err := buildbucketClient(c, host)
	if err != nil {
		return nil, err
	}
	return client.GetBuild(c, &bid)
}

var fullBuildMask = &field_mask.FieldMask{
	// TODO(hinoka): Add statusReason here.
	Paths: []string{
		"id",
		"builder",
		"number",
		"created_by",
		"create_time",
		"start_time",
		"end_time",
		"update_time",
		"status",
		"input",
		"output",
		"steps",
		"infra",
		"tags",
	},
}

// getBugLink attempts to formulate and return the build page bug link
// for the given build.
func getBugLink(c *router.Context, b *buildbucketpb.Build) (string, error) {
	project, err := common.GetProject(c.Context, b.Builder.GetProject())
	if err != nil || proto.Equal(&project.BuildBugTemplate, &config.BugTemplate{}) {
		return "", err
	}

	builderPath := fmt.Sprintf("/p/%s/builders/%s/%s", b.Builder.GetProject(), b.Builder.GetBucket(), b.Builder.GetBuilder())
	buildURL, err := c.Request.URL.Parse(builderPath + "/" + c.Params.ByName("numberOrId"))
	if err != nil {
		return "", errors.Annotate(err, "Unable to make build URL for build bug link.").Err()
	}
	builderURL, err := c.Request.URL.Parse(builderPath)
	if err != nil {
		return "", errors.Annotate(err, "Unable to make builder URL for build bug link.").Err()
	}

	return MakeBuildBugLink(&project.BuildBugTemplate, map[string]interface{}{
		"Build":          b,
		"MiloBuildUrl":   buildURL,
		"MiloBuilderUrl": builderURL,
	})
}

// GetBuildPage fetches the full set of information for a Milo build page from Buildbucket.
// Including the blamelist and other auxiliary information.
func GetBuildPage(c *router.Context, br buildbucketpb.GetBuildRequest) (*ui.BuildPage, error) {
	br.Fields = fullBuildMask
	host, err := getHost(c.Context)
	if err != nil {
		return nil, err
	}
	b, err := GetBuild(c.Context, host, br)
	if err != nil {
		return nil, err
	}
	blame, err := getBlame(c.Context, host, b)
	if err != nil {
		return nil, err
	}
	link, err := getBugLink(c, b)

	bp := ui.NewBuildPage(c.Context, b)
	bp.Blame = blame
	bp.BuildBugLink = link
	return bp, err
}
