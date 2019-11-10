package cmd

import (
	"github.com/spf13/cobra"
)

var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Manage merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(mrCmd)
}
