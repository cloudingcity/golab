package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var mrOpenCmd = &cobra.Command{
	Use:                   "open [MRID]",
	Short:                 "Open a merge request page in the default browser",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mrID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return projectManager(nil).MergeRequest.Open(mrID)
	},
}
