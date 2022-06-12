package web

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/model/loader"
	. "timenoteWeb/utils/config"
)

// HomePage 主页
func HomePage(c *gin.Context) {
	var data homeData
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "首页"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	data.Title = "首页"
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	data.NoteCount = timenoteData.NoteCount()
	data.CategoryCount = timenoteData.CategoryCount()
	data.TodoCountTotal = timenoteData.TodoCountTotal()
	data.TodoCountDone = timenoteData.TodoCountDone()
	data.TodoCountUndone = timenoteData.TodoCountUndone()
	c.HTML(200, "home.html", data)
}
