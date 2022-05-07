package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timenoteWeb/model"
)

func GetValues(context *gin.Context) {
	context.JSON(http.StatusOK, Rep{
		Code:   http.StatusOK,
		Msg:    "OK",
		Source: SourceBuiltIn,
		Data: map[string]interface{}{
			"weather": map[int]string{
				model.WeatherCloudy: "阴",
				model.WeatherSunny:  "晴",
				model.WeatherWindy:  "大风",
				model.WeatherSnowy:  "下雪",
				model.WeatherRainy:  "下雨",
			},
			"mood": map[string]string{
				model.MoodHappy:   "开心",
				model.MoodSad:     "难过",
				model.MoodAngry:   "生气",
				model.MoodGloomy:  "阴沉",
				model.MoodNormal:  "一般",
				model.MoodUnknown: "未知",
			},
		},
	})
}
