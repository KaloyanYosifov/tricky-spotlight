package models

import (
	"github.com/jinzhu/gorm"
)

type DesktopEntry struct {
	gorm.Model
	Icon           string
	EntryPath      string
	ExecutablePath string
}
