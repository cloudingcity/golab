package global

import (
	"io"
	"strconv"

	"github.com/cloudingcity/golab/internal/utils"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
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

	s.renderList(mrs)
	return nil
}

func (s *mergeRequestsService) renderList(mrs []*gitlab.MergeRequest) {
	var rows [][]string

	h := []string{"pid", "mrid", "project", "title"}

	for _, mr := range mrs {
		pID := strconv.Itoa(mr.ProjectID)
		mrID := strconv.Itoa(mr.IID)
		p := utils.ParseMRProject(mr.WebURL)
		rows = append(rows, []string{pID, mrID, p, mr.Title})
	}

	utils.RenderTable(s.out, h, rows)
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(projectID, mrID string) error {
	if _, err := strconv.Atoi(projectID); err != nil {
		return errors.Errorf("invalid project id: '%s'", projectID)
	}

	mrIntID, err := strconv.Atoi(mrID)
	if err != nil {
		return errors.Errorf("merge request id: '%s'", mrID)
	}

	mr, _, err := s.mr.GetMergeRequest(projectID, mrIntID, nil)
	if err != nil {
		return err
	}

	return browser.OpenURL(mr.WebURL)
}
