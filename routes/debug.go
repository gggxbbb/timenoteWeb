package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"timenoteWeb/loader/jsonLoader"
)

func DebugRoute(r *gin.Engine, logger *logrus.Logger) {
	r.GET("/debug/data", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			jsonLoader.LoadLastJSONFile(logger))
	})
	r.GET("/debug", func(context *gin.Context) {
		context.HTML(http.StatusOK, "debug_index.html", jsonLoader.LoadLastJSONFile(logger))
	})
	r.GET("/debug/note/:id", func(context *gin.Context) {
		data := jsonLoader.LoadLastJSONFile(logger)
		var opt jsonLoader.NoteData
		for _, note := range data.Notes {
			if strconv.FormatInt(note.ID, 10) == context.Param("id") {
				opt = note
				break
			}
		}
		context.HTML(http.StatusOK, "debug_note.html", opt)
	})
}
