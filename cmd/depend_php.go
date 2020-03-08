package cmd

import (
	"errors"

	errs "github.com/cloudingcity/golab/internal/errors"
	"github.com/spf13/cobra"
)

var dependPHPCmd = &cobra.Command{
	Use:   "php [PACKAGE] [--group]",
	Short: "List composer package (vendor/name) dependency",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &errs.FlagError{Err: errors.New("requires a package")}
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		group, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		return groupManager(group).Depend.PHP(args[0])
	},
}

func init() {
	dependPHPCmd.Flags().StringP("group", "g", "", "group to inspect")
	dependPHPCmd.MarkFlagRequired("group")
}
