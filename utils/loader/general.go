package loader

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	"timenoteWeb/model"
	. "timenoteWeb/utils/config"
)

// loadGeneralData 将 model.RawData 转换为 model.GeneralData
func loadGeneralData(data model.RawData) model.GeneralData {
	var generalData model.GeneralData

	// 未分类
	var defaultCategory = model.CategoryData{
		BgColor:          0,
		CategoryDesc:     "",
		CategoryName:     "未分类",
		ID:               -1,
		IsDefault:        0,
		IsLock:           0,
		NoteCount:        0,
		ParentCategoryID: 0,
		Time:             0,
	}

	generalData.Categories = append(generalData.Categories, defaultCategory)

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

// LoadLastDataFile 加载最新的数据文件
func LoadLastDataFile() (model.GeneralData, bool) {
	log := logging.WithField("源", "LoadLastDataFile")
	var data model.GeneralData

	dataPath := filepath.Join(AppConfig.Data.Root, AppConfig.Data.Dir)

	//find last modified data file in ./data/timeNote/
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.WithError(err).Error("读取数据文件夹失败")
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

		if (strings.HasSuffix(file.Name(), ".zip") || strings.HasSuffix(file.Name(), ".json")) && file.ModTime().After(lastModifiedFile.ModTime()) {
			lastModifiedFile = file
		}

	}

	if lastModifiedFile == nil {
		log.Error("未找到最新的数据文件")
		return data, false
	} else {
		log.WithField("文件名", lastModifiedFile.Name()).Info("找到最新的数据文件")
	}

	path := filepath.Join(AppConfig.Data.Root, AppConfig.Data.Dir, lastModifiedFile.Name())

	if strings.HasSuffix(lastModifiedFile.Name(), ".zip") {
		data = loadGeneralZipData(path)
	} else if strings.HasSuffix(lastModifiedFile.Name(), ".json") {
		data = loadGeneralJsonData(path)
	} else {
		log.WithField("文件名", lastModifiedFile.Name()).Error("数据文件后缀不正确")
		return data, false
	}
	data.Source = lastModifiedFile.Name()
	return data, true
}
