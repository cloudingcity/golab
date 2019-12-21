package global

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestRenderLint(t *testing.T) {
	buf := &bytes.Buffer{}
	validate := &validateService{out: buf}

	t.Run("valid", func(t *testing.T) {
		defer buf.Reset()

		result := &gitlab.LintResult{Status: "valid", Errors: nil}
		validate.renderLint(result)

		want := "Valid!\n"
		got := buf.String()

		assert.Equal(t, want, got)
	})

	t.Run("invalid", func(t *testing.T) {
		defer buf.Reset()

		result := &gitlab.LintResult{Status: "invalid", Errors: []string{"A error", "B error"}}
		validate.renderLint(result)

		wants := []string{"Invalid!", "Errors", "- A error", "- B error"}
		got := buf.String()
		for _, want := range wants {
			assert.Contains(t, got, want)
		}
	})
}
