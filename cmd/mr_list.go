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
			Scope:   gitlab.String(mrListFlag.optionScope()),
		}
		return projectManager().MergeRequest.List(opt, mrListFlag.url)
	},
}

type mrListFlagStruct struct {
	review bool
	url    bool
}

func (f *mrListFlagStruct) optionScope() string {
	if f.review {
		return "assigned_to_me"
	}
	return "all"
}

var mrListFlag *mrListFlagStruct

func init() {
	mrListFlag = &mrListFlagStruct{}
	mrListCmd.Flags().BoolVarP(&mrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	mrListCmd.Flags().BoolVarP(&mrListFlag.url, "url", "u", false, "with url column")

	mrCmd.AddCommand(mrListCmd)
}
