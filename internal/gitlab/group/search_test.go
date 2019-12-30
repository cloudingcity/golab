package group

import (
	"errors"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/group/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	group := "foo"
	query := "query string"
	search := &mocks.GitlabSearchService{}
	search.On("MergeRequestsByGroup", group, query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{group: group, search: search}
	err := s.MR(query)

	assert.Error(t, err)
	search.AssertExpectations(t)
}
