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
	mockGitlabMR := &mocks.GitlabMergeRequests{}
	mockGitlabMR.On("ListMergeRequests", opt).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{gitlabMR: mockGitlabMR, out: &bytes.Buffer{}}
	s.List(opt)

	mockGitlabMR.AssertExpectations(t)
}

func TestOpen(t *testing.T) {
	var openURL string
	stdout := &bytes.Buffer{}

	pID := 123
	mrID := 456
	mockGitlabMR := &mocks.GitlabMergeRequests{}
	mockGitlabMR.On("GetMergeRequest", pID, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{WebURL: "https://foo/bar"}, &gitlab.Response{}, nil)

	s := &mergeRequestsService{
		gitlabMR: mockGitlabMR,
		out:      stdout,
		openURL: func(url string) error {
			openURL = url
			return nil
		},
	}
	s.Open(pID, mrID)

	assert.Equal(t, "Opening https://foo/bar in your browser\n", stdout.String())
	assert.Equal(t, "https://foo/bar", openURL)
	mockGitlabMR.AssertExpectations(t)
}

func TestShow(t *testing.T) {
	pID := "123"
	mrID := 456
	mockGitlabMR := &mocks.GitlabMergeRequests{}
	mockGitlabMR.On("GetMergeRequest", pID, mrID, (*gitlab.GetMergeRequestsOptions)(nil)).
		Once().
		Return(&gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &mergeRequestsService{gitlabMR: mockGitlabMR, out: &bytes.Buffer{}}
	err := s.Show(pID, mrID)

	assert.Error(t, err)
	mockGitlabMR.AssertExpectations(t)
}
