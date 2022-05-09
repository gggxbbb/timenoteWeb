package loader

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	. "timenoteWeb/config"
	. "timenoteWeb/log"
	"timenoteWeb/model"
)

// loadGeneralJsonData 将指定 json 文件加载为 model.GeneralData
func loadGeneralJsonData(filename string) model.GeneralData {
	var data model.RawData

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		Logger.Panic(err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		Logger.Panic(err)
	}

	return loadGeneralData(data)
}

// LoadLastJSONFile 加载最新的 json 文件
func LoadLastJSONFile() model.GeneralData {

	log := logging.WithField("源", "LoadLastJSONFile")
	var data model.GeneralData

	dataPath := filepath.Join(AppConfig.Dav.DataPath, "/timeNote/")

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		Logger.Panic(err)
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
		log.Error("未找到最新的 json 文件")
		return data
	} else {
		log.WithField("文件名", lastModifiedFile.Name()).Info("找到最新的 json 文件")
	}

	data = loadGeneralJsonData(filepath.Join(AppConfig.Dav.DataPath, "/timeNote/", lastModifiedFile.Name()))
	data.Source = lastModifiedFile.Name()

	return data
}
