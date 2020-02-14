package cmd

import (
	"github.com/spf13/cobra"
)

var dependPHPCmd = &cobra.Command{
	Use:                   "php [PACKAGE] [--group]",
	Short:                 "List composer package (vendor/name) dependency",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return groupManager(dependPHPFlag.group).Depend.PHP(args[0])
	},
}

type dependPHPFlagStruct struct {
	group string
}

var dependPHPFlag *dependPHPFlagStruct

func init() {
	dependPHPFlag = &dependPHPFlagStruct{}
	dependPHPCmd.Flags().StringVarP(&dependPHPFlag.group, "group", "", "", "group to inspect")
	dependPHPCmd.MarkFlagRequired("group")
	dependPHPCmd.MarkFlagRequired("pkg")
}
