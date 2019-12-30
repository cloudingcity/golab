package global

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cloudingcity/golab/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestList(t *testing.T) {
	opt := &gitlab.ListMergeRequestsOptions{}
	mr := &mocks.GitlabMergeRequests{}
	mr.On("ListMergeRequests", opt).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{mr: mr, out: &bytes.Buffer{}}
	s.List(opt)

	mr.AssertExpectations(t)
}

func TestOpen(t *testing.T) {
	var got string
	pID := "123"
	mrID := 456
	mr := &mocks.GitlabMergeRequests{}
	mr.On("GetMergeRequest", pID, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{WebURL: "https://foo/bar"}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{
		mr: mr,
		openURL: func(url string) error {
			got = url
			return nil
		},
	}
	s.Open(pID, mrID)

	assert.Equal(t, "https://foo/bar", got)
	mr.AssertExpectations(t)
}

func TestShow(t *testing.T) {
	pID := "123"
	mrID := 456
	mr := &mocks.GitlabMergeRequests{}
	mr.On("GetMergeRequest", pID, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &mergeRequestsService{mr: mr, out: &bytes.Buffer{}}
	err := s.Show(pID, mrID)

	assert.Error(t, err)
	mr.AssertExpectations(t)
}
