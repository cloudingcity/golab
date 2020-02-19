package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type flagError struct {
	Err error
}

func (e *flagError) Error() string {
	return e.Err.Error()
}

func (e *flagError) Unwrap() error {
	return e.Err
}

func handleError(cmd *cobra.Command, err error) {
	fmt.Println(err)

	var flagError *flagError
	if errors.As(err, &flagError) || isUnknownCommand(err) || isDefaultFlagError(err) {
		fmt.Println()
		fmt.Println(cmd.UsageString())
	}
}

func isUnknownCommand(err error) bool {
	return strings.HasPrefix(err.Error(), "unknown command ")
}

func isDefaultFlagError(err error) bool {
	return strings.HasPrefix(err.Error(), "required flag")
}
