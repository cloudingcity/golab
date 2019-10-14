package cmd

import (
	"os"

	"github.com/cloudingcity/golab/internal/config"
	"github.com/spf13/cobra"
)

// configureListCmd represents the configureList command
var configureListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config.List(os.Stdout)
	},
}

func init() {
	configureCmd.AddCommand(configureListCmd)
}
