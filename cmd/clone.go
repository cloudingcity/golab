package cmd

import (
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [REPO] [--DIR]",
	Short: "Clone a repository from GitLab",
	Args:  cobra.MinimumNArgs(1),
	Run:   clone,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
