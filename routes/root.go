package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timenoteWeb/auth"
	. "timenoteWeb/logger"
)

func RootRoute(r *gin.Engine) {

	r.GET("/login", func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err == nil {
			if auth.CheckToken(token) {
				Logger.Info("Login successful, renew token")
				auth.RenewToken(token)
				c.Redirect(302, "/")
			}
		}
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		_, success := auth.RequireToken(c)
		if !success {
			c.Redirect(302, "/login")
		} else {
			c.Redirect(302, "/")
		}
	})
}
