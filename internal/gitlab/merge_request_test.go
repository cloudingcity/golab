package gitlab

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

type stubMergeRequestService struct {
}

func (s *stubMergeRequestService) ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	return []*gitlab.MergeRequest{
		{IID: 1, Title: "Title 1"},
		{IID: 2, Title: "Title 2"},
	}, nil, nil
}

func TestMergeRequestList(t *testing.T) {
	s := &stubMergeRequestService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequest{mr: s, out: buf}

	mr.List("foo", nil)

	wants := []string{"#1", "#2", "Title 1", "Title 2"}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}
