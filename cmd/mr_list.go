package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var mrListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := gitlabManager()
		opt := &gitlab.ListProjectMergeRequestsOptions{
			State:   gitlab.String("opened"),
			OrderBy: gitlab.String("updated_at"),
		}

		return m.MergeRequest.List(currentRepo(), opt)
	},
}

func init() {
	mrCmd.AddCommand(mrListCmd)
}
