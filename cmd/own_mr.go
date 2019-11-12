package cmd

import (
	"github.com/spf13/cobra"
)

var ownMrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Manage own merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	ownCmd.AddCommand(ownMrCmd)
}
