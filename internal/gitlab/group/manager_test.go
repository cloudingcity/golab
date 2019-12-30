package group

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestNewManager(t *testing.T) {
	c := gitlab.NewClient(nil, "foo")

	m := NewManager(c, "foo", nil)

	assert.NotNil(t, m.Search)
}
