package cmd

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var mrOpenCmd = &cobra.Command{
	Use:                   "open [MRID]",
	Short:                 "Open a merge request page in the default browser",
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
		return projectManager(nil).MergeRequest.Open(args[0])
	},
}

func init() {
	mrCmd.AddCommand(mrOpenCmd)
}
