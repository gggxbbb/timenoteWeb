package model

import (
	"encoding/json"
	"time"
)

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

type RawData struct {
	Version  int    `json:"version"`
	Platform string `json:"platform"`
	Tables   []struct {
		Data []struct {
			CategoryID       int64  `json:"categoryId"`
			CategoryName     string `json:"categoryName"`
			Content          string `json:"content"`
			ContentType      int    `json:"contentType"`
			ID               int64  `json:"id"`
			IsRemove         int    `json:"isRemove"`
			Location         string `json:"location"`
			Mood             string `json:"mood"`
			Music            string `json:"music"`
			Time             int64  `json:"time"`
			Title            string `json:"title"`
			Weather          int    `json:"weather"`
			BgColor          int    `json:"bgColor"`
			CategoryDesc     string `json:"categoryDesc"`
			IsDefault        int    `json:"isDefault"`
			IsLock           int    `json:"isLock"`
			NoteCount        int    `json:"noteCount"`
			ParentCategoryID int64  `json:"parentCategoryId"`
			ColorIndex       int    `json:"colorIndex"`
			Priority         int    `json:"priority"`
			State            int    `json:"state"`
			Tags             string `json:"tags"`
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
	switch d.Weather {
	case WeatherCloudy:
		return "阴"
	case WeatherSunny:
		return "晴"
	case WeatherWindy:
		return "大风"
	case WeatherSnowy:
		return "下雪"
	case WeatherRainy:
		return "下雨"
	default:
		return "未知"
	}
}

func (d NoteData) GetMoodStr() string {
	switch d.Mood {
	case MoodHappy:
		return "开心"
	case MoodSad:
		return "难过"
	case MoodAngry:
		return "生气"
	case MoodGloomy:
		return "阴沉"
	case MoodNormal:
		return "一般"
	default:
		return "未知"
	}
}

func (d NoteData) GetTimeStr() string {
	timestamp := d.Time / 1000
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
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
