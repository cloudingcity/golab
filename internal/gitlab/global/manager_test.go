package global

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestNewManager(t *testing.T) {
	c := gitlab.NewClient(nil, "foo")
	c.SetBaseURL("bar")

	m := NewManager(c, nil)

	assert.NotNil(t, m.MergeRequest)
	assert.NotNil(t, m.Validate)
}
