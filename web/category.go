package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
)

func CategoryListPage(c *gin.Context) {
	data := loader.LoadLastDataFile()
	var categories []simpleCategory
	for _, category := range data.Categories {
		if !category.IsSubCategory() {
			var d simpleCategory
			d.ID = strconv.FormatInt(category.ID, 10)
			d.Name = category.CategoryName
			d.NoteCount = len(data.FindSubNote(category))
			d.SubcategoryCount = len(data.FindSubCategory(category))
			categories = append(categories, d)
		}
	}
	var pData categoryListData
	pData.Categories = categories
	pData.Title = "分类列表"
	pData.Nickname = AppConfig.Web.Nickname
	pData.Source = data.Source
	c.HTML(200, "categories.html", pData)
}

func CategoryPage(c *gin.Context) {
	id := c.Param("id")
	data := loader.LoadLastDataFile()
	var pData categoryPageData
	for _, category := range data.Categories {
		if strconv.FormatInt(category.ID, 10) == id {
			pData.Title = "分类：" + category.CategoryName
			pData.ID = strconv.FormatInt(category.ID, 10)
			pData.Name = category.CategoryName
			notes := data.FindSubNote(category)
			pData.NoteCount = len(notes)
			for _, note := range notes {
				pData.Notes = append(pData.Notes, simpleNote{
					ID:           strconv.FormatInt(note.ID, 10),
					Title:        note.Title,
					Date:         note.GetDateStr(),
					Weather:      note.GetWeatherStr(),
					WeatherEmoji: note.GetWeatherEmoji(),
					Mood:         note.GetMoodStr(),
					MoodEmoji:    note.GetMoodEmoji(),
					Location:     note.Location,
					CategoryID:   strconv.FormatInt(note.CategoryID, 10),
					CategoryName: func() string {
						c, s := data.FindCategory(note)
						if !s {
							return ""
						} else {
							return c.CategoryName
						}
					}(),
				})
			}
			categories := data.FindSubCategory(category)
			pData.SubcategoryCount = len(categories)
			for _, subcategory := range categories {
				pData.Subcategories = append(pData.Subcategories, simpleCategory{
					ID:               strconv.FormatInt(subcategory.ID, 10),
					Name:             subcategory.CategoryName,
					NoteCount:        len(data.FindSubNote(category)),
					SubcategoryCount: len(data.FindSubCategory(category)),
				})
			}
			if category.IsSubCategory() {
				pData.ParentCategoryID = strconv.FormatInt(category.ParentCategoryID, 10)
			}
		}
	}
	pData.Nickname = AppConfig.Web.Nickname
	pData.Source = data.Source
	c.HTML(200, "category.html", pData)
}
