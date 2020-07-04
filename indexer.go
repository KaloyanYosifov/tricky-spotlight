package main

import (
	"fmt"
	"github.com/KaloyanYosifov/tricky-spotlight/logger"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"regexp"
)

func indexDesktopEntries(db *gorm.DB) {
	linuxDesktopEntriesPath := "/usr/share/applications"
	files, err := ioutil.ReadDir(linuxDesktopEntriesPath)

	if err != nil {
		panic("Applications folder doesn't exist in this distro!")
	}

	for _, file := range files {
		if matched, err := regexp.MatchString(`.*?\.desktop`, file.Name()); err != nil || !matched {
			if err != nil {
				logger.Debug("indexer: invalid regex")
			}

			continue
		}

		data, err := ioutil.ReadFile(linuxDesktopEntriesPath + file.Name())

		if err != nil {
			logger.Debug("indexer: file couldn't be opened " + file.Name())
			continue
		}

		fmt.Println("yes")

		fmt.Println(data)
		break
	}
}
