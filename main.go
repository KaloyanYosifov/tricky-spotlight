package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/logger"
	"github.com/therecipe/qt/core"
	"os"

	"github.com/therecipe/qt/widgets"
)

var env string

func main() {
	if env == "" {
		env = "production"
	}

	logger.Init(env)

	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := initMainWindow(app)

	db := initDatabase(core.QStandardPaths_WritableLocation(core.QStandardPaths__AppConfigLocation))
	defer db.Close()
	migrateModels(db)

	indexDesktopEntries(db)

	window.initKeyEventHandling()

	app.Exec()
}
