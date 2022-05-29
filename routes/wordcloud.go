package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/web"
	"timenoteWeb/utils/auth"
)

func WordcloudRoute(r *gin.Engine) {
	r.GET("/wordcloud", auth.CookieTokenAuthFunc(), web.Wordcloud)
}
