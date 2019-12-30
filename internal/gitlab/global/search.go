package global

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

// GitlabSearchService is go-gitlab search service interface.
type GitlabSearchService interface {
	MergeRequests(query string, opt *gitlab.SearchOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

type searchService struct {
	search GitlabSearchService
	out    io.Writer
}

func (s *searchService) MR(query string) error {
	mrs, _, err := s.search.MergeRequests(query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}
