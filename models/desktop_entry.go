package models

import (
	"github.com/jinzhu/gorm"
)

type DesktopEntry struct {
	gorm.Model
	Name           string `gorm:"index"`
	Icon           string
	EntryPath      string `gorm:"index"`
	ExecutablePath string
}

func SearchForDesktopEntry(name string, db *gorm.DB) []DesktopEntry {
	var desktopEntries []DesktopEntry
	db.Where("name like ?", "%"+name+"%").Find(&desktopEntries)

	return desktopEntries
}
