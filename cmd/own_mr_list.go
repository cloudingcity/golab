package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var ownMrListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List merge requests created by you",
	RunE: func(cmd *cobra.Command, args []string) error {
		return globalManager().MergeRequest.List(ownMrListFlag.option())
	},
}

type ownMrListFlagStruct struct {
	review bool
	state  string
	limit  int
}

func (f *ownMrListFlagStruct) option() *gitlab.ListMergeRequestsOptions {
	opt := &gitlab.ListMergeRequestsOptions{
		State:       gitlab.String(ownMrListFlag.state),
		OrderBy:     gitlab.String("updated_at"),
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: f.limit},
	}

	if f.review {
		opt.Scope = gitlab.String("assigned_to_me")
	} else {
		opt.Scope = gitlab.String("created_by_me")
	}

	return opt
}

var ownMrListFlag *ownMrListFlagStruct

func init() {
	ownMrListFlag = &ownMrListFlagStruct{}
	ownMrListCmd.Flags().IntVarP(&ownMrListFlag.limit, "limit", "l", 20, "number of merge requests to list (max 100)")
	ownMrListCmd.Flags().BoolVarP(&ownMrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	ownMrListCmd.Flags().StringVarP(&ownMrListFlag.state, "state", "s", "opened", "filter merge requests by state (opened/closed/locked/merged)")

	ownMrCmd.AddCommand(ownMrListCmd)
}
