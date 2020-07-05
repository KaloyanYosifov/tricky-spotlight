package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	activeWindow *widgets.QMainWindow
}

func initMainWindow(app *widgets.QApplication) *Window {
	screenData := app.Desktop().ScreenGeometry(0)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(screenData.Width()/2, 200)
	window.SetWindowTitle("Tricky Spotlight")

	// Center window
	x := (screenData.Width() - window.Width()) / 2
	y := (screenData.Height() - window.Height()) / 2
	window.Move2(x, y)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQGridLayout2())
	window.SetCentralWidget(widget)

	mainWindow := Window{
		activeWindow: window,
	}
	mainWindow.initAttributes()

	return &mainWindow
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

func (window *Window) initAttributes() {
	window.activeWindow.SetAttribute(core.Qt__WA_AlwaysStackOnTop, true)
	window.activeWindow.SetWindowFlag(core.Qt__WindowStaysOnTopHint, true)
	window.activeWindow.SetWindowFlag(core.Qt__Dialog|core.Qt__MSWindowsFixedSizeDialogHint, true)
	window.activeWindow.SetWindowFlag(core.Qt__FramelessWindowHint, true)
}
