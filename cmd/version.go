package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version means app version
var Version string

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print version number of golab",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("golab version " + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
