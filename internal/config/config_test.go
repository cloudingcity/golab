package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPath = "../../test"
	fakePath = "/foo/bar/baz"
)

func TestGet(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		err := Load(testPath)

		assert.NoError(t, err)
		assert.Equal(t, Get("host"), "https://abc.com")
		assert.Equal(t, Get("token"), "NjEaWdDcARhzYKdx4fA4")
	})

	t.Run("load fail", func(t *testing.T) {
		err := Load(fakePath)

		assert.Error(t, err)
		assert.Equal(t, Get("host"), "https://gitlab.com")
		assert.Equal(t, Get("token"), "None")
	})
}
