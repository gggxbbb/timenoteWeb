package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	. "timenoteWeb/logger"
)

type ServerConfig struct {
	Listen string `json:"listen" mapstructure:"listen"`
	Port   int    `json:"port" mapstructure:"port"`
	Debug  bool   `json:"debug" mapstructure:"debug"`
}

type DavConfig struct {
	DataPath string `json:"dataPath" mapstructure:"data_path"`
}

type AdminConfig struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

type WebConfig struct {
	Nickname string `json:"nickname" mapstructure:"nickname"`
	Title    string `json:"title" mapstructure:"title"`
}

type Config struct {
	Server ServerConfig `json:"server" mapstructure:"server"`
	Dav    DavConfig    `json:"dav" mapstructure:"dav"`
	Admin  AdminConfig  `json:"admin" mapstructure:"admin"`
	Web    WebConfig    `json:"web" mapstructure:"web"`
}

var AppConfig *Config

func init() {

	Logger.Info("init config")

	viper.SetDefault("server", ServerConfig{
		Listen: "0.0.0.0",
		Port:   8080,
		Debug:  false,
	})
	viper.SetDefault("dav", DavConfig{
		DataPath: "./data",
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

	err := viper.SafeWriteConfig()
	if err == nil {
		Logger.Info("Can't find config file, create a new one")
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	_ = viper.WriteConfig()

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		Logger.Fatal("Unmarshal config failed: %s", err.Error())
	}

	Logger.Info("config loaded, initializing server")

	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&AppConfig)
		Logger.Info("Config file changed: ", e.Name)
		Logger.Info("If you changed the config of SERVER, you need to restart the this application.")
		if err != nil {
			Logger.Error("Unmarshal config file failed: ", err)
		}
	})
	viper.WatchConfig()

	return
}
