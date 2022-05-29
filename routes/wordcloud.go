package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/utils/auth"
	"timenoteWeb/web"
)

func WordcloudRoute(r *gin.Engine) {
	r.GET("/wordcloud", auth.CookieTokenAuthFunc(), web.Wordcloud)
}
