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
		return globalManager().MergeRequest.List(opt, withURL)
	},
}

func init() {
	ownMrListCmd.Flags().BoolVarP(&review, "review", "r", false, "list merge requests assigned to you")
	ownMrListCmd.Flags().BoolVarP(&withURL, "url", "u", false, "with url column")

	ownMrCmd.AddCommand(ownMrListCmd)
}
