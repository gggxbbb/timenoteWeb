// Package config 用于处理程序配置文件
package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	. "timenoteWeb/log"
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

// logging 包内私有 logging
var logging = Logger.WithField("包", "config")

func init() {

	log := logging.WithField("源", "init")

	log.Info("初始化配置")

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
		log.Info("找不到配置文件, 已创建默认配置文件")
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	_ = viper.WriteConfig()

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.WithError(err).Fatal("解析配置文件失败")
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log := logging.WithField("源", "OnConfigChange")
		err := viper.Unmarshal(&AppConfig)
		log.WithField("文件", e.Name).Info("配置文件变更")
		log.Warn("如果修改了 SERVER 或 DAV 配置, 需要重启服务")
		log.Warn("如果修改了 DAV 配置, 需要手动迁移文件")
		if err != nil {
			log.WithError(err).Fatal("解析配置文件失败")
		}
	})
	viper.WatchConfig()

	log.Info("配置初始化完成, 启动服务器")
	return
}
