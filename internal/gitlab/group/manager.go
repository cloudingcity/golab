package group

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	MergeRequest *mergeRequestsService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, w io.Writer) *Manager {
	m := &Manager{}
	m.MergeRequest = &mergeRequestsService{
		mr:  c.MergeRequests,
		out: w,
	}

	return m
}
