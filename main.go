package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"timenoteWeb/auth"
	"timenoteWeb/config"
	"timenoteWeb/routes"
	"timenoteWeb/utils"
	"timenoteWeb/webdav"
)

//go:embed assets/*
var assets embed.FS

//go:embed templates/*
var templates embed.FS

var logger *logrus.Logger

func main() {

	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Out = os.Stdout

	// If no data folder, create one
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err := os.Mkdir("./data", 0777)
		if err != nil {
			return
		}
	}

	// Load config
	appConfig, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	// init gin
	r := gin.Default()

	// load templates
	templates, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(templates)

	// load assets files
	r.StaticFS("/assets/", http.FS(assets))

	//setup logger
	r.Use(utils.LoggerMiddleware(logger))

	// setup webdav
	r.Use(webdav.DavServer(
		"/dav",
		"./data",
		func(c *gin.Context) bool {
			return auth.BasicAuth(c, appConfig)
		},
		func(req *http.Request, err error) {
			logger.Error(err)
		}),
	)

	if gin.Mode() == gin.DebugMode {
		routes.DebugRoute(r, appConfig, logger)
	}

	// run
	srv := &http.Server{
		Addr:    appConfig.Listen + ":" + strconv.Itoa(appConfig.Port),
		Handler: r,
	}
	srv.SetKeepAlivesEnabled(true)
	logger.Info("Listening on " + appConfig.Listen + ":" + strconv.Itoa(appConfig.Port))
	logger.Fatal(srv.ListenAndServe())
}
