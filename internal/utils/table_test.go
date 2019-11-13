package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTable(t *testing.T) {
	buf := &bytes.Buffer{}
	h := []string{"title", "content"}
	rows := [][]string{
		{"Day 1", "Go to school"},
		{"Day 2", "Go to shop"},
	}
	RenderTable(buf, h, rows)

	got := buf.String()
	assert.Contains(t, got, "TITLE")
	assert.Contains(t, got, "CONTENT")
	assert.Contains(t, got, "Day 1")
	assert.Contains(t, got, "Day 2")
	assert.Contains(t, got, "Go to school")
	assert.Contains(t, got, "Go to shop")
}
