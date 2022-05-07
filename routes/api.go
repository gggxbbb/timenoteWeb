package routes

import (
	"github.com/gin-gonic/gin"
	apiFunc "timenoteWeb/api"
)

func ApiRoute(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/values", apiFunc.GetValues)
}
