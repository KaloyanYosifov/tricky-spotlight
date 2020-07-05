package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/logger"
	"github.com/KaloyanYosifov/tricky-spotlight/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"regexp"
	"strings"
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

		desktopEntryPath := linuxDesktopEntriesPath + "/" + file.Name()
		var desktopEntryFoundCount int

		// check if we already have an entry with such path
		db.Model(&models.DesktopEntry{}).Where("entry_path = ?", desktopEntryPath).Count(&desktopEntryFoundCount)
		if desktopEntryFoundCount > 0 {
			continue
		}

		data, err := ioutil.ReadFile(desktopEntryPath)

		if err != nil {
			logger.Debug("indexer: file couldn't be opened " + file.Name())
			continue
		}

		params := retrieveParamsFromEntries(strings.Split(string(data), "\n"))
		executablePath, execOk := params["Exec"]
		iconPath, iconOk := params["Icon"]

		if !execOk || !iconOk {
			logger.Debug("indexer: cannot retrieve either icon or executable paths")
			continue
		}

		db.Create(&models.DesktopEntry{
			Icon:           iconPath,
			EntryPath:      desktopEntryPath,
			ExecutablePath: executablePath,
		})
		continue
	}
}

func retrieveParamsFromEntries(data []string) map[string]string {
	params := map[string]string{}

	for _, line := range data {
		keyValuePairs := strings.Split(line, "=")

		// if the keyValuePairs are less than 1
		// then we haven't split by the = sign
		// so in that case we haven't taken a param
		if len(keyValuePairs) <= 1 {
			continue
		}

		params[keyValuePairs[0]] = keyValuePairs[1]
	}

	return params
}
