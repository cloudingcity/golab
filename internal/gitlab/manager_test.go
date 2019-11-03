package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	t.Run("invalid host", func(t *testing.T) {
		_, err := NewManager("%", "", nil)

		assert.Error(t, err)
	})

	t.Run("member not nil", func(t *testing.T) {
		m, _ := NewManager("", "", nil)

		assert.NotNil(t, m.MergeRequest)
	})
}
