package cmd

import (
	"github.com/spf13/cobra"
)

var dependCmd = &cobra.Command{
	Use:   "depend",
	Short: "Shows project which depend on a certain package",
}

func init() {
	rootCmd.AddCommand(dependCmd)
}
