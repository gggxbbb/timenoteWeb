package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"timenoteWeb/model/loader"
	. "timenoteWeb/utils/config"
	"timenoteWeb/utils/map"
)

// MapPage 地图页面
func MapPage(c *gin.Context) {
	if AppConfig.Map.TokenWeb == "" {
		var data errorPageData
		data.Title = "日记地图"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoMapTokenWeb
		c.HTML(200, "error.html", data)
		return
	}
	if AppConfig.Map.TokenApi == "" {
		var data errorPageData
		data.Title = "日记地图"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoMapTokenApi
		c.HTML(200, "error.html", data)
		return
	}
	var data locationMapData
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "日记地图"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	tempL := _map.GetLocationNotes(timenoteData.Notes)
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

// LocationListPage 地点列表页面
func LocationListPage(c *gin.Context) {
	var data locationListPageData
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "地点列表"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	tempL := _map.GetLocationNotes(timenoteData.Notes)
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
	data.Title = "地点列表"
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	c.HTML(200, "locations.html", data)
}

// LocationPage 地点页面
func LocationPage(c *gin.Context) {
	location := c.Param("name")
	var data locationPageData
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "地点"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	tempL := _map.GetLocationNotes(timenoteData.Notes)
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
			break
		}
	}
	if data.Name == "" {
		var data errorPageData
		data.Title = "地点"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoSuchLocation
		c.HTML(errNoSuchLocation.Code, "error.html", data)
		return
	}
	data.Source = timenoteData.Source
	data.Nickname = AppConfig.Web.Nickname
	c.HTML(200, "location.html", data)
}
