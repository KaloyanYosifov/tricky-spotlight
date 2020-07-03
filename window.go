package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	activeWindow *widgets.QMainWindow
}

func initMainWindow() *Window {
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

	return &Window{
		activeWindow: window,
	}
}

func (window *Window) initKeyEventHandling() {
	keyEventHandler := keylogger.NewKeyEventHandler(func(eventHandler *keylogger.KeyEventHandler) {
		if eventHandler.IsKeyCombinationActive([]keylogger.GlobalKey{keylogger.KEY_SPACE, keylogger.KEY_CTRL}) {
			if window.activeWindow.IsVisible() {
				window.activeWindow.Hide()
			} else {
				window.activeWindow.Show()
			}
		}
	}, func(eventHandler *keylogger.KeyEventHandler) {

	})
	go keylogger.WaitForKeyEvents(keyEventHandler)
}
