package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func StaticServer(prefix string, fs http.FileSystem) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, prefix) {
			c.Status(http.StatusOK)
			http.FileServer(fs).ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
