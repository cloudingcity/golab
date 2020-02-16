package global

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type searchService struct {
	gitlabSearch contract.GitlabSearch
	out          io.Writer
}

func (s *searchService) MR(query string) error {
	mrs, _, err := s.gitlabSearch.MergeRequests(query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}

func (s *searchService) Project(query string) error {
	projects, _, err := s.gitlabSearch.Projects(query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).Projects(projects)
	return nil
}
