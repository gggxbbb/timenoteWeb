package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/utils/auth"
	"timenoteWeb/web"
)

func LocationsRoute(r *gin.Engine) {
	g := r.Group("/locations", auth.CookieTokenAuthFunc())

	g.GET("/", web.LocationListPage)
	g.GET("/map", web.MapPage)
	g.GET("/:name", web.LocationPage)

}
