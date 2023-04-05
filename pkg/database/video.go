package database

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Url        string     `json:"url"`
	Categories []Category `json:"categories" gorm:"many2many:video_categories;"`
}

type CustomizedVideo struct {
	gorm.Model
	Url    string `json:"url"`
	UserID uint
}
