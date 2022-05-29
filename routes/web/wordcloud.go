package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ego/gse"
	. "timenoteWeb/utils/config"
	"timenoteWeb/utils/loader"
)

func Wordcloud(c *gin.Context) {
	//TODO: 无意义词汇过滤、标点剔除，待完成 (似乎不做也不是不行)
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "词云"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	var dataString string
	for _, item := range timenoteData.Notes {
		dataString += item.GetContentText() + "\n"
	}

	x, _ := gse.NewEmbed()

	words := x.Cut(dataString, true)
	words = x.Trim(words)

	tempData := make(map[string]int)
	for _, word := range words {
		if _, ok := tempData[word]; ok {
			tempData[word]++
		} else {
			tempData[word] = 1
		}
	}

	var data []wordcloudData
	for word, count := range tempData {
		data = append(data, wordcloudData{word, count})
	}

	var pData wordcloudPageData
	pData.Title = "词云α"
	pData.Nickname = AppConfig.Web.Nickname
	pData.Words = data
	pData.Source = timenoteData.Source

	c.HTML(200, "wordcloud.html", pData)

}
