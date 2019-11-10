package gitlab

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

// Manager manages gitlab services.
type Manager struct {
	MergeRequest *mergeRequest
}

// NewManager returns a gitlab service manager.
func NewManager(host, token string, out io.Writer) (*Manager, error) {
	client := gitlab.NewClient(nil, token)
	if err := client.SetBaseURL(host); err != nil {
		return nil, err
	}

	m := &Manager{}
	m.MergeRequest = &mergeRequest{url: client.BaseURL(), mr: client.MergeRequests, out: out}

	return m, nil
}
