package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudingcity/golab/internal/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "golab",
	Short: "A CLI tool for gitlab",
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
	c          *config.Config
	configPath string
	configured = false
)

// initConfig reads in config file.
func initConfig() {
	home, _ := os.UserHomeDir()
	configPath = filepath.Join(home, ".config")

	c = config.New()

	if err := c.Load(configPath); err != nil {
		if err := c.Configure(configPath, os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
		configured = true
	}
}
