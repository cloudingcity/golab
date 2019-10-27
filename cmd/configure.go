package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure golab CLI options",
	RunE: func(cmd *cobra.Command, args []string) error {
		if configured {
			return nil
		}

		return c.Configure(configPath, os.Stdin, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
