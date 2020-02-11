package cmd

import (
	"github.com/spf13/cobra"
)

var projectSearchCmd = &cobra.Command{
	Use:   "search [QUERY]",
	Short: "Search projects",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		q := args[0]
		if len(projectSearchFlag.group) != 0 {
			return groupManager(projectSearchFlag.group).Search.Project(q)
		}
		return globalManager().Search.Project(q)
	},
}

type projectSearchFlagStruct struct {
	group string
}

var projectSearchFlag *projectSearchFlagStruct

func init() {
	projectSearchFlag = &projectSearchFlagStruct{}
	projectSearchCmd.Flags().StringVarP(&projectSearchFlag.group, "group", "", "", "specify group to search")

	projectCmd.AddCommand(projectSearchCmd)
}
