package group

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	Search *searchService
	Depend *dependService
}

// NewManager returns a gitlab service manager.
func NewManager(c *gitlab.Client, group string, w io.Writer) *Manager {
	m := &Manager{}
	m.Search = &searchService{
		group:        group,
		gitlabSearch: c.Search,
		out:          w,
	}
	m.Depend = &dependService{
		group:          group,
		gitlabGroup:    c.Groups,
		gitlabRepoFile: c.RepositoryFiles,
		out:            w,
	}

	return m
}
