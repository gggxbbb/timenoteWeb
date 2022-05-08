package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"timenoteWeb/auth"
	. "timenoteWeb/config"
	. "timenoteWeb/logger"
	"timenoteWeb/routes"
	"timenoteWeb/utils"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

func main() {

	// If no data folder, create one
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err := os.Mkdir("./data", 0777)
		if err != nil {
			return
		}
	}

	// setup debug mode
	if AppConfig.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// init gin
	r := gin.New()

	// setup recovery
	r.Use(gin.Recovery())

	// load templates
	templates, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(templates)

	// setup static files
	r.Use(utils.StaticServer("/static", http.FS(static)))

	// setup logger
	r.Use(utils.LoggerMiddleware())

	// setup webdav
	r.Use(utils.DavServer(
		"/dav",
		AppConfig.Dav.DataPath,
		func(c *gin.Context) bool {
			return auth.BasicAuth(c)
		},
		func(req *http.Request, err error) {
			Logger.Error(err)
		}),
	)

	if gin.Mode() == gin.DebugMode {
		routes.DebugRoute(r)
	}

	routes.RootRoute(r)
	routes.ApiRoute(r)

	// run
	srv := &http.Server{
		Addr:    AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port),
		Handler: r,
	}
	srv.SetKeepAlivesEnabled(true)
	Logger.Info("Listening on " + AppConfig.Server.Listen + ":" + strconv.Itoa(AppConfig.Server.Port))
	Logger.Fatal(srv.ListenAndServe())
}
