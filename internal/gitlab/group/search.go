package group

import (
	"io"

	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type searchService struct {
	group  string
	search contract.GitlabSearch
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
