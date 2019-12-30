package group

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	Search *searchService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, group string, w io.Writer) *Manager {
	m := &Manager{}
	m.Search = &searchService{
		group:  group,
		search: c.Search,
		out:    w,
	}

	return m
}
