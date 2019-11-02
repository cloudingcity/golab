package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if configured {
			return nil
		}

		return c.Edit(os.Stdin, os.Stdout)
	},
}

func init() {
	configCmd.AddCommand(configEditCmd)
}
