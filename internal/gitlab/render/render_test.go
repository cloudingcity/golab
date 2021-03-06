package render

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestProjectMRs(t *testing.T) {
	buf := &bytes.Buffer{}
	mrs := []*gitlab.MergeRequest{
		{IID: 1, Title: "Foo 1", WebURL: "https://foo/1"},
		{IID: 2, Title: "Foo 2", WebURL: "https://foo/2"},
	}

	New(buf).ProjectMRs(mrs)

	wants := []string{
		"MRID", "TITLE", "URL",
		"1", "Foo 1", "https://foo/1",
		"2", "Foo 2", "https://foo/2",
	}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}

func TestGlobalMRs(t *testing.T) {
	buf := &bytes.Buffer{}
	mrs := []*gitlab.MergeRequest{
		{ProjectID: 100, IID: 1, Title: "Title 1", WebURL: "https://gitlab.com/foo/bar/merge_requests/1"},
		{ProjectID: 200, IID: 2, Title: "Title 2", WebURL: "https://gitlab.com/foo/bar/baz/merge_requests/999"},
	}

	New(buf).GlobalMRs(mrs)

	wants := []string{
		"PID", "MRID", "PROJECT", "TITLE", "URL",
		"100", "1", "foo/bar", "Title 1", "https://gitlab.com/foo/bar/merge_requests/1",
		"100", "2", "foo/bar/baz", "Title 2", "https://gitlab.com/foo/bar/baz/merge_requests/999",
	}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}

func TestMR(t *testing.T) {
	buf := &bytes.Buffer{}
	mr := &gitlab.MergeRequest{
		Title:        "How are you?",
		Description:  "I'm fine thank you! And you?",
		ProjectID:    100,
		IID:          5,
		SourceBranch: "staging",
		TargetBranch: "master",
		State:        "merged",
		Author:       &gitlab.BasicUser{Name: "Jax"},
		Assignee:     &gitlab.BasicUser{Name: "Nocture"},
		CreatedAt:    gitlab.Time(time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:    gitlab.Time(time.Date(2019, time.December, 31, 23, 59, 59, 0, time.UTC)),
		WebURL:       "https://gitlab.com/foo/bar/merge_requests/123",
	}

	New(buf).MR(mr)

	want := `
How are you?
--------------------------------------------------
I'm fine thank you! And you?
--------------------------------------------------
PID         100  
MRID        5
Project     foo/bar
Branch      staging -> master
State       merged
Author      Jax
Assignee    Nocture
CreatedAt   2019-12-31 00:00:00
UpdatedAt   2019-12-31 23:59:59
Url         https://gitlab.com/foo/bar/merge_requests/123
`
	assert.Equal(t, want, buf.String())
}

func TestLintCI(t *testing.T) {
	buf := &bytes.Buffer{}
	r := New(buf)

	t.Run("valid", func(t *testing.T) {
		defer buf.Reset()

		result := &gitlab.LintResult{Status: "valid", Errors: nil}
		r.LintCI(result)

		assert.Equal(t, "Valid!\n", buf.String())
	})

	t.Run("invalid", func(t *testing.T) {
		defer buf.Reset()

		result := &gitlab.LintResult{Status: "invalid", Errors: []string{"A error", "B error"}}
		r.LintCI(result)

		want := `Invalid!

Errors:
  - A error
  - B error
`
		assert.Equal(t, want, buf.String())
	})
}

func TestDepends(t *testing.T) {
	buf := &bytes.Buffer{}
	results := []*DependResult{
		{Project: "bar", Version: "v1.0.0", Branch: "master", URL: "https://gitlab.com/foo/bar"},
	}

	New(buf).Depends(results)

	wants := []string{
		"PROJECT", "VERSION", "BRANCH", "URL",
		"bar", "v1.0.0", "master", "https://gitlab.com/foo/bar",
	}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}

func TestProjects(t *testing.T) {
	buf := &bytes.Buffer{}
	projects := []*gitlab.Project{
		{ID: 100, NameWithNamespace: "pokemon/pikachu", WebURL: "https://gitlab.com/pokemon/pikachu"},
		{ID: 200, NameWithNamespace: "pokemon/pichu", WebURL: "https://gitlab.com/pokemon/pichu"},
	}

	New(buf).Projects(projects)

	wants := []string{
		"PID", "PROJECT", "URL",
		"100", "pokemon/pikachu", "https://gitlab.com/pokemon/pikachu",
		"200", "pokemon/pichu", "https://gitlab.com/pokemon/pichu",
	}
	got := buf.String()
	for _, want := range wants {
		assert.Contains(t, got, want)
	}
}
