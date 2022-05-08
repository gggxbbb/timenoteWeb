package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/web"
)

func RootRoute(r *gin.Engine) {

	r.GET("/", web.HomePage)

	r.GET("/login", web.LoginPage)

	r.POST("/login", web.LoginAction)
}
