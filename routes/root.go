package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/web"
	"timenoteWeb/utils/auth"
)

func RootRoute(r *gin.Engine) {

	r.GET("/", auth.CookieTokenAuthFunc(), web.HomePage)

	r.GET("/login", web.LoginPage)

	r.POST("/login", web.LoginAction)
}
