package cmd

import (
	"errors"

	errs "github.com/cloudingcity/golab/internal/errors"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [REPO] [--DIR]",
	Short: "Clone a repository from GitLab",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &errs.FlagError{Err: errors.New("requires a repository")}
		}
		return nil
	},
	RunE: clone,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
