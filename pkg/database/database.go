package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := viper.GetString("POSTGRES_DSN")
	var err error

	if db, err = gorm.Open(postgres.Open(dsn)); err != nil {
		log.Fatalf("error connecting to database %v", err)
	}

	db.AutoMigrate(&User{}, &Category{}, &Video{}, &CustomizedVideo{})
}

func GetDB() *gorm.DB {
	return db
}
