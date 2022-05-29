package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/utils/auth"
	"timenoteWeb/web"
)

func NotesRoute(r *gin.Engine) {
	g := r.Group("/notes", auth.CookieTokenAuthFunc())

	g.GET("/", web.NoteListPage)
	g.GET("/:id", web.NotePage)
}
