package database

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	dsn := viper.GetString("POSTGRES_DSN")
	var err error

	if db, err = sql.Open("postgres", dsn); err != nil {
		log.Fatalf("error connecting to database %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
