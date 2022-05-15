package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
	"timenoteWeb/utils"
)

func SearchPage(c *gin.Context) {
	c.HTML(200, "search.html", basicData{
		Title:    "搜索",
		Nickname: AppConfig.Web.Nickname,
	})
}

func SearchResultPage(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		keyword = c.Query("keyword")
	}
	var data searchResultPageData
	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "搜索"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}

	data.Title = "搜索: " + keyword
	data.Nickname = AppConfig.Web.Nickname
	data.Keyword = keyword
	data.Source = timenoteData.Source

	var NotesWithKeyResult []simpleNoteWithKey

	for _, note := range timenoteData.Notes {
		if strings.Contains(note.Title, keyword) || strings.Contains(note.Content, keyword) {
			var nData simpleNoteWithKey
			nData.Title = note.Title
			nData.ID = strconv.FormatInt(note.ID, 10)
			nData.Date = note.GetDateStr()
			nData.Weather = note.GetWeatherStr()
			nData.WeatherEmoji = note.GetWeatherEmoji()
			nData.Mood = note.GetMoodStr()
			nData.MoodEmoji = note.GetMoodEmoji()
			nData.CategoryID = strconv.FormatInt(note.CategoryID, 10)
			nData.CategoryName = func() string {
				c, s := timenoteData.FindCategory(note)
				if !s {
					return ""
				} else {
					return c.CategoryName
				}
			}()
			nData.Location = note.Location
			if strings.Contains(note.Content, keyword) {
				startIndex := strings.Index(note.Content, keyword)
				endIndex := startIndex + len(keyword)
				keyStart := startIndex - 50
				if keyStart < 0 {
					keyStart = 0
				}
				keyEnd := endIndex + 50
				if keyEnd > len(note.Content) {
					keyEnd = len(note.Content)
				}
				nData.KeyContent = note.Content[keyStart:keyEnd]
			}
			NotesWithKeyResult = append(NotesWithKeyResult, nData)
		}
	}

	data.Notes = NotesWithKeyResult
	data.NoteCount = len(NotesWithKeyResult)

	var CategoriesResult []simpleCategoryWithKey

	for _, category := range timenoteData.Categories {
		if strings.Contains(category.CategoryName, keyword) || strings.Contains(category.CategoryDesc, keyword) {
			var cData simpleCategoryWithKey
			cData.ID = strconv.FormatInt(category.ID, 10)
			cData.Name = category.CategoryName
			cData.NoteCount = len(timenoteData.FindSubNote(category))
			cData.SubcategoryCount = len(timenoteData.FindSubCategory(category))
			if strings.Contains(category.CategoryDesc, keyword) {
				cData.KeyContent = category.CategoryDesc
			}
			CategoriesResult = append(CategoriesResult, cData)
		}
	}

	data.Categories = CategoriesResult
	data.CategoryCount = len(CategoriesResult)

	locations := utils.GetLocationNotes(timenoteData.Notes)
	var locationResult []simpleLocation
	for _, location := range locations {
		if strings.Contains(location.Name, keyword) {
			var lData simpleLocation
			lData.Name = location.Name
			lData.Count = len(location.Notes)
			locationResult = append(locationResult, lData)
		}
	}

	data.Locations = locationResult
	data.LocationCount = len(locationResult)

	c.HTML(200, "search_result.html", data)

}
