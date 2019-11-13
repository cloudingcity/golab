package global

import (
	"io"
	"strconv"

	"github.com/cloudingcity/golab/internal/utils"
	"github.com/xanzy/go-gitlab"
)

type gitlabMergeRequestsService interface {
	ListMergeRequests(opt *gitlab.ListMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
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
	table := utils.NewTable(s.out)
	table.SetHeader([]string{"id", "project", "title"})
	for _, mr := range mrs {
		id := strconv.Itoa(mr.ProjectID) + "-" + strconv.Itoa(mr.IID)
		p := utils.ParseMRProject(mr.WebURL)
		table.Append([]string{id, p, mr.Title})
	}
	table.Render()
}
