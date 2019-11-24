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
			State:       gitlab.String(mrListFlag.state),
			OrderBy:     gitlab.String("updated_at"),
			Scope:       gitlab.String(mrListFlag.optionScope()),
			ListOptions: mrListFlag.optionList(),
		}
		return projectManager().MergeRequest.List(opt, mrListFlag.url)
	},
}

type mrListFlagStruct struct {
	review bool
	state  string
	url    bool
	limit  int
}

func (f *mrListFlagStruct) optionScope() string {
	if f.review {
		return "assigned_to_me"
	}
	return "all"
}

func (f *mrListFlagStruct) optionList() gitlab.ListOptions {
	return gitlab.ListOptions{
		Page:    1,
		PerPage: f.limit,
	}
}

var mrListFlag *mrListFlagStruct

func init() {
	mrListFlag = &mrListFlagStruct{}
	mrListCmd.Flags().IntVarP(&mrListFlag.limit, "limit", "l", 20, "number of merge requests to list (max 100)")
	mrListCmd.Flags().BoolVarP(&mrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	mrListCmd.Flags().StringVarP(&mrListFlag.state, "state", "s", "opened", "filter merge requests by state (opened/closed/locked/merged)")
	mrListCmd.Flags().BoolVarP(&mrListFlag.url, "url", "u", false, "with url column")

	mrCmd.AddCommand(mrListCmd)
}
