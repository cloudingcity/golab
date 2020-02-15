package cmd

import (
	"github.com/spf13/cobra"
)

var ciLintCmd = &cobra.Command{
	Use:                   "lint [FILE]",
	Short:                 "Validate the .gitlab-ci.yml",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return globalManager().Validate.Lint(args[0])
	},
}
