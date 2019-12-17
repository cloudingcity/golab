package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGitProject(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{give: "git@gitlab.com:foo/bar.git", want: "foo/bar"},
		{give: "git@gitlab.com:foo/bar/baz.git", want: "foo/bar/baz"},
		{give: "https://gitlab.com/foo/bar.git", want: "foo/bar"},
		{give: "https://gitlab.com/foo/bar/baz.git", want: "foo/bar/baz"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ParseGitProject(tt.give))
	}
}

func TestParseMRProject(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{give: "https://gitlab.com/foo/bar/merge_requests/1", want: "foo/bar"},
		{give: "https://gitlab.com/foo/bar/baz/merge_requests/999", want: "foo/bar/baz"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ParseMRProject(tt.give))
	}
}
