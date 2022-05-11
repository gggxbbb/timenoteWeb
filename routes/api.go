package routes

import (
	"github.com/gin-gonic/gin"
	apiFunc "timenoteWeb/api"
	"timenoteWeb/auth"
)

func ApiRoute(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/values", apiFunc.GetValues)

	apiAuthed := api.Group("/authed", auth.CookieTokenAuthFunc())

	apiAuthed.GET("/locations", apiFunc.GetLocations)

}
