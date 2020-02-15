package cmd

import (
	"github.com/spf13/cobra"
)

var dependGOCmd = &cobra.Command{
	Use:                   "go [PACKAGE] [--group]",
	Short:                 "List go module package (example.com/hello) dependency",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		group, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		return groupManager(group).Depend.GO(args[0])
	},
}

func init() {
	dependGOCmd.Flags().StringP("group", "g", "", "group to inspect")
	dependGOCmd.MarkFlagRequired("group")
}
