package main

import (
	"context"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
	"timenoteWeb/routes"
	. "timenoteWeb/utils/config"
	"timenoteWeb/utils/live"
	. "timenoteWeb/utils/log"
	"timenoteWeb/utils/server"
)

//go:embed static/*
var staticData embed.FS

//go:embed templates/*
var templatesData embed.FS

var (
	VERSION string
	BUILD   string
)

func unescaped(x string) interface{} { return template.HTML(x) }

func main() {

	log := Logger.WithField("源", "main")

	banner, _ := staticData.ReadFile("static/banner.txt")
	bannerStr := strings.Split(string(banner), "\n")
	for _, v := range bannerStr {
		log.Info(v)
	}

	log.Info("记时光 WebViewer")
	log.Info("版本: ", VERSION)
	log.Info("构建: ", BUILD)

	// 初始化数据目录
	if _, err := os.Stat(AppConfig.Data.Root); os.IsNotExist(err) {
		log.Info("数据目录不存在, 初始化数据目录")
		err := os.Mkdir(AppConfig.Data.Root, 0777)
		if err != nil {
			log.WithError(err).Fatal("无法新建数据目录!")
		}
	}

	// 配置调试模式
	if AppConfig.Server.Debug {
		gin.SetMode(gin.DebugMode)
		log.Warn("调试模式已开启!")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化 gin
	r := gin.New()

	// 使用 gin 自带 Recovery
	r.Use(gin.Recovery())

	// 加载前端页面模板
	templates := template.New("")
	templates = templates.Funcs(template.FuncMap{"unescaped": unescaped})
	templates, err := templates.ParseFS(templatesData, "templates/*.html")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(templates)

	// 初始化日志
	r.Use(LoggerMiddleware())

	// 初始化静态文件
	r.Use(server.StaticServer("/static", http.FS(staticData)))

	// 初始化记时光 assets 文件服务
	r.Use(server.AssetsServer("/assets"))

	// 初始化 WebDav 服务
	if AppConfig.Server.EnableWebDav {
		log.Info("WebDav 服务已开启")
		r.Use(server.DavServer(
			"/dav",
			AppConfig.Data.Root),
		)
	} else {
		log.Info("WebDav 服务已关闭")
	}

	// 应用路由
	routes.EnableRoute(r)

	// Run
	live.WatchBackupDataPath()
	srv := &http.Server{
		Addr:    AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port),
		Handler: r,
	}
	log.Info("listen on " + AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port))
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("listen error")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Server 停止中...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("出现异常")
	}
	log.Info("再见, 我的朋友!")
}
