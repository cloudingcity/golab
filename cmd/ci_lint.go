package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var ciLintCmd = &cobra.Command{
	Use:   "lint [FILE]",
	Short: "Validate the .gitlab-ci.yml",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("requires a file")}
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return globalManager().Validate.Lint(args[0])
	},
}
