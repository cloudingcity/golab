package project

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type searchService struct {
	project string
	search  contract.GitlabSearch
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
