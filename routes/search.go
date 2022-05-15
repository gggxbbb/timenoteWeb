package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/auth"
	"timenoteWeb/web"
)

func SearchRoute(r *gin.Engine) {
	g := r.Group("/search", auth.CookieTokenAuthFunc())

	g.GET("/", web.SearchPage)
	g.GET("/:keyword", web.SearchResultPage)
}
