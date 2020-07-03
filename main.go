package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := initMainWindow(app)
	window.initKeyEventHandling()

	app.Exec()
}
