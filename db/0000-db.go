package db

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var wg sync.WaitGroup

func init() {
	wg.Add(1)

	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func WaitInit() {
	wg.Wait()
}
