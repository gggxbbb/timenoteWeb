package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	. "timenoteWeb/config"
	. "timenoteWeb/database"
	"timenoteWeb/model"
)

type geocoderRep struct {
	Location struct {
		Lon   float64 `json:"lon"`
		Level string  `json:"level"`
		Lat   float64 `json:"lat"`
	} `json:"location"`
	Status        string `json:"status"`
	Msg           string `json:"msg"`
	SearchVersion string `json:"searchVersion"`
}

func GetLocationByName(name string) (location Location) {
	var log = logging.WithField("源", "GetLocationByName")
	log.WithField("名称", name).Info("开始获取地点")
	DB.Where("name = ?", name).First(&location)
	if location.Name != "" {
		log.WithField("名称", name).Info("获取地点成功")
		return
	} else {
		log.Info("地点不存在, 查询 API")
		//goland:noinspection HttpUrlsUsage
		resp, err := http.Get("https://api.tianditu.gov.cn/geocoder?ds={\"keyWord\":\"" + name + "\"}&tk=" + AppConfig.Map.TokenApi)
		if err != nil {
			log.WithError(err).WithField("名称", name).Error("获取地点失败")
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.WithError(err).WithField("名称", name).Error("关闭连接失败")
			}
		}(resp.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		var rep geocoderRep
		err = json.Unmarshal(body, &rep)
		if err != nil {
			log.WithError(err).WithField("名称", name).Error("解析地点失败")
		}
		if rep.Status == "0" {
			location.Name = name
			location.Lon = rep.Location.Lon
			location.Lat = rep.Location.Lat
			location.Level = rep.Location.Level
			DB.Create(&location)
			log.WithField("名称", name).Info("地点已存储")
			return
		}
		log.WithField("名称", name).Error("获取地点失败")
		return
	}
}

type locationNotes struct {
	Location
	Notes []model.NoteData `json:"notes"`
}

func GetLocationNotes(notes []model.NoteData) (data map[string]locationNotes) {
	var log = logging.WithField("源", "GetLocationNotes")
	log.Info("开始获取地点")
	var tempData = make(map[string][]model.NoteData)
	for _, note := range notes {
		if tempData[note.Location] == nil {
			tempData[note.Location] = []model.NoteData{note}
		} else {
			tempData[note.Location] = append(tempData[note.Location], note)
		}
	}
	data = make(map[string]locationNotes)
	for location, notes := range tempData {
		data[location] = locationNotes{Location: GetLocationByName(location), Notes: notes}
	}
	return
}
