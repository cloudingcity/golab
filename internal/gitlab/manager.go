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
	err := client.SetBaseURL(host)
	if err != nil {
		return nil, err
	}

	m := &Manager{}
	m.MergeRequest = &mergeRequest{client.MergeRequests, out}

	return m, nil
}
