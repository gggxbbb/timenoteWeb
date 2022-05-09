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
	log := logging.WithField("源", "loadGeneralZipData")
	var rawData model.RawData
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		log.WithError(err).Error("打开压缩文件失败")
	}
	defer func(zipReader *zip.ReadCloser) {
		_ = zipReader.Close()
	}(zipReader)

	var buf bytes.Buffer

	for _, file := range zipReader.File {
		if file.Name == "backup/data.json" {
			fileReader, err := file.Open()
			if err != nil {
				log.WithError(err).WithField("文件名", file.Name).Error("打开文件失败")
			}

			_, err = io.CopyN(&buf, fileReader, int64(file.UncompressedSize64))

			if err != nil {
				log.WithError(err).Error("读取文件失败")
				return model.GeneralData{}
			}
		}
	}

	err = json.Unmarshal(buf.Bytes(), &rawData)
	if err != nil {
		log.WithError(err).Error("解析 json 失败")
	}
	return loadGeneralData(rawData)
}

//goland:noinspection GoUnusedExportedFunction
func LoadLastZipFile() model.GeneralData {

	log := logging.WithField("源", "LoadLastZipFile")
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

		if strings.HasSuffix(file.Name(), ".zip") && file.ModTime().After(lastModifiedFile.ModTime()) {
			lastModifiedFile = file
		}

	}

	if lastModifiedFile == nil {
		log.Error("未找到最新的压缩文件")
		return data
	} else {
		log.WithField("文件名", lastModifiedFile.Name()).Info("找到最新的压缩文件")
	}

	data = loadGeneralZipData(filepath.Join(AppConfig.Dav.DataPath, "/timeNote/", lastModifiedFile.Name()))
	data.Source = lastModifiedFile.Name()

	return data
}
