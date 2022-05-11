package web

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/auth"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
)

// HomePage 主页
func HomePage(c *gin.Context) {
	ifLogin := auth.CookieTokenAuth(c)
	if !ifLogin {
		c.Redirect(302, "/login")
	} else {
		var data homeData
		timenoteData := loader.LoadLastDataFile()
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
}
