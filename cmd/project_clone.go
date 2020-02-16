package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/cloudingcity/golab/internal/git"
	"github.com/spf13/cobra"
)

var clone = func(cmd *cobra.Command, args []string) {
	var dir, project string

	project = args[0]
	if len(args) > 1 {
		dir = args[1]
	}

	host, err := url.Parse(config.Get("host"))
	if err != nil {
		log.Fatal(err)
	}

	repo := fmt.Sprintf("git@%s:%s.git", host.Host, project)

	git.Clone(repo, dir)
}

var projectCloneCmd = &cobra.Command{
	Use:   "clone [REPO] [--DIR]",
	Short: "Clone a repository from GitLab",
	Args:  cobra.MinimumNArgs(1),
	Run:   clone,
}
