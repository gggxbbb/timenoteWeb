package database

import "gorm.io/gorm"

// Location 地点数据模型
type Location struct {
	gorm.Model `json:"-"`

	// Name 地点名称
	Name string `json:"name"`

	// Lon 经度
	Lon float64 `json:"lon"`

	// Lat 纬度
	Lat float64 `json:"lat"`

	// Level 见 天地图 API 文档
	Level string `json:"level"`
}
