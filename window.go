package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	localWidgets "github.com/KaloyanYosifov/tricky-spotlight/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	*widgets.QMainWindow
}

func initMainWindow(app *widgets.QApplication) *Window {
	screenData := app.Desktop().ScreenGeometry(0)
	window := Window{widgets.NewQMainWindow(nil, 0)}
	window.SetMinimumSize2(screenData.Width()/2, 70)
	window.SetWindowTitle("Tricky Spotlight")

	// Center window
	x := (screenData.Width() - window.Width()) / 2
	y := (screenData.Height() - window.Height()) / 2
	window.Move2(x, y)

	window.initAttributes()
	window.SetCentralWidget(localWidgets.InitMainWidgets())

	return &window
}

func (window *Window) initKeyEventHandling() {
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
}

func (window *Window) initAttributes() {
	window.SetAttribute(core.Qt__WA_AlwaysStackOnTop, true)
	window.SetWindowFlag(core.Qt__WindowStaysOnTopHint, true)
	window.SetWindowFlag(core.Qt__Dialog|core.Qt__MSWindowsFixedSizeDialogHint, true)
	window.SetWindowFlag(core.Qt__FramelessWindowHint, true)
}
