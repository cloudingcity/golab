package global

import (
	"io"
	"strconv"

	"github.com/cloudingcity/golab/internal/utils"
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
func (s *mergeRequestsService) List(opt *gitlab.ListMergeRequestsOptions, withURL bool) error {
	mrs, _, err := s.mr.ListMergeRequests(opt)
	if err != nil {
		return err
	}

	s.renderList(mrs, withURL)
	return nil
}

func (s *mergeRequestsService) renderList(mrs []*gitlab.MergeRequest, withURL bool) {
	var (
		rows [][]string
		row  []string
		h    []string
	)

	if withURL {
		h = []string{"pid", "mrid", "project", "title", "url"}
	} else {
		h = []string{"pid", "mrid", "project", "title"}
	}

	for _, mr := range mrs {
		pID := strconv.Itoa(mr.ProjectID)
		mrID := strconv.Itoa(mr.IID)
		p := utils.ParseMRProject(mr.WebURL)

		if withURL {
			row = []string{pID, mrID, p, mr.Title, mr.WebURL}
		} else {
			row = []string{pID, mrID, p, mr.Title}
		}

		rows = append(rows, row)
	}

	utils.RenderTable(s.out, h, rows)
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(pID string, mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(pID, mrID, nil)
	if err != nil {
		return err
	}

	return browser.OpenURL(mr.WebURL)
}
