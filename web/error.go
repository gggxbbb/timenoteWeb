package web

var (
	errNoDataFile = simpleError{
		Code:  500,
		Title: "NoDataFile",
		Intro: "找不到 记时光 备份文件",
	}
	errNoMapTokenWeb = simpleError{
		Code:  500,
		Title: "NoMapTokenWeb",
		Intro: "天地图 浏览器端 API Token 未设置",
	}
	errNoMapTokenApi = simpleError{
		Code:  500,
		Title: "NoMapTokenApi",
		Intro: "天地图 服务器端 API Token 未设置",
	}
)
