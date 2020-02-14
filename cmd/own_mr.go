package cmd

import (
	"github.com/spf13/cobra"
)

var ownMrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Manage own merge requests",
}

func init() {
	ownCmd.AddCommand(ownMrCmd)
}
