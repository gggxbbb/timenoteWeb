package model

//goland:noinspection GoUnusedConst
const (
	// WeatherCloudy 阴
	WeatherCloudy = 104
	// WeatherSunny 晴
	WeatherSunny = 150
	// WeatherWindy 大风
	WeatherWindy = 250
	// WeatherSnowy 下雪
	WeatherSnowy = 350
	// WeatherRainy 下雨
	WeatherRainy = 450

	// MoodUnknown 未知
	MoodUnknown = "MOOD_UNKNOWN"
	// MoodHappy 开心
	MoodHappy = "MOOD_HAPPY"
	// MoodSad 难过
	MoodSad = "MOOD_SAD"
	// MoodAngry 生气
	MoodAngry = "MOOD_ANGRY"
	// MoodGloomy 阴沉
	MoodGloomy = "MOOD_GLOOMY"
	// MoodNormal 一般
	MoodNormal = "MOOD_NORMAL"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	WeatherStrMap = map[int]string{
		WeatherCloudy: "阴",
		WeatherSunny:  "晴",
		WeatherWindy:  "大风",
		WeatherSnowy:  "下雪",
		WeatherRainy:  "下雨",
	}
	MoodStrMap = map[string]string{
		MoodHappy:   "开心",
		MoodSad:     "难过",
		MoodAngry:   "生气",
		MoodGloomy:  "阴沉",
		MoodNormal:  "一般",
		MoodUnknown: "未知",
	}
)

//goland:noinspection GoUnusedGlobalVariable
var (
	WeatherEmojiMap = map[int]string{
		WeatherCloudy: "☁️",
		WeatherSunny:  "☀️",
		WeatherWindy:  "🍃",
		WeatherSnowy:  "❄️",
		WeatherRainy:  "🌧️",
	}
	MoodEmojiMap = map[string]string{
		MoodHappy:   "😄",
		MoodSad:     "😢",
		MoodAngry:   "😠",
		MoodGloomy:  "😔",
		MoodNormal:  "😐",
		MoodUnknown: "🤔",
	}
)
