package global

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/pkg/browser"
	"github.com/xanzy/go-gitlab"
)

type gitlabMergeRequestsService interface {
	ListMergeRequests(opt *gitlab.ListMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
}

type mergeRequestsService struct {
	mr  gitlabMergeRequestsService
	out io.Writer
}

// List lists merge requests on a project.
func (s *mergeRequestsService) List(opt *gitlab.ListMergeRequestsOptions) error {
	mrs, _, err := s.mr.ListMergeRequests(opt)
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(pID string, mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(pID, mrID, nil)
	if err != nil {
		return err
	}

	return browser.OpenURL(mr.WebURL)
}

// Show show a merge request on a project
func (s *mergeRequestsService) Show(pID string, mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(pID, mrID, &gitlab.GetMergeRequestsOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).MR(mr)
	return nil
}
