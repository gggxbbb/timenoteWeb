package web

// basicData 基本前端数据
type basicData struct {
	Title    string `json:"title"`
	Nickname string `json:"nickname"`
	Source   string `json:"source"`
}

type simpleNote struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Weather      string `json:"weather"`
	WeatherEmoji string `json:"weatherEmoji"`
	Mood         string `json:"mood"`
	MoodEmoji    string `json:"moodEmoji"`
	CategoryName string `json:"categoryName"`
	CategoryID   string `json:"categoryID"`
	Location     string `json:"location"`
}

type noteListData struct {
	basicData
	Notes []simpleNote `json:"notes"`
}

type note struct {
	simpleNote
	Content string `json:"content"`
}

type notePageData struct {
	basicData
	Note note `json:"note"`
}

type simpleCategory struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parentCategoryID"`
	SubcategoryCount int    `json:"subcategoryCount"`
	NoteCount        int    `json:"noteCount"`
}

type categoryListData struct {
	basicData
	Categories []simpleCategory `json:"categories"`
}

type categoryPageData struct {
	basicData
	simpleCategory
	Notes         []simpleNote     `json:"notes"`
	Subcategories []simpleCategory `json:"subcategories"`
}

// homeData 主页数据
type homeData struct {
	basicData
	NoteCount       int `json:"note_count"`
	CategoryCount   int `json:"category_count"`
	TodoCountTotal  int `json:"todo_count_total"`
	TodoCountDone   int `json:"todo_count_done"`
	TodoCountUndone int `json:"todo_count_undone"`
}

type simpleLocation struct {
	Name  string  `json:"name"`
	Lon   float64 `json:"lon"`
	Lat   float64 `json:"lat"`
	Count int     `json:"count"`
}

type locationMapData struct {
	basicData
	Locations []simpleLocation `json:"locations"`
	Token     string           `json:"token"`
}

type locationListPageData struct {
	basicData
	Locations []simpleLocation `json:"locations"`
}

type locationPageData struct {
	basicData
	simpleLocation
	Notes []simpleNote `json:"notes"`
}

type simpleError struct {
	Title string `json:"title"`
	Intro string `json:"intro"`
}

type errorPageData struct {
	basicData
	Error simpleError `json:"error"`
}
