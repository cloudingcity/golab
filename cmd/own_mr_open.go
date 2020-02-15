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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("missing arguments")
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
		return globalManager().MergeRequest.Open(pID, mrID)
	},
}
