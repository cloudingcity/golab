package errors

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestIsUnknownCommand(t *testing.T) {
	err := errors.New("unknown command")

	assert.True(t, isUnknownCommand(err))
}

func TestIsDefaultFlagError(t *testing.T) {
	err := errors.New("required flag")

	assert.True(t, isDefaultFlagError(err))
}

func ExampleHandle() {
	cmd := &cobra.Command{}
	err := &FlagError{Err: errors.New("something wrong")}

	Handle(cmd, err)

	// Output: something wrong
	//
	// Usage:
}
