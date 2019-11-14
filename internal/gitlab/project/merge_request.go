package project

import (
	"io"
	"net/url"
	"path"
	"strconv"

	"github.com/cloudingcity/golab/internal/utils"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type gitlabMergeRequestsService interface {
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

type mergeRequestsService struct {
	project string
	mr      gitlabMergeRequestsService
	out     io.Writer
	url     *url.URL
}

// List lists merge requests on a project.
func (s *mergeRequestsService) List(opt *gitlab.ListProjectMergeRequestsOptions, withURL bool) error {
	mrs, _, err := s.mr.ListProjectMergeRequests(s.project, opt)
	if err != nil {
		return err
	}

	s.renderList(mrs, withURL)
	return nil
}

func (s *mergeRequestsService) renderList(mrs []*gitlab.MergeRequest, withURL bool) {
	var (
		rows   [][]string
		row, h []string
	)

	if withURL {
		h = []string{"mrid", "title", "url"}
	} else {
		h = []string{"mrid", "title"}
	}

	for _, mr := range mrs {
		if withURL {
			row = []string{strconv.Itoa(mr.IID), mr.Title, mr.WebURL}
		} else {
			row = []string{strconv.Itoa(mr.IID), mr.Title}
		}
		rows = append(rows, row)
	}

	utils.RenderTable(s.out, h, rows)
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return errors.Errorf("invalid merge request id: '%s'", id)
	}
	s.url.Path = path.Join(s.project, "merge_requests", id)

	return browser.OpenURL(s.url.String())
}
