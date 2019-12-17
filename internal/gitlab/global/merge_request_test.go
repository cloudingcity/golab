package global

import (
	"bytes"
	"errors"
	"testing"
	"time"

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
	if pid == "test-show" {
		return &gitlab.MergeRequest{
			Author:    &gitlab.BasicUser{},
			Assignee:  &gitlab.BasicUser{},
			CreatedAt: gitlab.Time(time.Now()),
			UpdatedAt: gitlab.Time(time.Now()),
			WebURL:    "https://gitlab.com/foo/bar/merge_requests/123",
		}, nil, nil
	}

	return nil, nil, errors.New("error")
}

func TestMergeRequestsServiceList(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	t.Run("list", func(t *testing.T) {
		mr.List(nil, false)

		wants := []string{"PID", "MRID", "PROJECT", "TITLE", "100", "1", "200", "2", "foo/bar", "foo/bar/baz", "Title 1", "Title 2"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})

	t.Run("list with url", func(t *testing.T) {
		mr.List(nil, true)

		wants := []string{"PID", "MRID", "PROJECT", "TITLE", "100", "1", "200", "2", "foo/bar", "foo/bar/baz", "Title 1", "Title 2",
			"https://gitlab.com/foo/bar/merge_requests/1", "https://gitlab.com/foo/bar/baz/merge_requests/999",
		}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})
}

func TestMergeRequestShow(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	t.Run("show", func(t *testing.T) {
		mr.Show("test-show", 123)

		wants := []string{"PID", "MRID", "Project", "Branch", "State", "Author", "Assignee", "CreatedAt", "UpdatedAt"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})
}
