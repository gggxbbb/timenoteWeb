package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
)

type simpleNoteData struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Weather      string `json:"weather"`
	WeatherEmoji string `json:"weatherEmoji"`
	Mood         string `json:"mood"`
	MoodEmoji    string `json:"moodEmoji"`
}

type noteListData struct {
	basicData
	Notes []simpleNoteData `json:"notes"`
}

func NoteListPage(c *gin.Context) {
	data := loader.LoadLastDataFile()
	notes := make([]simpleNoteData, len(data.Notes))
	for i, note := range data.Notes {
		notes[i] = simpleNoteData{
			ID:           strconv.FormatInt(note.ID, 10),
			Title:        note.Title,
			Date:         note.GetDateStr(),
			Weather:      note.GetWeatherStr(),
			WeatherEmoji: note.GetWeatherEmoji(),
			Mood:         note.GetMoodStr(),
			MoodEmoji:    note.GetMoodEmoji(),
		}
	}
	var pData noteListData
	pData.Notes = notes
	pData.Title = "日记列表"
	pData.Nickname = AppConfig.Web.Nickname
	pData.Source = data.Source
	c.HTML(200, "notes.html", pData)
}
