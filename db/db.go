package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func Migrate() {
	MigrateUser()
	MigrateProfile()
}

func Setup() error {
	sqlite, err := GetDb()
	if err != nil {
		return err
	}

	db = sqlite
	Migrate()

	return nil
}
