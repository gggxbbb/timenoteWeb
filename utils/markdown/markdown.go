package markdown

import (
	"timenoteWeb/model"
)

//var logging = Logger.WithField("包", "utils.markdown")

func FormatFileName(ipt model.NoteData) string {
	return ipt.GetDateStr() + " " + ipt.Title + ".md"
}
