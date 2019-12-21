package cmd

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var mrShowCmd = &cobra.Command{
	Use:                   "show [MRID]",
	Short:                 "Show information about a merge request",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing MRID")
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return errors.Errorf("invalid MRID %q", args[0])
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		mrID, _ := strconv.Atoi(args[0])
		return projectManager().MergeRequest.Show(mrID)
	},
}

func init() {
	mrCmd.AddCommand(mrShowCmd)
}