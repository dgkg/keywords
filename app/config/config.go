package config

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	ModeEnv string
}

func New() *Config {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	var config Config
	config.Port = viper.GetString("port")
	config.ModeEnv = viper.GetString("mode")
	return &config
}

// any approach to require this configuration into your program.
var yamlExample = []byte(`
mode: production
port: 8081
`)
