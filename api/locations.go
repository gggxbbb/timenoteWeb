package api

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/loader"
	"timenoteWeb/utils"
)

func GetLocations(c *gin.Context) {
	cData := loader.LoadLastDataFile()
	data := utils.GetLocationNotes(cData.Notes)
	c.JSON(200, data)
}
