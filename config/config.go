package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Listen   string `json:"listen"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadConfig() (config *Config, err error) {

	viper.SetDefault("listen", "0.0.0.0")
	viper.SetDefault("port", 8080)
	viper.SetDefault("username", "admin")
	viper.SetDefault("password", "admin123456")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	_ = viper.SafeWriteConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
