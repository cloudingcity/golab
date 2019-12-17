package global

import (
	"fmt"
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

// Show show a merge request on a project
func (s *mergeRequestsService) Show(pID string, mrID int) error {
	mr, _, err := s.mr.GetMergeRequest(pID, mrID, &gitlab.GetMergeRequestsOptions{})
	if err != nil {
		return err
	}

	s.renderShow(mr)
	return nil
}

func (s *mergeRequestsService) renderShow(mr *gitlab.MergeRequest) {
	var assignee string
	if mr.Assignee != nil {
		assignee = mr.Assignee.Username
	}
	createdAt := mr.CreatedAt.Format("2006-01-02 15:04:05")
	updatedAt := mr.UpdatedAt.Format("2006-01-02 15:04:05")

	format := `
%s
--------------------------------------------------
%s
--------------------------------------------------
PID         %d  
MRID        %d
Project     %s
Branch      %s -> %s
State       %s
Author      %s
Assignee    %s
CreatedAt   %s
UpdatedAt   %s
Url         %s
`
	fmt.Fprintf(s.out, format,
		mr.Title,
		mr.Description,
		mr.ProjectID,
		mr.IID,
		utils.ParseMRProject(mr.WebURL),
		mr.SourceBranch, mr.TargetBranch,
		mr.State,
		mr.Author.Username,
		assignee,
		createdAt,
		updatedAt,
		mr.WebURL,
	)
}
