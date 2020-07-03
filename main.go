package main

import (
	"github.com/therecipe/qt/core"
	"os"

	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(1200, 900)
	window.SetWindowTitle("Tricky Spotlight")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQGridLayout2())
	window.SetCentralWidget(widget)
	window.SetAttribute(core.Qt__WA_AlwaysStackOnTop, true)
	window.SetWindowFlag(core.Qt__WindowStaysOnTopHint, true)

	keyEventHandler := keylogger.NewKeyEventHandler(func(eventHandler *keylogger.KeyEventHandler) {
		if eventHandler.IsKeyCombinationActive([]keylogger.GlobalKey{keylogger.KEY_SPACE, keylogger.KEY_CTRL}) {
			if window.IsVisible() {
				window.Hide()
			} else {
				window.Show()
			}
		}
	}, func(eventHandler *keylogger.KeyEventHandler) {

	})
	go keylogger.WaitForKeyEvents(keyEventHandler)

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}
