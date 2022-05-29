package api

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/utils/loader"
	"timenoteWeb/utils/map"
)

// GetLocations 获取根据地点分类的日记和地点的详细坐标
func GetLocations(c *gin.Context) {
	cData, _ := loader.LoadLastDataFile()
	data := _map.GetLocationNotes(cData.Notes)
	c.JSON(200, data)
}
