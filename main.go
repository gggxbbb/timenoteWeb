package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"timenoteWeb/auth"
	"timenoteWeb/utils"
	"timenoteWeb/webdav"
)

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

	// run
	err := r.Run(":7080")
	if err != nil {
		return
	}
}
