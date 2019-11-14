package project

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

type stubMergeRequestsService struct {
}

func (s *stubMergeRequestsService) ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	return []*gitlab.MergeRequest{
		{IID: 1, Title: "Title 1"},
		{IID: 2, Title: "Title 2"},
	}, nil, nil
}

func TestMergeRequestList(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	mr.List(nil)

	wants := []string{"MRID", "TITLE", "1", "2", "Title 1", "Title 2"}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}

func TestMergeRequestOpen(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		mr := &mergeRequestsService{}
		err := mr.Open("aaa")

		assert.Error(t, err)
	})
}
