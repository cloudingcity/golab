package project

import (
	"bytes"
	"errors"
	"net/url"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/project/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestList(t *testing.T) {
	project := "foo/bar"
	opt := &gitlab.ListProjectMergeRequestsOptions{}
	mr := &mocks.GitlabMergeRequestsService{}
	mr.On("ListProjectMergeRequests", project, opt).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{project: project, mr: mr, out: &bytes.Buffer{}}
	s.List(opt)

	mr.AssertExpectations(t)
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
	mr := &mocks.GitlabMergeRequestsService{}
	mr.On("GetMergeRequest", project, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &mergeRequestsService{project: project, mr: mr, out: &bytes.Buffer{}}
	err := s.Show(mrID)

	assert.Error(t, err)
	mr.AssertExpectations(t)
}
