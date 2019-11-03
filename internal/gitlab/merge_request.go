package gitlab

import (
	"fmt"
	"io"
	"strconv"

	"github.com/xanzy/go-gitlab"
)

type mergeRequestService interface {
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
}

type mergeRequest struct {
	mr  mergeRequestService
	out io.Writer
}

// List lists merge requests on a project.
func (s *mergeRequest) List(project string, opt *gitlab.ListProjectMergeRequestsOptions) error {
	mrs, _, err := s.mr.ListProjectMergeRequests(project, opt)
	if err != nil {
		return err
	}

	s.render(mrs)
	return nil
}

func (s *mergeRequest) render(mrs []*gitlab.MergeRequest) {
	f := "  #%s  %s\n"
	for _, mr := range mrs {
		id := strconv.Itoa(mr.IID)
		fmt.Fprintf(s.out, f, id, mr.Title)
	}
}
