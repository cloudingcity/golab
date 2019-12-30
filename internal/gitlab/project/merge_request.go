package project

import (
	"io"
	"net/url"
	"path"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type mergeRequestsService struct {
	project string
	mr      contract.GitlabMergeRequests
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
