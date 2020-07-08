package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	localWidgets "github.com/KaloyanYosifov/tricky-spotlight/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"strconv"
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

	localWidgets.InitAppController()
	window.SetCentralWidget(localWidgets.GetAppController().GetCentralWidget())

	// launch a goroutine to handle the controller renders
	go func(controller *localWidgets.AppController) {
		for app != nil && !app.ClosingDown() {
			controller.Render()
		}
	}(localWidgets.GetAppController())

	initWidgets()

	return &window
}

func initWidgets() {
	localWidgets.GetAppController().AddController(localWidgets.NewInputController2("input-1"))
	localWidgets.GetAppController().AddController(localWidgets.NewInputController2("input-2"))
}

func (window *Window) initKeyEventHandling() {
	keyEventHandler := keylogger.NewKeyEventHandler(func(eventHandler *keylogger.KeyEventHandler) {
		if eventHandler.IsKeyActive(keylogger.KEY_a) {
			m := localWidgets.GetModelManager().GetModel("input-1")
			m.SetText("testingggg" + strconv.Itoa(rand.Int()))
			m.Update()
		}

		if eventHandler.IsKeyActive(keylogger.KEY_b) {
			m := localWidgets.GetModelManager().GetModel("input-2")
			m.SetText("test" + strconv.Itoa(rand.Int()))
			m.Update()
		}

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

	window.SetStyleSheet("background-color: #1a2138")
}
