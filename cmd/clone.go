package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [REPO] [--DIR]",
	Short: "Clone a repository from GitLab",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("requires a repository")}
		}
		return nil
	},
	RunE: clone,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
