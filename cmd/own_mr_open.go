package cmd

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ownMrOpenCmd = &cobra.Command{
	Use:                   "open [PID] [MRID]",
	Short:                 "Open a merge request page in the default browser",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pID, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.Errorf("invalid PID %q", args[0])
		}
		mrID, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.Errorf("invalid MRID %q", args[1])
		}
		return globalManager().MergeRequest.Open(pID, mrID)
	},
}
