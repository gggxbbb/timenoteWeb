package web

import (
	"github.com/gin-gonic/gin"
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
