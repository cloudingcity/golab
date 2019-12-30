package project

import (
	"errors"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	project := "foo/bar"
	query := "query string"
	search := &mocks.GitlabSearch{}
	search.On("MergeRequestsByProject", project, query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{project: project, search: search}
	err := s.MR(query)

	assert.Error(t, err)
	search.AssertExpectations(t)
}
