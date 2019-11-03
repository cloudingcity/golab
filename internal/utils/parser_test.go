package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseRepo(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{give: "git@gitlab.com:foo/bar.git", want: "foo/bar"},
		{give: "https://gitlab.com/foo/bar.git", want: "foo/bar"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ParseRepo(tt.give))
	}
}
