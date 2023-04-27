package repositories

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var _connection *gorm.DB

func connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	dsn := os.Getenv("CONNECTION_STRING")

	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	_connection = db

	return db
}

func Connected() bool {
	return !connection().Config.DryRun
}
