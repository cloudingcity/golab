package git

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	git := New(u)

	git.Clone("foo/bar", "my-dir")
	assert.Contains(t, git.String(), "git clone git@example.com:foo/bar.git my-dir")
}
