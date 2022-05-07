package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	wd "golang.org/x/net/webdav"
)

func DavServer(prefix string, rootDir string,
	validator func(c *gin.Context) bool,
	logger func(req *http.Request, err error)) gin.HandlerFunc {
	w := wd.Handler{
		Prefix:     prefix,
		FileSystem: wd.Dir(rootDir),
		LockSystem: wd.NewMemLS(),
		Logger:     logger,
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, w.Prefix) {
			if validator != nil && !validator(c) {
				c.AbortWithStatus(403)
				return
			}
			c.Status(200)
			w.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
