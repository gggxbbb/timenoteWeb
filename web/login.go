package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timenoteWeb/auth"
)

// LoginPage 登录页
func LoginPage(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil {
		if auth.CheckToken(token) {
			auth.RenewToken(token)
			c.Redirect(302, "/")
		}
	}
	c.HTML(http.StatusOK, "login.html", basicData{
		Title: "登录",
	})
}

// LoginAction 登录请求
func LoginAction(c *gin.Context) {
	success := auth.RequireToken(c)
	if !success {
		c.Redirect(302, "/login")
	} else {
		c.Redirect(302, "/")
	}
}
