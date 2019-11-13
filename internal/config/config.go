package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudingcity/golab/internal/utils"
	"github.com/spf13/viper"
)

const (
	fileName  = "golab.yaml"
	tokenPath = "profile/personal_access_tokens"
)

// Config wraps Viper.
type Config struct {
	viper *viper.Viper
	path  string
}

// New returns an initialized Config instance.
func New(path string) *Config {
	v := viper.New()
	v.SetDefault("host", "https://gitlab.com")
	v.SetDefault("token", "None")

	return &Config{viper: v, path: path}
}

// Get returns the value associated with key.
func (c *Config) Get(key string) string {
	return c.viper.GetString(key)
}

// Load read config from the given path.
func (c *Config) Load() error {
	c.viper.SetConfigFile(filepath.Join(c.path, fileName))

	return c.viper.ReadInConfig()
}

// Init create a config file, update it if exists.
func (c *Config) Init(r io.Reader, w io.Writer) error {
	reader := bufio.NewReader(r)

	if host := c.readHost(w, reader); host != "" {
		c.viper.Set("host", host)
	}

	if token := c.readToken(w, reader); token != "" {
		c.viper.Set("token", token)
	}

	os.MkdirAll(c.path, os.ModePerm)
	c.viper.SetConfigFile(filepath.Join(c.path, fileName))
	if err := c.viper.WriteConfig(); err != nil {
		return err
	}

	fmt.Fprintf(w, "\nConfig saved to %s\n", c.viper.ConfigFileUsed())

	return nil
}

func (c *Config) readHost(w io.Writer, reader *bufio.Reader) string {
	fmt.Fprintf(w, "Gitlab Host [%s]: ", c.Get("host"))
	host, _ := reader.ReadString('\n')

	return strings.TrimSpace(host)
}

func (c *Config) readToken(w io.Writer, reader *bufio.Reader) string {
	tokenURL := c.Get("host") + "/" + tokenPath
	fmt.Fprintf(w, "Create a token here: %s\n", tokenURL)
	fmt.Fprintf(w, "Gitlab Token (scope: api) [%s]: ", c.Get("token"))
	token, _ := reader.ReadString('\n')

	return strings.TrimSpace(token)
}

// List show config content.
func (c *Config) List(w io.Writer) {
	h := []string{"name", "value"}
	rows := [][]string{
		{"host", c.Get("host")},
		{"token", c.Get("token")},
	}
	utils.RenderTable(w, h, rows)
}
