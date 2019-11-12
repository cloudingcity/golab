package cmd

import (
	"github.com/spf13/cobra"
)

var ownCmd = &cobra.Command{
	Use:   "own",
	Short: "Manage own resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(ownCmd)
}
