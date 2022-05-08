package loader

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	. "timenoteWeb/config"
	. "timenoteWeb/logger"
	"timenoteWeb/model"
)

func loadGeneralJsonData(filename string) model.GeneralData {
	var data model.RawData

	// read the file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		Logger.Panic(err)
	}

	//load the data
	err = json.Unmarshal(file, &data)
	if err != nil {
		Logger.Panic(err)
	}

	return loadGeneralData(data)
}

func LoadLastJSONFile() model.GeneralData {

	var data model.GeneralData

	dataPath := filepath.Join(AppConfig.Dav.DataPath, "/timeNote/")

	//find last modified json file in ./data/timeNote/
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
		Logger.Error("No json file found")
		return data
	} else {
		Logger.Info("Last modified json file: " + lastModifiedFile.Name())
	}

	data = loadGeneralJsonData("./data/timeNote/" + lastModifiedFile.Name())
	data.Source = lastModifiedFile.Name()

	return data
}
