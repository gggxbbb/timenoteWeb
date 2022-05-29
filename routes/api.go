package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/api"
	"timenoteWeb/utils/auth"
)

func ApiRoute(r *gin.Engine) {

	apiR := r.Group("/api")

	apiR.GET("/values", api.GetValues)

	apiAuthed := apiR.Group("/authed", auth.CookieTokenAuthFunc())

	apiAuthed.GET("/locations", api.GetLocations)

}
