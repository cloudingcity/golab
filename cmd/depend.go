package cmd

import (
	"github.com/spf13/cobra"
)

var dependCmd = &cobra.Command{
	Use:   "depend",
	Short: "Shows project which depend on a certain package",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(dependCmd)
}
