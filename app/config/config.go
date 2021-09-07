package config

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	ModeEnv    string
	JWTSignKey string
}

func New() *Config {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(err)
	}
	var config Config
	config.Port = viper.GetString("port")
	config.ModeEnv = viper.GetString("mode")
	config.JWTSignKey = viper.GetString("jwtSignKey")
	return &config
}

// any approach to require this configuration into your program.
var yamlExample = []byte(`
mode: production
port: 8081
jwtSignKey: secret
`)
