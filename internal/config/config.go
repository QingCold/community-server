package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver string
	Dsn    string
}

type Config struct {
	Database DatabaseConfig
}

var Conf *Config

func LoadConfig(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	Conf = &Config{}
	if err := v.Unmarshal(Conf); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
}
