package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Config{
		DatabaseURL: viper.GetString("database.url"),
	}, nil
}
