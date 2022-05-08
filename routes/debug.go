package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"timenoteWeb/auth"
	. "timenoteWeb/config"
	"timenoteWeb/loader/jsonLoader"
	"timenoteWeb/model"
)

func DebugRoute(r *gin.Engine) {

	debug := r.Group("/debug", auth.CookieTokenAuthFunc())

	debug.GET("/data", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			jsonLoader.LoadLastJSONFile())
	})
	debug.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "debug_index.html", jsonLoader.LoadLastJSONFile())
	})
	debug.GET("/note/:id", func(context *gin.Context) {
		data := jsonLoader.LoadLastJSONFile()
		var opt model.NoteData
		for _, note := range data.Notes {
			if strconv.FormatInt(note.ID, 10) == context.Param("id") {
				opt = note
				break
			}
		}
		context.HTML(http.StatusOK, "debug_note.html", gin.H{
			"title": opt.Title,
			"category": func() string {
				d, b := data.FindCategory(opt)
				if b {
					return d.CategoryName
				} else {
					return "nil"
				}
			}(),
			"content":  opt.Content,
			"time":     opt.GetTimeStr(),
			"mood":     opt.GetMoodStr(),
			"weather":  opt.GetWeatherStr(),
			"location": opt.Location,
		})
	})
	debug.GET("/config", func(context *gin.Context) {
		context.JSON(http.StatusOK, AppConfig)
	})
	debug.GET("/count", func(context *gin.Context) {
		context.JSON(http.StatusOK, func() gin.H {
			data := jsonLoader.LoadLastJSONFile()
			return gin.H{
				"source":      data.Source,
				"notes":       data.NoteCount(),
				"category":    data.CategoryCount(),
				"todo_all":    data.TodoCountTotal(),
				"todo_done":   data.TodoCountDone(),
				"todo_undone": data.TodoCountUndone(),
			}
		}())
	})
}
