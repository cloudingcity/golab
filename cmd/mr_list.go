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
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}
		review, err := cmd.Flags().GetBool("review")
		if err != nil {
			return err
		}
		state, err := cmd.Flags().GetString("state")
		if err != nil {
			return err
		}

		opt := &gitlab.ListProjectMergeRequestsOptions{
			State:       gitlab.String(state),
			OrderBy:     gitlab.String("updated_at"),
			ListOptions: gitlab.ListOptions{Page: 1, PerPage: limit},
		}

		if review {
			opt.Scope = gitlab.String("assigned_to_me")
		}

		return projectManager(nil).MergeRequest.List(opt)
	},
}

func init() {
	mrListCmd.Flags().IntP("limit", "l", 20, "number of merge requests to list (max 100)")
	mrListCmd.Flags().BoolP("review", "r", false, "list merge requests assigned to you")
	mrListCmd.Flags().StringP("state", "s", "opened", "filter by state (opened/closed/locked/merged)")
}
