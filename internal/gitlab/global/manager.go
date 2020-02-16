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
	Search       *searchService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, w io.Writer) *Manager {
	m := &Manager{}
	m.MergeRequest = &mergeRequestsService{
		gitlabMR: c.MergeRequests,
		out:      w,
		openURL:  browser.OpenURL,
	}
	m.Validate = &validateService{
		gitlabValidate: c.Validate,
		out:            w,
	}
	m.Search = &searchService{
		gitlabSearch: c.Search,
		out:          w,
	}

	return m
}
