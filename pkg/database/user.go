package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string `json:"username" gorm:"unique"`
	Password         string
	Name             string            `json:"name"`
	CustomizedVideos []CustomizedVideo `json:"videos"`
}
