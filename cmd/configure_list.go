package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configureListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration",
	Run: func(cmd *cobra.Command, args []string) {
		c.List(os.Stdout)
	},
}

func init() {
	configureCmd.AddCommand(configureListCmd)
}
