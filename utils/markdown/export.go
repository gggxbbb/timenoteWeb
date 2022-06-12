package markdown

import "timenoteWeb/model"

func ExportWithFrontMatter(ipt model.NoteData, data model.GeneralData) string {
	opt := "---"
	opt += "title: " + ipt.Title + "\n"
	opt += "date: " + ipt.GetTimeStr() + "\n"
	category, success := data.FindCategory(ipt)
	if success {
		opt += "categories: " + category.CategoryName + "\n"
		opt += "categories_desc: " + category.CategoryDesc
	}
	opt += "location: " + ipt.Location + "\n"
	opt += "weather: " + ipt.GetWeatherStr() + "\n"
	opt += "mood: " + ipt.GetMoodStr() + "\n"
	opt += "---\n"
	opt += ipt.Content
	return opt
}
