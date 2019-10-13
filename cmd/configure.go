package cmd

import (
	"os"

	"github.com/cloudingcity/golab/internal/config"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure golab CLI options",
	Run: func(cmd *cobra.Command, args []string) {
		config.Configure(os.Stdin, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
