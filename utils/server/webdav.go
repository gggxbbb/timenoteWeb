package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"timenoteWeb/utils/auth"

	"github.com/gin-gonic/gin"
	wd "golang.org/x/net/webdav"
)

func logger(req *http.Request, err error) {
	log := logging.WithField("源", "DavServer")
	if err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"来源": req.RemoteAddr,
			"路径": req.URL.Path,
			"方法": req.Method,
		}).Error("WebDav 服务异常!")
	}
}

func DavServer(prefix string, rootDir string) gin.HandlerFunc {
	//log := logging.WithField("源", "DavServer")
	w := wd.Handler{
		Prefix:     prefix,
		FileSystem: wd.Dir(rootDir),
		LockSystem: wd.NewMemLS(),
		Logger:     logger,
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, w.Prefix) {
			auth.BasicAuthFunc()(c)
			c.Status(200)
			w.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
