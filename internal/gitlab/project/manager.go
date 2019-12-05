package project

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	MergeRequest *mergeRequestsService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, project string, w io.Writer) *Manager {
	m := &Manager{}
	m.MergeRequest = &mergeRequestsService{
		project: project,
		mr:      c.MergeRequests,
		out:     w,
		baseURL: c.BaseURL(),
	}

	return m
}
