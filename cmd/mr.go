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
}
