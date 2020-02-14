package cmd

import (
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Manage gitlab ci",
}

func init() {
	rootCmd.AddCommand(ciCmd)
}
