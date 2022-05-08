package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timenoteWeb/auth"
	. "timenoteWeb/log"
)

func LoginPage(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil {
		if auth.CheckToken(token) {
			Logger.Info("Login successful, renew token")
			auth.RenewToken(token)
			c.Redirect(302, "/")
		}
	}
	c.HTML(http.StatusOK, "login.html", BasicData{
		Title: "登录",
	})
}

func LoginAction(c *gin.Context) {
	success := auth.RequireToken(c)
	if !success {
		c.Redirect(302, "/login")
	} else {
		c.Redirect(302, "/")
	}
}
