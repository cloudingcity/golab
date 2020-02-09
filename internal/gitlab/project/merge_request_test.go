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
	var got string

	project := "foo/bar"
	baseURL, _ := url.Parse("https://gitlab.com")

	s := &mergeRequestsService{
		project: project,
		baseURL: baseURL,
		openURL: func(url string) error {
			got = url
			return nil
		},
	}
	s.Open("123")

	assert.Equal(t, "https://gitlab.com/foo/bar/merge_requests/123", got)
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
