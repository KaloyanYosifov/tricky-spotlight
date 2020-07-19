package models

import (
	"bytes"
	"github.com/KaloyanYosifov/tricky-spotlight/logger"
	"github.com/jinzhu/gorm"
	"os/exec"
	"strings"
)

type DesktopEntry struct {
	gorm.Model
	Name           string `gorm:"index"`
	Icon           string
	EntryPath      string `gorm:"index"`
	ExecutablePath string
	TimesTriggered int `gorm:"index"`
}

func SearchForDesktopEntry(name string, db *gorm.DB) []DesktopEntry {
	var desktopEntries []DesktopEntry
	db.Where("name like ?", "%"+name+"%").Order("times_triggered desc").Find(&desktopEntries)

	return desktopEntries
}

func (de *DesktopEntry) Execute() {
	userToRunCommandAs := findOutAsWhoToExecuteEntryAs()
	cmd := exec.Command("runuser", "-l", userToRunCommandAs, "-c", "export DISPLAY=:0 && "+de.ExecutablePath+" </dev/null &>/dev/null &")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		panic("Something went wrong when trying to run the program")
	}

	logger.Debug("Executed command " + de.ExecutablePath)
}

func findOutAsWhoToExecuteEntryAs() string {
	wCmd := exec.Command("w", "-h")
	var output bytes.Buffer
	wCmd.Stdout = &output

	if err := wCmd.Run(); err != nil {
		logger.Debug("desktop_entry: cannot run \"w\" command!")
		return "root"
	}

	awkCmd := exec.Command("awk", "{print $1}")
	awkCmd.Stdin = strings.NewReader(output.String())
	var output2 bytes.Buffer
	awkCmd.Stdout = &output2

	if err := awkCmd.Run(); err != nil {
		logger.Debug("desktop_entry: cannot run \"awk\" command!")
		return "root"
	}

	users := strings.Split(output2.String(), "\n")

	for _, user := range users {
		if user != "root" {
			return user
		}
	}

	return "root"
}
