package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"timenoteWeb/auth"
	"timenoteWeb/routes"
	"timenoteWeb/utils"
	"timenoteWeb/webdav"
)

//go:embed assets/*
var assets embed.FS

//go:embed templates/*
var templates embed.FS

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Out = os.Stdout

	// If no data folder, create one
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err := os.Mkdir("./data", 0777)
		if err != nil {
			return
		}
	}

	// init gin
	r := gin.Default()

	// load templates
	templates, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(templates)

	// load static files
	r.StaticFS("/assets/", http.FS(assets))

	//setup logger
	r.Use(utils.LoggerMiddleware(logger))

	// setup webdav
	r.Use(webdav.DavServer(
		"/dav",
		"./data",
		func(c *gin.Context) bool {
			return auth.BasicAuth(c)
		},
		func(req *http.Request, err error) {
			logger.Error(err)
		}),
	)

	routes.DebugRoute(r, logger)

	// run
	err = r.Run(":7080")
	if err != nil {
		return
	}
}
