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
			State:   gitlab.String("opened"),
			OrderBy: gitlab.String("updated_at"),
			Scope:   gitlab.String(ownMrListFlag.optionScope()),
		}
		return globalManager().MergeRequest.List(opt, ownMrListFlag.url)
	},
}

type ownMrListFlagStruct struct {
	review bool
	url    bool
}

func (f *ownMrListFlagStruct) optionScope() string {
	if f.review {
		return "assigned_to_me"
	}
	return "created_by_me"
}

var ownMrListFlag *ownMrListFlagStruct

func init() {
	ownMrListFlag = &ownMrListFlagStruct{}
	ownMrListCmd.Flags().BoolVarP(&ownMrListFlag.review, "review", "r", false, "list merge requests assigned to you")
	ownMrListCmd.Flags().BoolVarP(&ownMrListFlag.url, "url", "u", false, "with url column")

	ownMrCmd.AddCommand(ownMrListCmd)
}
