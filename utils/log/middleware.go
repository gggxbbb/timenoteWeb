package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		if comment == "" {
			Logger.WithFields(logrus.Fields{
				"状态码":   statusCode,
				"耗时":    latency,
				"来源 IP": clientIP,
				"请求方式":  method,
				"请求路径":  path,
			}).Info("请求已处理")
		} else {
			Logger.WithFields(logrus.Fields{
				"状态码":   statusCode,
				"耗时":    latency,
				"来源 IP": clientIP,
				"请求方式":  method,
				"请求路径":  path,
				"错误信息":  comment,
			}).Warn("请求处理出现问题")
		}
	}
}
