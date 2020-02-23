package cmd

import (
	"github.com/spf13/cobra"
)

var mrCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a merge request",
	Long: `Create a merge request

1. Push local branch to remote repository
2. Opening merge request page in browser
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return projectManager(nil).MergeRequest.Create()
	},
}
