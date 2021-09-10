package config

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	ModeEnv string
	DBName  string
}

func New() *Config {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(err)
	}
	var c Config
	c.Port = viper.GetString("port")
	c.ModeEnv = viper.GetString("mode")
	c.DBName = viper.GetString("dbname")
	return &c
}

// any approach to require this configuration into your program.
var yamlExample = []byte(`
mode: testing
port: 8080
dbname: mylittle.db
`)
