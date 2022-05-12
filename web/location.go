package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
	"timenoteWeb/utils"
)

func MapPage(c *gin.Context) {
	var data locationMapData
	timenoteData := loader.LoadLastDataFile()
	tempL := utils.GetLocationNotes(timenoteData.Notes)
	var tempM []simpleLocation
	for _, v := range tempL {
		tempM = append(tempM, simpleLocation{
			Name:  v.Name,
			Lat:   v.Lat,
			Lon:   v.Lon,
			Count: len(v.Notes),
		})
	}
	data.Locations = tempM
	data.Title = "日记地图"
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	data.Token = AppConfig.Map.TokenWeb
	c.HTML(200, "map.html", data)
}

func LocationListPage(c *gin.Context) {
	var data locationListPageData
	timenoteData := loader.LoadLastDataFile()
	tempL := utils.GetLocationNotes(timenoteData.Notes)
	var tempM []simpleLocation
	for _, v := range tempL {
		tempM = append(tempM, simpleLocation{
			Name:  v.Name,
			Lat:   v.Lat,
			Lon:   v.Lon,
			Count: len(v.Notes),
		})
	}
	data.Locations = tempM
	data.Title = "地点"
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	c.HTML(200, "locations.html", data)
}

func LocationPage(c *gin.Context) {
	location := c.Param("name")
	var data locationPageData
	timenoteData := loader.LoadLastDataFile()
	tempL := utils.GetLocationNotes(timenoteData.Notes)
	for _, v := range tempL {
		if v.Name == location {
			data.Title = "地点: " + v.Name
			data.Name = v.Name
			var notes []simpleNote
			for _, note := range v.Notes {
				notes = append(notes, simpleNote{
					ID:           strconv.FormatInt(note.ID, 10),
					Title:        note.Title,
					Date:         note.GetDateStr(),
					Weather:      note.GetWeatherStr(),
					WeatherEmoji: note.GetWeatherEmoji(),
					Mood:         note.GetMoodStr(),
					MoodEmoji:    note.GetMoodEmoji(),
					Location:     note.Location,
					CategoryID:   strconv.FormatInt(note.CategoryID, 10),
					CategoryName: func() string {
						c, s := timenoteData.FindCategory(note)
						if !s {
							return ""
						} else {
							return c.CategoryName
						}
					}(),
				})
			}
			data.Notes = notes
			data.Count = len(notes)
		}
	}
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	c.HTML(200, "location.html", data)
}