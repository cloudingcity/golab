package project

import (
	"io"
	"net/url"
	"path"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

// GitlabMergeRequestsService is go-gitlab merge request service interface.
type GitlabMergeRequestsService interface {
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
}

type mergeRequestsService struct {
	project string
	mr      GitlabMergeRequestsService
	out     io.Writer
	baseURL *url.URL
	openURL func(url string) error
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

	return s.openURL(u.String())
}

// Show show a merge request on a project
func (s *mergeRequestsService) Show(mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(s.project, mrID, nil)
	if err != nil {
		return err
	}

	render.New(s.out).MR(mr)
	return nil
}
