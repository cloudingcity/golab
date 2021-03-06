package cmd

import (
	"github.com/spf13/cobra"
)

var ownCmd = &cobra.Command{
	Use:   "own",
	Short: "Manage own resources",
}

func init() {
	rootCmd.AddCommand(ownCmd)
	ownCmd.AddCommand(ownMrCmd)
	ownMrCmd.AddCommand(ownMrListCmd)
	ownMrCmd.AddCommand(ownMrOpenCmd)
	ownMrCmd.AddCommand(ownMrShowCmd)
}
