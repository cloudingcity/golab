package cmd

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ownMrOpenCmd = &cobra.Command{
	Use:                   "open [PID] [MRID]",
	Short:                 "Open a merge request page in the default browser",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("missing MRID")}
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return &flagError{fmt.Errorf("invalid MRID %q", args[0])}
		}
		if len(args) < 2 {
			return &flagError{errors.New("missing PID")}
		}
		if _, err := strconv.Atoi(args[1]); err != nil {
			return &flagError{fmt.Errorf("invalid PID %q", args[0])}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pID, _ := strconv.Atoi(args[0])
		mrID, _ := strconv.Atoi(args[1])
		return globalManager().MergeRequest.Open(pID, mrID)
	},
}
