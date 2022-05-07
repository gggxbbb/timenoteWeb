package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	. "timenoteWeb/logger"
)

type ServerConfig struct {
	Listen string `json:"listen"`
	Port   int    `json:"port"`
}

type AdminConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WebConfig struct {
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
}

type Config struct {
	Server ServerConfig `json:"server"`
	Admin  AdminConfig  `json:"admin"`
	Web    WebConfig    `json:"web"`
}

var AppConfig *Config

func init() {

	viper.SetDefault("server", ServerConfig{
		Listen: "0.0.0.0",
		Port:   8080,
	})
	viper.SetDefault("admin", AdminConfig{
		Username: "admin",
		Password: "admin123456",
	})
	viper.SetDefault("web", WebConfig{
		Nickname: "timenoteUser",
		Title:    "timenoteWeb",
	})

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	_ = viper.SafeWriteConfig()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	_ = viper.WriteConfig()

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&AppConfig)
		Logger.Info("Config file changed: ", e.Name)
		if err != nil {
			Logger.Error("Unmarshal config file failed: ", err)
		}
	})
	viper.WatchConfig()

	return
}
