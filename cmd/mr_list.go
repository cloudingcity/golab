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
		return projectManager(nil).MergeRequest.List(mrListFlag.option())
	},
}

type mrListFlagStruct struct {
	review bool
	state  string
	limit  int
}

func (f *mrListFlagStruct) option() *gitlab.ListProjectMergeRequestsOptions {
	opt := &gitlab.ListProjectMergeRequestsOptions{
		State:       gitlab.String(mrListFlag.state),
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

var mrListFlag *mrListFlagStruct

func init() {
	mrListFlag = &mrListFlagStruct{}
	mrListCmd.Flags().IntVarP(&mrListFlag.limit, "limit", "l", 20, "number of merge requests to list (max 100)")
	mrListCmd.Flags().BoolVarP(&mrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	mrListCmd.Flags().StringVarP(&mrListFlag.state, "state", "s", "opened", "filter merge requests by state (opened/closed/locked/merged)")

	mrCmd.AddCommand(mrListCmd)
}
