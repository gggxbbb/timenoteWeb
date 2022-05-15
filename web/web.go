package web

// basicData 基本前端数据
type basicData struct {
	Title    string `json:"title"`
	Nickname string `json:"nickname"`
	Source   string `json:"source"`
}

// simpleNote 基本笔记数据
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

// noteListData 笔记列表数据
type noteListData struct {
	basicData
	Notes []simpleNote `json:"notes"`
}

// note 某篇笔记数据
type note struct {
	simpleNote
	Content string `json:"content"`
}

// notePageData 笔记页面数据
type notePageData struct {
	basicData
	Note note `json:"note"`
}

// simpleCategory 基本分类数据
type simpleCategory struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parentCategoryID"`
	SubcategoryCount int    `json:"subcategoryCount"`
	NoteCount        int    `json:"noteCount"`
}

// categoryListData 分类列表数据
type categoryListData struct {
	basicData
	Categories []simpleCategory `json:"categories"`
}

type category struct {
	simpleCategory
	Description string `json:"description"`
}

// categoryPageData 分类页面数据
type categoryPageData struct {
	basicData
	category
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

// simpleLocation 基本地点数据
type simpleLocation struct {
	Name  string  `json:"name"`
	Lon   float64 `json:"lon"`
	Lat   float64 `json:"lat"`
	Count int     `json:"count"`
}

// locationMapData 地点地图数据
type locationMapData struct {
	basicData
	Locations []simpleLocation `json:"locations"`
	Token     string           `json:"token"`
}

// locationListPageData 地点页面数据
type locationListPageData struct {
	basicData
	Locations []simpleLocation `json:"locations"`
}

// locationPageData 地点页面数据
type locationPageData struct {
	basicData
	simpleLocation
	Notes []simpleNote `json:"notes"`
}

type simpleNoteWithKey struct {
	simpleNote
	KeyContent string `json:"key_content"`
}

type simpleCategoryWithKey struct {
	simpleCategory
	KeyContent string `json:"key_content"`
}

// searchResultPageData 搜索页面数据
type searchResultPageData struct {
	basicData
	Keyword       string                  `json:"keyword"`
	Notes         []simpleNoteWithKey     `json:"notes_with_key"`
	NoteCount     int                     `json:"note_with_key_count"`
	Categories    []simpleCategoryWithKey `json:"categories"`
	CategoryCount int                     `json:"category_count"`
	Locations     []simpleLocation        `json:"locations"`
	LocationCount int                     `json:"location_count"`
}

// simpleError 错误数据
type simpleError struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
	Intro string `json:"intro"`
}

// errorPageData 错误页面数据
type errorPageData struct {
	basicData
	Error simpleError `json:"error"`
}
