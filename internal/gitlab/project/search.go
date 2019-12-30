package project

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

// GitlabSearchService is go-gitlab search service interface.
type GitlabSearchService interface {
	MergeRequestsByProject(pid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

type searchService struct {
	project string
	search  GitlabSearchService
	out     io.Writer
}

func (s *searchService) MR(query string) error {
	mrs, _, err := s.search.MergeRequestsByProject(s.project, query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).ProjectMRs(mrs)
	return nil
}
