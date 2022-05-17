package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/auth"
	"timenoteWeb/web"
)

func TimelineRoute(r *gin.Engine) {
	// Timeline
	r.GET("/timeline", auth.CookieTokenAuthFunc(), web.TimelinePage)
}