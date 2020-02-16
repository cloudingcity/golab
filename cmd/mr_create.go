package cmd

import (
	"github.com/spf13/cobra"
)

var mrCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a merge request",
	RunE: func(cmd *cobra.Command, args []string) error {
		return projectManager(nil).MergeRequest.Create()
	},
}
