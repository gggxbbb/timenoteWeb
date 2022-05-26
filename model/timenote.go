package model

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
	"strings"
	"time"
)

type RawData struct {
	Version  int    `json:"version"`
	Platform string `json:"platform"`
	Tables   []struct {
		Data []struct {
			CategoryID       int64  `json:"categoryId,omitempty"`
			CategoryName     string `json:"categoryName,omitempty"`
			Content          string `json:"content,omitempty"`
			ContentType      int    `json:"contentType,omitempty"`
			ID               int64  `json:"id,omitempty"`
			IsRemove         int    `json:"isRemove,omitempty"`
			Location         string `json:"location,omitempty"`
			Mood             string `json:"mood,omitempty"`
			Music            string `json:"music,omitempty"`
			Time             int64  `json:"time,omitempty"`
			Title            string `json:"title,omitempty"`
			Weather          int    `json:"weather,omitempty"`
			BgColor          int    `json:"bgColor,omitempty"`
			CategoryDesc     string `json:"categoryDesc,omitempty"`
			IsDefault        int    `json:"isDefault,omitempty"`
			IsLock           int    `json:"isLock,omitempty"`
			NoteCount        int    `json:"noteCount,omitempty"`
			ParentCategoryID int64  `json:"parentCategoryId,omitempty"`
			ColorIndex       int    `json:"colorIndex,omitempty"`
			Priority         int    `json:"priority,omitempty"`
			State            int    `json:"state,omitempty"`
			Tags             string `json:"tags,omitempty"`
		} `json:"data"`
		Name string `json:"name"`
	} `json:"tables"`
}

type NoteData struct {
	CategoryID int64 `json:"categoryId"`
	//CategoryName always empty
	CategoryName string `json:"categoryName"`
	Content      string `json:"content"`
	ContentType  int    `json:"contentType"`
	ID           int64  `json:"id"`
	IsRemove     int    `json:"isRemove"`
	Location     string `json:"location"`
	Mood         string `json:"mood"`
	Music        string `json:"music"`
	Time         int64  `json:"time"`
	Title        string `json:"title"`
	Weather      int    `json:"weather"`
}

func (d NoteData) GetWeatherStr() string {
	return WeatherStrMap[d.Weather]
}

func (d NoteData) GetWeatherEmoji() string {
	return WeatherEmojiMap[d.Weather]
}

func (d NoteData) GetMoodStr() string {
	return MoodStrMap[d.Mood]
}

func (d NoteData) GetMoodEmoji() string {
	return MoodEmojiMap[d.Mood]
}

func (d NoteData) GetTimeStr() string {
	timestamp := d.Time / 1000
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func (d NoteData) GetDateStr() string {
	timestamp := d.Time / 1000
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02")
}

func (d NoteData) GetContentHTML() string {
	data := string(blackfriday.Run([]byte(d.Content),
		blackfriday.WithExtensions(blackfriday.CommonExtensions)))
	pData, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return d.Content
	}
	pData.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		if strings.HasPrefix(src, "assets://") {
			newSrc := strings.Replace(src, "assets://", "/assets", -1)
			s.SetAttr("src", newSrc)
		}
	})
	data, err = pData.Html()
	if err != nil {
		return d.Content
	}
	return data
}

func (d NoteData) GetContentText() string {
	data := string(blackfriday.Run([]byte(d.Content),
		blackfriday.WithExtensions(blackfriday.CommonExtensions)))
	pData, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return d.Content
	}
	data = pData.Text()
	return data
}

type CategoryData struct {
	BgColor          int    `json:"bgColor"`
	CategoryDesc     string `json:"categoryDesc"`
	CategoryName     string `json:"categoryName"`
	ID               int64  `json:"id"`
	IsDefault        int    `json:"isDefault"`
	IsLock           int    `json:"isLock"`
	NoteCount        int    `json:"noteCount"`
	ParentCategoryID int64  `json:"parentCategoryId"`
	Time             int64  `json:"time"`
}

func (d CategoryData) IsSubCategory() bool {
	return d.ParentCategoryID != 0
}

type TodoData struct {
	ColorIndex int    `json:"colorIndex"`
	ID         int64  `json:"id"`
	Location   string `json:"location"`
	Priority   int    `json:"priority"`
	//State when 1, not done, when 0, done
	State int    `json:"state"`
	Tags  string `json:"tags"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
}

type GeneralData struct {
	Source     string         `json:"source"`
	Notes      []NoteData     `json:"notes"`
	Todos      []TodoData     `json:"todos"`
	Categories []CategoryData `json:"categories"`
}

func (c GeneralData) DumpBackupLikeJSON() string {

	opt := map[string]interface{}{
		"version":  2,
		"platform": "web",
		"tables": []map[string]interface{}{
			{
				"name": "note",
				"data": c.Notes,
			},
			{
				"name": "todo",
				"data": c.Todos,
			},
			{
				"name": "category",
				"data": c.Categories,
			},
		},
	}

	data, err := json.Marshal(opt)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (c GeneralData) NoteCount() int {
	return len(c.Notes)
}

func (c GeneralData) CategoryCount() int {
	return len(c.Categories)
}

func (c GeneralData) TodoCountTotal() int {
	return len(c.Todos)
}

func (c GeneralData) TodoCountDone() int {
	var count = 0
	for _, todo := range c.Todos {
		if todo.State == 2 {
			count++
		}
	}
	return count
}

func (c GeneralData) TodoCountUndone() int {
	var count = 0
	for _, todo := range c.Todos {
		if todo.State == 1 {
			count++
		}
	}
	return count
}

func (c GeneralData) FindCategory(note NoteData) (CategoryData, bool) {
	for _, category := range c.Categories {
		if category.ID == note.CategoryID {
			return category, true
		}
	}
	return CategoryData{}, false
}

func (c GeneralData) FindParentCategory(childCategory CategoryData) (CategoryData, bool) {
	if !childCategory.IsSubCategory() {
		return CategoryData{}, false
	}
	for _, category := range c.Categories {
		if category.ID == childCategory.ParentCategoryID {
			return category, true
		}
	}
	return CategoryData{}, false
}

func (c GeneralData) FindSubCategory(parentCategory CategoryData) []CategoryData {
	var subCategories []CategoryData
	for _, category := range c.Categories {
		if category.ParentCategoryID == parentCategory.ID {
			subCategories = append(subCategories, category)
		}
	}
	return subCategories
}

func (c GeneralData) FindSubNote(category CategoryData) []NoteData {
	var subNotes []NoteData
	for _, note := range c.Notes {
		if note.CategoryID == category.ID {
			subNotes = append(subNotes, note)
		}
	}
	return subNotes
}
