package cmd

import (
	"github.com/spf13/cobra"
)

var dependGOCmd = &cobra.Command{
	Use:                   "go [PKG] [--group]",
	Short:                 "List go module package (example.com/hello) dependency",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return groupManager(dependGOFlag.group).Depend.GO(args[0])
	},
}

type dependGOFlagStruct struct {
	group string
}

var dependGOFlag *dependGOFlagStruct

func init() {
	dependGOFlag = &dependGOFlagStruct{}
	dependGOCmd.Flags().StringVarP(&dependGOFlag.group, "group", "g", "", "group to inspect")
	dependGOCmd.MarkFlagRequired("group")

	dependCmd.AddCommand(dependGOCmd)
}
