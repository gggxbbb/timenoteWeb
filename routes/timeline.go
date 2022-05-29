package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/web"
	"timenoteWeb/utils/auth"
)

func TimelineRoute(r *gin.Engine) {
	// Timeline
	r.GET("/timeline", auth.CookieTokenAuthFunc(), web.TimelinePage)
}
