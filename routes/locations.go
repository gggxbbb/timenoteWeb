package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/auth"
	"timenoteWeb/web"
)

func LocationsRoute(r *gin.Engine) {
	g := r.Group("/locations", auth.CookieTokenAuthFunc())

	g.GET("/map", web.MapPage)
}
