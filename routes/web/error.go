package web

var (
	// errNoDataFile 找不到数据文件
	errNoDataFile = simpleError{
		Code:  500,
		Title: "NoDataFile",
		Intro: "找不到 记时光 备份文件",
	}
	// errNoMapTokenWeb 找不到 map token
	errNoMapTokenWeb = simpleError{
		Code:  500,
		Title: "NoMapTokenWeb",
		Intro: "天地图 浏览器端 API Token 未设置",
	}
	// errNoMapTokenApi 找不到 map token
	errNoMapTokenApi = simpleError{
		Code:  500,
		Title: "NoMapTokenApi",
		Intro: "天地图 服务器端 API Token 未设置",
	}
	// errNoSuchNote 找不到笔记
	errNoSuchNote = simpleError{
		Code:  404,
		Title: "NoteNotFound",
		Intro: "找不到该笔记",
	}
	// errNoSuchCategory 找不到分类
	errNoSuchCategory = simpleError{
		Code:  404,
		Title: "CategoryNotFound",
		Intro: "找不到该分类",
	}
	// errNoSuchLocation 找不到地点
	errNoSuchLocation = simpleError{
		Code:  404,
		Title: "LocationNotFound",
		Intro: "找不到该地点",
	}
)
