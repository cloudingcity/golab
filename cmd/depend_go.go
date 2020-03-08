package cmd

import (
	"errors"

	errs "github.com/cloudingcity/golab/internal/errors"
	"github.com/spf13/cobra"
)

var dependGOCmd = &cobra.Command{
	Use:   "go [PACKAGE] [--group]",
	Short: "List go module package (example.com/hello) dependency",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &errs.ArgError{Err: errors.New("requires a package")}
		}
		return nil
	},
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
