package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:                   "list",
	Aliases:               []string{"ls"},
	Short:                 "List all configuration",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		config.List(os.Stdout)
	},
}
