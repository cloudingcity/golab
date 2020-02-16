package project

import (
	"bytes"
	"errors"
	"net/url"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestList(t *testing.T) {
	project := "foo/bar"
	opt := &gitlab.ListProjectMergeRequestsOptions{}
	mockGitlabMR := &mocks.GitlabMergeRequests{}
	mockGitlabMR.On("ListProjectMergeRequests", project, opt).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{project: project, mr: mockGitlabMR, out: &bytes.Buffer{}}
	s.List(opt)

	mockGitlabMR.AssertExpectations(t)
}

func TestOpen(t *testing.T) {
	var openURL string
	stdout := &bytes.Buffer{}

	project := "foo/bar"
	baseURL, _ := url.Parse("https://gitlab.com")

	s := &mergeRequestsService{
		project: project,
		baseURL: baseURL,
		out:     stdout,
		openURL: func(url string) error {
			openURL = url
			return nil
		},
	}
	s.Open(123)

	assert.Equal(t, "Opening https://gitlab.com/foo/bar/merge_requests/123 in your browser\n", stdout.String())
	assert.Equal(t, "https://gitlab.com/foo/bar/merge_requests/123", openURL)
}

func TestShow(t *testing.T) {
	project := "foo/bar"
	mrID := 123
	mockGitlabMR := &mocks.GitlabMergeRequests{}
	mockGitlabMR.On("GetMergeRequest", project, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &mergeRequestsService{project: project, mr: mockGitlabMR, out: &bytes.Buffer{}}
	err := s.Show(mrID)

	assert.Error(t, err)
	mockGitlabMR.AssertExpectations(t)
}
