package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	conf "github.com/cloudingcity/golab/internal/config"
	"github.com/cloudingcity/golab/internal/gitlab/global"
	"github.com/cloudingcity/golab/internal/gitlab/project"
	"github.com/cloudingcity/golab/internal/utils"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
	"github.com/xanzy/go-gitlab"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "golab",
	Short:         "A CLI tool for gitlab",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cobra.OnInitialize(initConfig)
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

func currentProject() string {
	url, err := gitconfig.OriginURL()
	if err != nil {
		log.Fatal("not a git repository")
	}
	return utils.ParseGitProject(url)
}

func gitlabClient() *gitlab.Client {
	c := gitlab.NewClient(nil, config.Get("token"))
	if err := c.SetBaseURL(config.Get("host")); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func projectManager() *project.Manager {
	return project.NewManager(gitlabClient(), currentProject(), os.Stdout)
}

func globalManager() *global.Manager {
	return global.NewManager(gitlabClient(), os.Stdout)
}
