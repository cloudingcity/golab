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
		opt := &gitlab.ListMergeRequestsOptions{
			State:       gitlab.String(ownMrListFlag.state),
			OrderBy:     gitlab.String("updated_at"),
			Scope:       gitlab.String(ownMrListFlag.optionScope()),
			ListOptions: ownMrListFlag.optionList(),
		}
		return globalManager().MergeRequest.List(opt, ownMrListFlag.url)
	},
}

type ownMrListFlagStruct struct {
	review bool
	state  string
	url    bool
	limit  int
}

func (f *ownMrListFlagStruct) optionScope() string {
	if f.review {
		return "assigned_to_me"
	}
	return "created_by_me"
}

func (f *ownMrListFlagStruct) optionList() gitlab.ListOptions {
	return gitlab.ListOptions{
		Page:    1,
		PerPage: f.limit,
	}
}

var ownMrListFlag *ownMrListFlagStruct

func init() {
	ownMrListFlag = &ownMrListFlagStruct{}
	ownMrListCmd.Flags().IntVarP(&ownMrListFlag.limit, "limit", "l", 20, "number of merge requests to list (max 100)")
	ownMrListCmd.Flags().BoolVarP(&ownMrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	ownMrListCmd.Flags().StringVarP(&ownMrListFlag.state, "state", "s", "opened", "filter merge requests by state (opened/closed/locked/merged)")
	ownMrListCmd.Flags().BoolVarP(&ownMrListFlag.url, "url", "u", false, "with url column")

	ownMrCmd.AddCommand(ownMrListCmd)
}
