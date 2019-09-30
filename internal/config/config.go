package config

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

// Load read config with path
func Load(path string) error {
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
