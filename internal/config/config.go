package config

import (
	"bufio"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var (
	v          *viper.Viper
	configPath string
)

// Load read config with path
func Load(path string) error {
	configPath = path

	v = viper.New()
	v.SetConfigName("golab")
	v.SetConfigType("yaml")
	v.SetDefault("host", "https://gitlab.com")
	v.SetDefault("token", "None")
	v.AddConfigPath(path)

	return v.ReadInConfig()
}

// Get returns the value associated with key from config
func Get(key string) string {
	return v.GetString(key)
}

// Configure configure and save file
func Configure(r io.Reader, w io.Writer) error {
	reader := bufio.NewReader(r)

	if host := readHost(w, reader); host != "" {
		v.Set("host", host)
	}

	if token := readToken(w, reader); token != "" {
		v.Set("token", token)
	}

	filePath := path.Join(configPath, "golab.yaml")
	if err := v.WriteConfigAs(filePath); err != nil {
		return err
	}

	fmt.Fprintf(w, "\nConfig saved to %s\n", filePath)

	return nil
}

func readHost(w io.Writer, reader *bufio.Reader) string {
	fmt.Fprintf(w, "Gitlab Host [%s]: ", Get("host"))
	host, _ := reader.ReadString('\n')

	return strings.TrimSpace(host)
}

func readToken(w io.Writer, reader *bufio.Reader) string {
	fmt.Fprintf(w, "Gitlab Token (scope: api) [%s]: ", Get("token"))
	token, _ := reader.ReadString('\n')

	return strings.TrimSpace(token)
}
