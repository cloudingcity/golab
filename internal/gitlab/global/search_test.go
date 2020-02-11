package global

import (
	"errors"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("MergeRequests", query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{search: mockGitlabSearch}
	err := s.MR(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}

func TestProject(t *testing.T) {
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("Projects", query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.Project{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{search: mockGitlabSearch}
	err := s.Project(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}
