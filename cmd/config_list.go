package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all configuration",
	Run: func(cmd *cobra.Command, args []string) {
		c.List(os.Stdout)
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
