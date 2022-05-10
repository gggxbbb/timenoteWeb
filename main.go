package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"strconv"
	. "timenoteWeb/config"
	. "timenoteWeb/log"
	"timenoteWeb/routes"
	"timenoteWeb/utils"
)

//go:embed static/*
var staticData embed.FS

//go:embed templates/*
var templatesData embed.FS

func unescaped(x string) interface{} { return template.HTML(x) }

func main() {

	log := Logger.WithField("源", "main")

	// 初始化数据目录
	if _, err := os.Stat(AppConfig.Dav.DataPath); os.IsNotExist(err) {
		err := os.Mkdir(AppConfig.Dav.DataPath, 0777)
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
	r.Use(utils.LoggerMiddleware())

	// 初始化静态文件
	r.Use(utils.StaticServer("/static", http.FS(staticData)))

	// 初始化时光记 assets 文件服务
	r.Use(utils.AssetsServer("/assets"))

	// 初始化 WebDav 服务
	r.Use(utils.DavServer(
		"/dav",
		AppConfig.Dav.DataPath),
	)

	// 应用 debug 路由
	if gin.Mode() == gin.DebugMode {
		routes.DebugRoute(r)
	}

	// 应用根路由
	routes.RootRoute(r)
	// 应用 API 路由
	routes.ApiRoute(r)
	// 应用 Notes 路由
	routes.NotesRoute(r)

	// Run
	srv := &http.Server{
		Addr:    AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port),
		Handler: r,
	}
	srv.SetKeepAlivesEnabled(true)
	log.Info("Listening on " + AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port))
	err = srv.ListenAndServe()
	if err != nil {
		log.WithError(err).Fatal("服务器启动失败!")
	}
}
