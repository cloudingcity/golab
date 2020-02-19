package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var mrSearchCmd = &cobra.Command{
	Use:   "search [QUERY]",
	Short: "Search merge requests",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("requires a query")}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		global, err := cmd.Flags().GetBool("global")
		if err != nil {
			return err
		}
		group, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		q := args[0]

		if global {
			return globalManager().Search.MR(q)
		}
		if len(group) == 0 {
			return projectManager(nil).Search.MR(q)
		}
		if len(project) == 0 {
			return groupManager(group).Search.MR(q)
		}
		p := group + "/" + project
		return projectManager(&p).Search.MR(q)
	},
}

func init() {
	mrSearchCmd.Flags().BoolP("global", "G", false, "search everywhere")
	mrSearchCmd.Flags().StringP("group", "g", "", "filter by group")
	mrSearchCmd.Flags().StringP("project", "p", "", "filter by project")
}
