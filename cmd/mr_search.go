package cmd

import (
	"github.com/spf13/cobra"
)

var mrSearchCmd = &cobra.Command{
	Use:   "search [QUERY]",
	Short: "Search merge requests",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		q := args[0]
		if mrSearchFlag.isGlobal {
			return globalManager().Search.MR(q)
		}
		if len(mrSearchFlag.group) != 0 {
			if len(mrSearchFlag.project) != 0 {
				p := mrSearchFlag.group + "/" + mrSearchFlag.project
				return projectManager(&p).Search.MR(q)
			}
			return groupManager(mrSearchFlag.group).Search.MR(q)
		}
		return projectManager(nil).Search.MR(q)
	},
}

type mrSearchFlagStruct struct {
	group    string
	project  string
	isGlobal bool
}

var mrSearchFlag *mrSearchFlagStruct

func init() {
	mrSearchFlag = &mrSearchFlagStruct{}
	mrSearchCmd.Flags().StringVarP(&mrSearchFlag.group, "group", "g", "", "specify group to search")
	mrSearchCmd.Flags().StringVarP(&mrSearchFlag.project, "project", "p", "", "specify project to search")
	mrSearchCmd.Flags().BoolVarP(&mrSearchFlag.isGlobal, "global", "", false, "search all places")

	mrCmd.AddCommand(mrSearchCmd)
}
