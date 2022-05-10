package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timenoteWeb/model"
)

// GetValues 返回基本常量
func GetValues(context *gin.Context) {
	context.JSON(http.StatusOK, Rep{
		Code:   http.StatusOK,
		Msg:    "OK",
		Source: SourceBuiltIn,
		Data: map[string]interface{}{
			"weather": model.WeatherStrMap,
			"mood":    model.MoodStrMap,
		},
	})
}
