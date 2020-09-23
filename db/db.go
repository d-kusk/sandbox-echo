package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./db/sandbox.sqlite3")

	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db
}
