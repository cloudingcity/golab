package global

import (
	"fmt"
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type mergeRequestsService struct {
	gitlabMR contract.GitlabMergeRequests
	out      io.Writer
	openURL  func(url string) error
}

// List lists merge requests on a project.
func (s *mergeRequestsService) List(opt *gitlab.ListMergeRequestsOptions) error {
	mrs, _, err := s.gitlabMR.ListMergeRequests(opt)
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(pID, mrID int) error {
	mr, _, err := s.gitlabMR.GetMergeRequest(pID, mrID, nil)
	if err != nil {
		return err
	}

	fmt.Fprintf(s.out, "Opening %s in your browser\n", mr.WebURL)
	return s.openURL(mr.WebURL)
}

// Show show a merge request on a project
func (s *mergeRequestsService) Show(pID string, mrID int) error {
	mr, _, err := s.gitlabMR.GetMergeRequest(pID, mrID, nil)
	if err != nil {
		return err
	}

	render.New(s.out).MR(mr)
	return nil
}
