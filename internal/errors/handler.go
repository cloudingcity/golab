package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type FlagError struct {
	Err error
}

func (e *FlagError) Error() string {
	return e.Err.Error()
}

func (e *FlagError) Unwrap() error {
	return e.Err
}

func Handle(cmd *cobra.Command, err error) {
	fmt.Println(err)

	var flagError *FlagError
	if errors.As(err, &flagError) || isUnknownCommand(err) || isDefaultFlagError(err) {
		fmt.Println()
		fmt.Println(cmd.UsageString())
	}
}

func isUnknownCommand(err error) bool {
	return strings.HasPrefix(err.Error(), "unknown command")
}

func isDefaultFlagError(err error) bool {
	return strings.HasPrefix(err.Error(), "required flag")
}
