package database

import (
	"github.com/KaloyanYosifov/tricky-spotlight/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"os"
)

type Database struct {
	*gorm.DB
}

var globalDB *Database

func InitDatabase(path string) *Database {
	if globalDB != nil {
		return globalDB
	}

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

	globalDB = &Database{db}

	return globalDB
}

func GetDatabase() *Database {
	return globalDB
}

func (db *Database) MigrateModels() {
	db.AutoMigrate(&models.DesktopEntry{})
}

func (db *Database) GetUnderilyingDB() *gorm.DB {
	return db.DB
}

func (db *Database) Close() error {
	err := db.DB.Close()

	if err != nil {
		return err
	}

	globalDB = nil

	return nil
}
