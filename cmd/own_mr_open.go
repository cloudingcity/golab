package cmd

import (
	"github.com/spf13/cobra"
)

var ownMrOpenCmd = &cobra.Command{
	Use:                   "open [PROJECT-ID] [MR-ID]",
	Short:                 "Open a merge request page in the default browser",
	Args:                  cobra.MinimumNArgs(2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return globalManager().MergeRequest.Open(args[0], args[1])
	},
}

func init() {
	ownMrCmd.AddCommand(ownMrOpenCmd)
}
