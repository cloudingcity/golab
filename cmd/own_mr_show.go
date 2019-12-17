package cmd

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ownMrShowCmd = &cobra.Command{
	Use:                   "show [PID] [MRID]",
	Short:                 "Show information about a merge request",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing PID")
		}
		if len(args) < 2 {
			return errors.New("missing MRID")
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return errors.Errorf("invalid PID %q", args[0])
		}
		if _, err := strconv.Atoi(args[1]); err != nil {
			return errors.Errorf("invalid MRID %q", args[1])
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pID := args[0]
		mrID, _ := strconv.Atoi(args[1])
		return globalManager().MergeRequest.Show(pID, mrID)
	},
}

func init() {
	ownMrCmd.AddCommand(ownMrShowCmd)
}
