package routes

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/routes/web"
	"timenoteWeb/utils/auth"
)

func NotesRoute(r *gin.Engine) {
	g := r.Group("/notes", auth.CookieTokenAuthFunc())

	g.GET("/", web.NoteListPage)
	g.GET("/:id", web.NotePage)
}
