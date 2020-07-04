package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := initMainWindow(app)

	db := initDatabase(core.QStandardPaths_WritableLocation(core.QStandardPaths__AppConfigLocation))
	defer db.Close()

	migrateModels(db)

	window.initKeyEventHandling()

	app.Exec()

	fmt.Println('a')
}
