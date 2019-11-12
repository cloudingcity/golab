package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var ownMrListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List merge requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &gitlab.ListMergeRequestsOptions{
			State:   gitlab.String("opened"),
			OrderBy: gitlab.String("updated_at"),
			Scope:   scope(),
		}
		return groupManager().MergeRequest.List(opt)
	},
}

func init() {
	ownMrListCmd.Flags().BoolVar(&review, "review", false, "List merge requests assigned to you")
	ownMrCmd.AddCommand(ownMrListCmd)
}

var review bool

func scope() *string {
	scope := "created_by_me"

	if review {
		scope = "assigned_to_me"
	}

	return gitlab.String(scope)
}
