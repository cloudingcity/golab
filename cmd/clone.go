package cmd

import (
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [REPO] [--DIR]",
	Short: "Clone a repository from GitLab",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string

		repo := args[0]
		if len(args) > 1 {
			dir = args[1]
		}

		gitCmd().Clone(repo, dir).Run()
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
