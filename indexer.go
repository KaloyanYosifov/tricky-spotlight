package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/database"
	"github.com/KaloyanYosifov/tricky-spotlight/logger"
	"github.com/KaloyanYosifov/tricky-spotlight/models"
	"io/ioutil"
	"regexp"
	"strings"
)

func indexDesktopEntries(db *database.Database) {
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
		name, nameOk := params["Name"]

		if !execOk || !iconOk || !nameOk {
			logger.Debug("indexer: cannot retrieve either icon, executable paths or name")
			continue
		}

		db.Create(&models.DesktopEntry{
			Icon:           iconPath,
			Name:           name,
			EntryPath:      desktopEntryPath,
			ExecutablePath: executablePath,
		})
		continue
	}
}

func retrieveParamsFromEntries(data []string) map[string]string {
	params := map[string]string{}

	for _, line := range data {
		// if we have a space somewhere
		// then we are entering another desktop entry so stop
		if line == "" || line == " " {
			break
		}

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
