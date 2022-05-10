package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
)

type noteData struct {
	simpleNoteData
	Content string `json:"content"`
}

type notePageData struct {
	basicData
	Note noteData `json:"note"`
}

func NotePage(c *gin.Context) {
	id := c.Param("id")
	data := loader.LoadLastDataFile()
	var nData noteData
	for _, n := range data.Notes {
		if strconv.FormatInt(n.ID, 10) == id {
			nData.Title = n.Title
			nData.Content = n.GetContentHTML()
			nData.ID = strconv.FormatInt(n.ID, 10)
			nData.Date = n.GetDateStr()
			nData.Weather = n.GetWeatherStr()
			nData.WeatherEmoji = n.GetWeatherEmoji()
			nData.Mood = n.GetMoodStr()
			nData.MoodEmoji = n.GetMoodEmoji()
		}
	}
	var pData notePageData
	pData.Note = nData
	pData.Title = "日记《" + nData.Title + "》"
	pData.Nickname = AppConfig.Web.Nickname
	pData.Source = data.Source
	c.HTML(200, "note.html", pData)
}
