package web

var (
	errNoMapTokenWeb = simpleError{
		Title: "NoMapTokenWeb",
		Intro: "天地图 浏览器端 API Token 未设置",
	}
	errNoMapTokenApi = simpleError{
		Title: "NoMapTokenApi",
		Intro: "天地图 服务器端 API Token 未设置",
	}
)
