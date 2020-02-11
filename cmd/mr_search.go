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
		if mrSearchFlag.global {
			return globalManager().Search.MR(q)
		}
		if len(mrSearchFlag.group) == 0 {
			return projectManager(nil).Search.MR(q)
		}
		if len(mrSearchFlag.project) == 0 {
			return groupManager(mrSearchFlag.group).Search.MR(q)
		}
		p := mrSearchFlag.group + "/" + mrSearchFlag.project
		return projectManager(&p).Search.MR(q)
	},
}

type mrSearchFlagStruct struct {
	group   string
	project string
	global  bool
}

var mrSearchFlag *mrSearchFlagStruct

func init() {
	mrSearchFlag = &mrSearchFlagStruct{}
	mrSearchCmd.Flags().BoolVarP(&mrSearchFlag.global, "global", "g", false, "search all places")
	mrSearchCmd.Flags().StringVarP(&mrSearchFlag.group, "group", "", "", "specify group to search")
	mrSearchCmd.Flags().StringVarP(&mrSearchFlag.project, "project", "", "", "specify project to search")

	mrCmd.AddCommand(mrSearchCmd)
}
