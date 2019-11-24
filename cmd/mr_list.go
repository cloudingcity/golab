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
		opt := &gitlab.ListProjectMergeRequestsOptions{
			State:   gitlab.String("opened"),
			OrderBy: gitlab.String("updated_at"),
			Scope:   projectScope(),
		}
		return projectManager().MergeRequest.List(opt, withURL)
	},
}

func init() {
	mrListCmd.Flags().BoolVarP(&review, "review", "r", false, "list merge requests assigned to you")
	mrListCmd.Flags().BoolVarP(&withURL, "url", "u", false, "with url column")

	mrCmd.AddCommand(mrListCmd)
}

var (
	review  bool
	withURL bool
)

func projectScope() *string {
	if review {
		return gitlab.String("assigned_to_me")
	}
	return gitlab.String("all")
}
