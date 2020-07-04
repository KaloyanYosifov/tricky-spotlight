package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"os"
)

func initDatabase(path string) *gorm.DB {
	// test if we path exists
	// if not create directories
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0600); err != nil {
			panic("Couldn't create database path!")
		}
	}

	databaseFile := path + "/test.db"

	// check if we have a database file
	// if not create it
	if _, err := os.Stat(databaseFile); os.IsNotExist(err) {
		if err := ioutil.WriteFile(databaseFile, []byte{}, 0600); err != nil {
			panic("Couldn't create database!")
		}
	}

	db, err := gorm.Open("sqlite3", databaseFile)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.DesktopEntry{})
}
