package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"timenoteWeb/auth"
	"timenoteWeb/config"
	"timenoteWeb/loader/jsonLoader"
)

func DebugRoute(r *gin.Engine, config *config.Config, logger *logrus.Logger) {

	debug := r.Group("/debug", auth.BasicAuthFunc(config))

	debug.GET("/data", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			jsonLoader.LoadLastJSONFile(logger))
	})
	debug.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "debug_index.html", jsonLoader.LoadLastJSONFile(logger))
	})
	debug.GET("/note/:id", func(context *gin.Context) {
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
