package jsonLoader

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
	CategoryID   int64  `json:"categoryId"`
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

type TodoData struct {
	ColorIndex int    `json:"colorIndex"`
	ID         int64  `json:"id"`
	Location   string `json:"location"`
	Priority   int    `json:"priority"`
	State      int    `json:"state"`
	Tags       string `json:"tags"`
	Time       int64  `json:"time"`
	Title      string `json:"title"`
}

type GeneralData struct {
	Notes      []NoteData     `json:"notes"`
	Todos      []TodoData     `json:"todos"`
	Categories []CategoryData `json:"categories"`
}
