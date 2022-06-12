// Package config 用于处理程序配置文件
package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	. "timenoteWeb/utils/log"
)

// AppConfig 全局应用配置
var AppConfig *Config

// logging 包内私有 logging
var logging = Logger.WithField("包", "utils.config")

// 初始化配置文件
func init() {

	log := logging.WithField("源", "init")

	log.Info("初始化配置")

	// 默认值
	viper.SetDefault("server", ServerConfig{
		Listen:       "0.0.0.0",
		Port:         8080,
		Debug:        false,
		EnableWebDav: true,
	})
	viper.SetDefault("data", DataConfig{
		Root: "./data",
		Dir:  "/timeNote/",
	})
	viper.SetDefault("admin", AdminConfig{
		Username: "admin",
		Password: "admin123456",
	})
	viper.SetDefault("web", WebConfig{
		Nickname: "timenoteUser",
		Title:    "timenoteWeb",
	})
	viper.SetDefault("map", MapConfig{
		TokenApi: "",
		TokenWeb: "",
	})
	viper.SetDefault("live", LiveConfig{
		Enable:  true,
		DataDir: "timenoteDoc",
	})

	// 配置文件默认存储于 ./config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 输出默认的配置文件
	err := viper.SafeWriteConfig()
	if err == nil {
		log.Info("找不到配置文件, 已创建默认配置文件")
	}

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// 保存配置文件, 以实现配置文件的迭代
	_ = viper.WriteConfig()

	// 加载配置文件
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.WithError(err).Fatal("解析配置文件失败")
	}

	// 监听配置文件改动
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
