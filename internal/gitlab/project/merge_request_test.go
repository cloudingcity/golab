package project

import (
	"bytes"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

type stubMergeRequestsService struct {
}

func (s *stubMergeRequestsService) ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	return []*gitlab.MergeRequest{
		{IID: 1, Title: "Title 1", WebURL: "https://foo/1"},
		{IID: 2, Title: "Title 2", WebURL: "https://foo/2"},
	}, nil, nil
}

func (s *stubMergeRequestsService) GetMergeRequest(pid interface{}, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	return &gitlab.MergeRequest{
		Author:    &gitlab.BasicUser{},
		Assignee:  &gitlab.BasicUser{},
		CreatedAt: gitlab.Time(time.Now()),
		UpdatedAt: gitlab.Time(time.Now()),
	}, nil, nil
}

func TestMergeRequestList(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	t.Run("list", func(t *testing.T) {
		mr.List(nil, false)

		wants := []string{"MRID", "TITLE", "1", "2", "Title 1", "Title 2"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})

	t.Run("list with url", func(t *testing.T) {
		mr.List(nil, true)

		wants := []string{"MRID", "TITLE", "URL", "1", "2", "Title 1", "Title 2", "https://foo/1", "https://foo/2"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})
}

func TestMRURL(t *testing.T) {
	u, _ := url.Parse("https://foo.com")
	mr := &mergeRequestsService{
		project: "my-project",
		baseURL: u,
	}

	want := "https://foo.com/my-project/merge_requests/123"
	got := mr.mrURL("123")

	assert.Equal(t, want, got)
}

func TestMergeRequestShow(t *testing.T) {
	s := &stubMergeRequestsService{}
	buf := &bytes.Buffer{}
	mr := &mergeRequestsService{mr: s, out: buf}

	t.Run("show", func(t *testing.T) {
		mr.Show(123)

		wants := []string{"PID", "MRID", "Project", "Branch", "State", "Author", "Assignee", "CreatedAt", "UpdatedAt"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})
}
