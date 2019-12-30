package group

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

// GitlabSearchService is go-gitlab search service interface.
type GitlabSearchService interface {
	MergeRequestsByGroup(gid interface{}, query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

type searchService struct {
	group  string
	search GitlabSearchService
	out    io.Writer
}

func (s *searchService) MR(query string) error {
	mrs, _, err := s.search.MergeRequestsByGroup(s.group, query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}
