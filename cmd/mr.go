package cmd

import (
	"github.com/spf13/cobra"
)

var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Manage merge requests",
}

func init() {
	rootCmd.AddCommand(mrCmd)
	mrCmd.AddCommand(mrCreateCmd)
	mrCmd.AddCommand(mrListCmd)
	mrCmd.AddCommand(mrOpenCmd)
	mrCmd.AddCommand(mrSearchCmd)
	mrCmd.AddCommand(mrShowCmd)
}
