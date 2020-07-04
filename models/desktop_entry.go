package models

import (
	"github.com/jinzhu/gorm"
)

type DesktopEntry struct {
	gorm.Model
	Icon string
	Path string
}
