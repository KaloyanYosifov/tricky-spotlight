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

	localWidgets.InitAppController()
	window.SetCentralWidget(localWidgets.GetAppController().GetCentralWidget())

	// launch a goroutine to handle the controller renders
	go func(controller *localWidgets.AppController) {
		for app != nil && !app.ClosingDown() {
			controller.Render()
		}
	}(localWidgets.GetAppController())

	initWidgets(window)

	return &window
}

func initWidgets(window Window) {
	inputController := localWidgets.NewInputController2("input-1")
	localWidgets.
		GetAppController().
		AddController2("input-controller-1", inputController, 0, core.Qt__AlignTop)

	listModel := localWidgets.NewListModel("list-1")
	listController := localWidgets.NewListController(listModel, localWidgets.NewListView())
	listController.Hide()

	localWidgets.
		GetAppController().
		AddController2("list-controller-1", listController, 0, core.Qt__AlignBottom)

	inputController.GetInput().ConnectTextChanged(func(text string) {
		if len(text) > 2 {
			listController.Show()
		} else {
			listController.Hide()
		}
	})
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

	window.SetStyleSheet("background-color: #1a2138")
}
