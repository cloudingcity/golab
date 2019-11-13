package global

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

type stubMergeRequestsService struct {
}

func (s *stubMergeRequestsService) ListMergeRequests(opt *gitlab.ListMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	return []*gitlab.MergeRequest{
		{ProjectID: 100, IID: 1, Title: "Title 1", WebURL: "https://gitlab.com/foo/bar/merge_requests/1"},
		{ProjectID: 200, IID: 2, Title: "Title 2", WebURL: "https://gitlab.com/foo/bar/baz/merge_requests/999"},
	}, nil, nil
}

func (s *stubMergeRequestsService) GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	return nil, nil, errors.New("error")
}

func TestMergeRequestsServiceList(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	mr.List(nil)

	wants := []string{"ID", "PROJECT", "TITLE", "100 1", "200 2", "foo/bar", "foo/bar/baz", "Title 1", "Title 2"}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}

func TestMergeRequestsOpen(t *testing.T) {
	s := &stubMergeRequestsService{}
	mr := &mergeRequestsService{mr: s}

	t.Run("invalid project id", func(t *testing.T) {
		err := mr.Open("error project id", "123")

		assert.Error(t, err)
	})

	t.Run("invalid merge request id", func(t *testing.T) {
		err := mr.Open("123", "error mr id")

		assert.Error(t, err)
	})

	t.Run("get merge requests error", func(t *testing.T) {
		err := mr.Open("123", "123")

		assert.Error(t, err)
	})
}
