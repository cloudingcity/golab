package cmd

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	conf "github.com/cloudingcity/golab/internal/config"
	"github.com/cloudingcity/golab/internal/errors"
	"github.com/cloudingcity/golab/internal/git"
	"github.com/cloudingcity/golab/internal/gitlab/global"
	"github.com/cloudingcity/golab/internal/gitlab/group"
	"github.com/cloudingcity/golab/internal/gitlab/project"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "golab",
	Short:         "A CLI tool for gitlab",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if cmd, err := rootCmd.ExecuteC(); err != nil {
		errors.Handle(cmd, err)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return &errors.FlagError{Err: err}
	})
	cobra.OnInitialize(initConfig)

	log.SetFlags(0)
}

var (
	config     *conf.Config
	configured = false
)

// initConfig reads in config file.
func initConfig() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config")
	config = conf.New(path)

	if err := config.Load(); err != nil {
		if err := config.Init(os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
		configured = true
	}
}

func gitlabClient() *gitlab.Client {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	c := gitlab.NewClient(client, config.Get("token"))
	if err := c.SetBaseURL(config.Get("host")); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func projectManager(p *string) *project.Manager {
	if p == nil {
		temp := git.CurrentRepo()
		p = &temp
	}

	return project.NewManager(gitlabClient(), *p, os.Stdout)
}

func groupManager(g string) *group.Manager {
	return group.NewManager(gitlabClient(), g, os.Stdout)
}

func globalManager() *global.Manager {
	return global.NewManager(gitlabClient(), os.Stdout)
}
