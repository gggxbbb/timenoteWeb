package live

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"timenoteWeb/model/loader"
	. "timenoteWeb/utils/config"
	"timenoteWeb/utils/markdown"
)

func exportNotes() {
	var log = logging.WithField("源", "exportNotes")
	timenoteData, success := loader.LoadLastDataFile()
	exportDir := filepath.Join(AppConfig.Data.Root, AppConfig.Live.DataDir)
	if _, err := os.Stat(exportDir); os.IsNotExist(err) {
		log.Info("导出目录不存在, 初始化导出目录")
		err := os.Mkdir(exportDir, 0777)
		if err != nil {
			log.WithError(err).Fatal("无法新建数据目录!")
		}
	}
	if success {
		log.WithField("源文件", timenoteData.Source).Info("读取最新数据成功")
		for i := range timenoteData.Notes {
			note := timenoteData.Notes[i]
			filename := markdown.FormatFileName(note)
			fileContent := markdown.ExportWithFrontMatter(note, timenoteData)
			filePath := filepath.Join(exportDir, filename)
			log.WithField("文件名", filename).Info("导出日记")
			err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
			if err != nil {
				log.WithError(err).WithField("文件名", filename).Warn("导出失败")
			}
		}
	}
}
