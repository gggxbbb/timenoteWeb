package database

import "gorm.io/gorm"

type Location struct {
	gorm.Model `json:"-"`
	Name       string  `json:"name"`
	Lon        float64 `json:"lon"`
	Lat        float64 `json:"lat"`
	Level      string  `json:"level"`
}
