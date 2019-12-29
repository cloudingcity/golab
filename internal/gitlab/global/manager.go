package global

import (
	"io"

	"github.com/pkg/browser"
	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	MergeRequest *mergeRequestsService
	Validate     *validateService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, w io.Writer) *Manager {
	m := &Manager{}
	m.MergeRequest = &mergeRequestsService{
		mr:      c.MergeRequests,
		out:     w,
		openURL: browser.OpenURL,
	}
	m.Validate = &validateService{
		validate: c.Validate,
		out:      w,
	}

	return m
}
