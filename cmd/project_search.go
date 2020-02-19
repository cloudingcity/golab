package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var projectSearchCmd = &cobra.Command{
	Use:   "search [QUERY]",
	Short: "Search projects",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("requires a query")}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		group, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		q := args[0]

		if len(group) != 0 {
			return groupManager(group).Search.Project(q)
		}
		return globalManager().Search.Project(q)
	},
}

func init() {
	projectSearchCmd.Flags().StringP("group", "g", "", "filter by group")
}
