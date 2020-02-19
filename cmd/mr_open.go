package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var mrOpenCmd = &cobra.Command{
	Use:                   "open [MRID]",
	Short:                 "Open a merge request page in the default browser",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &flagError{errors.New("missing MRID")}
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return &flagError{fmt.Errorf("invalid MRID %q", args[0])}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		mrID, _ := strconv.Atoi(args[0])
		return projectManager(nil).MergeRequest.Open(mrID)
	},
}
