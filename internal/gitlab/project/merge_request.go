package project

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"strconv"

	"github.com/cloudingcity/golab/internal/git"
	"github.com/cloudingcity/golab/internal/gitlab/contract"
	"github.com/cloudingcity/golab/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type mergeRequestsService struct {
	project       string
	gitlabMR      contract.GitlabMergeRequests
	gitlabProject contract.GitlabProject
	out           io.Writer
	baseURL       *url.URL
	openURL       func(url string) error
}

// List lists merge requests on a project.
func (s *mergeRequestsService) List(opt *gitlab.ListProjectMergeRequestsOptions) error {
	mrs, _, err := s.gitlabMR.ListProjectMergeRequests(s.project, opt)
	if err != nil {
		return err
	}

	render.New(s.out).ProjectMRs(mrs)
	return nil
}

// Open browse merge request in the default browser.
func (s *mergeRequestsService) Open(mrID int) error {
	u := *s.baseURL
	u.Path = path.Join(s.project, "merge_requests", strconv.Itoa(mrID))

	fmt.Fprintf(s.out, "Opening %s in your browser\n", u.String())
	return s.openURL(u.String())
}

// Show show a merge request on a project
func (s *mergeRequestsService) Show(mrID int) error {
	mr, _, err := s.gitlabMR.GetMergeRequest(s.project, mrID, nil)
	if err != nil {
		return err
	}

	render.New(s.out).MR(mr)
	return nil
}

// Create create a merge request.
func (s *mergeRequestsService) Create() error {
	project, _, err := s.gitlabProject.GetProject(s.project, &gitlab.GetProjectOptions{})
	if err != nil {
		return err
	}

	defaultBranch := project.DefaultBranch
	currentBranch := git.CurrentBranch()
	if defaultBranch == currentBranch {
		return fmt.Errorf("must be on a branch named differently than %q", defaultBranch)
	}

	if err := git.Push(currentBranch); err != nil {
		return err
	}

	u := *s.baseURL
	u.Path = path.Join(s.project, "merge_requests", "new")
	q := make(url.Values)
	q.Set("merge_request[source_branch]", currentBranch)
	u.RawQuery = q.Encode()

	fmt.Fprintf(s.out, "Opening %s in your browser\n", u.String())
	return s.openURL(u.String())
}
