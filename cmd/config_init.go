package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configInitCmd = &cobra.Command{
	Use:                   "init",
	Short:                 "Create a config file, update it if exists",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if configured {
			return nil
		}

		return config.Init(os.Stdin, os.Stdout)
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
}
