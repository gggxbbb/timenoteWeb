package loader

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	. "timenoteWeb/config"
	. "timenoteWeb/log"
	"timenoteWeb/model"
)

func loadGeneralZipData(filename string) model.GeneralData {
	var rawData model.RawData
	// read the zip file
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		Logger.WithError(err).Panic("Failed to open zip file")
	}
	defer func(zipReader *zip.ReadCloser) {
		_ = zipReader.Close()
	}(zipReader)

	var buf bytes.Buffer

	// iterate through the files
	for _, file := range zipReader.File {
		// read the file
		if file.Name == "backup/data.json" {
			// read the file
			fileReader, err := file.Open()
			if err != nil {
				Logger.WithError(err).Panic("Error opening file: " + file.Name)
			}

			_, err = io.CopyN(&buf, fileReader, int64(file.UncompressedSize64))

			if err != nil {
				Logger.WithError(err).Panic("Error reading file")
				return model.GeneralData{}
			}
		}
	}

	// parse the json
	err = json.Unmarshal(buf.Bytes(), &rawData)
	if err != nil {
		Logger.WithError(err).Panic("Error unmarshalling json")
	}
	return loadGeneralData(rawData)
}

//goland:noinspection GoUnusedExportedFunction
func LoadLastZipFile() model.GeneralData {

	var data model.GeneralData

	dataPath := filepath.Join(AppConfig.Dav.DataPath, "/timeNote/")

	//find last modified zip file in ./data/timeNote/
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

		if strings.HasSuffix(file.Name(), ".zip") && file.ModTime().After(lastModifiedFile.ModTime()) {
			lastModifiedFile = file
		}

	}

	if lastModifiedFile == nil {
		Logger.Error("No zip file found")
		return data
	} else {
		Logger.Info("Last modified zip file: " + lastModifiedFile.Name())
	}

	data = loadGeneralZipData("./data/timeNote/" + lastModifiedFile.Name())
	data.Source = lastModifiedFile.Name()

	return data
}
