package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
	"timenoteWeb/auth"
	. "timenoteWeb/config"
)

func AssetsServer(prefix string) gin.HandlerFunc {

	dataPath := path.Join(AppConfig.Dav.DataPath, "timeNote")
	fs := http.FileServer(http.Dir(dataPath))

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, prefix) {
			if auth.CookieTokenAuth(c) {
				c.Status(http.StatusOK)
				fs.ServeHTTP(c.Writer, c.Request)
				c.Abort()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
	}
}