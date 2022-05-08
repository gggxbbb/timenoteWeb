package jsonLoader

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strings"
	. "timenoteWeb/logger"
	"timenoteWeb/model"
)

func loadRawData(filename string) model.RawData {
	var data model.RawData

	// read the file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//load the data
	err = json.Unmarshal(file, &data)

	return data
}

func LoadGeneralData(filename string) model.GeneralData {
	var data model.RawData

	var generalData model.GeneralData

	data = loadRawData(filename)

	for _, v := range data.Tables {
		if v.Name == "note" {
			for _, v2 := range v.Data {
				generalData.Notes = append(generalData.Notes, model.NoteData{
					CategoryID:   v2.CategoryID,
					CategoryName: v2.CategoryName,
					Content:      v2.Content,
					ContentType:  v2.ContentType,
					ID:           v2.ID,
					IsRemove:     v2.IsRemove,
					Location:     v2.Location,
					Mood:         v2.Mood,
					Music:        v2.Music,
					Time:         v2.Time,
					Title:        v2.Title,
					Weather:      v2.Weather,
				})
			}
		} else if v.Name == "category" {
			for _, v2 := range v.Data {
				generalData.Categories = append(generalData.Categories, model.CategoryData{
					BgColor:          v2.BgColor,
					CategoryDesc:     v2.CategoryDesc,
					CategoryName:     v2.CategoryName,
					ID:               v2.ID,
					IsDefault:        v2.IsDefault,
					IsLock:           v2.IsLock,
					NoteCount:        v2.NoteCount,
					ParentCategoryID: v2.ParentCategoryID,
					Time:             v2.Time,
				})
			}
		} else if v.Name == "todo" {
			for _, v2 := range v.Data {
				generalData.Todos = append(generalData.Todos, model.TodoData{
					ColorIndex: v2.ColorIndex,
					ID:         v2.ID,
					Location:   v2.Location,
					Priority:   v2.Priority,
					State:      v2.State,
					Tags:       v2.Tags,
					Time:       v2.Time,
					Title:      v2.Title,
				})
			}
		}
	}

	return generalData
}

func LoadLastJSONFile() model.GeneralData {

	var data model.GeneralData

	//find last modified json file in ./data/timeNote/
	files, err := ioutil.ReadDir("./data/timeNote/")
	if err != nil {
		panic(err)
	}

	var lastModifiedFile fs.FileInfo

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if lastModifiedFile == nil {
			lastModifiedFile = file
			continue
		}

		if strings.HasSuffix(file.Name(), ".json") && file.ModTime().After(lastModifiedFile.ModTime()) {
			lastModifiedFile = file
		}

	}

	if lastModifiedFile == nil {
		Logger.Error("No json file found")
		return data
	} else {
		Logger.Info("Last modified json file: " + lastModifiedFile.Name())
	}

	data = LoadGeneralData("./data/timeNote/" + lastModifiedFile.Name())
	data.Source = lastModifiedFile.Name()

	return data
}
