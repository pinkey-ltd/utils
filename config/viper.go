package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

type Config struct {
	raw interface{}
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("[u.config] Failed to read configuration file: ", err)
		return nil, err
	}
	cnf := viper.Get("config")
	return &Config{raw: cnf}, nil
}
