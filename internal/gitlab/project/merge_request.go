package project

import (
	"io"
	"net/url"
	"path"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/pkg/browser"
	"github.com/xanzy/go-gitlab"
)

type gitlabMergeRequestsService interface {
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
}

type mergeRequestsService struct {
	project string
	mr      gitlabMergeRequestsService
	out     io.Writer
	baseURL *url.URL
}

// List lists merge requests on a project.
func (s *mergeRequestsService) List(opt *gitlab.ListProjectMergeRequestsOptions) error {
	mrs, _, err := s.mr.ListProjectMergeRequests(s.project, opt)
	if err != nil {
		return err
	}

	render.New(s.out).ProjectMRs(mrs)
	return nil
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(mrID string) error {
	u := *s.baseURL
	u.Path = path.Join(s.project, "merge_requests", mrID)

	return browser.OpenURL(u.String())
}

// Show show a merge request on a project
func (s *mergeRequestsService) Show(mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(s.project, mrID, &gitlab.GetMergeRequestsOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).MR(mr)
	return nil
}
