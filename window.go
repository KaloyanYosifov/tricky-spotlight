package main

import (
	"github.com/KaloyanYosifov/tricky-spotlight/database"
	"github.com/KaloyanYosifov/tricky-spotlight/keylogger"
	"github.com/KaloyanYosifov/tricky-spotlight/models"
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

	initWidgets()

	return &window
}

func initWidgets() {
	inputController := localWidgets.NewInputController2("input-1")
	localWidgets.
		GetAppController().
		AddController2("input-controller-1", inputController, 0, core.Qt__AlignTop)

	listModel := localWidgets.NewListModel("list-1")
	listController := localWidgets.NewListController(listModel, localWidgets.NewListView())

	localWidgets.
		GetAppController().
		AddController2("list-controller-1", listController, 0, core.Qt__AlignBottom)

	inputController.GetInput().ConnectTextChanged(func(text string) {
		entries := models.SearchForDesktopEntry(text, database.GetDatabase().GetUnderilyingDB())
		listModel.Clear()

		for _, entry := range entries {
			listModel.Add(localWidgets.ListData{Icon: entry.Icon, Name: entry.Name, Executable: entry.ExecutablePath})
		}
	})
}

func (window *Window) initKeyEventHandling() {
	keyEventHandler := keylogger.NewKeyEventHandler(func(eventHandler *keylogger.KeyEventHandler) {
		if window.IsVisible() && eventHandler.IsOnlyKeyActive(keylogger.KEY_ESC) {
			window.Hide()
		}

		if window.IsVisible() && eventHandler.IsOnlyKeyActive(keylogger.KEY_ENTER) {
			listController := localWidgets.GetAppController().GetController("list-controller-1").(*localWidgets.ListController)
			inputController := localWidgets.GetAppController().GetController("input-controller-1").(*localWidgets.InputController)
			selectedIndexes := listController.GetList().SelectedIndexes()
			var data *localWidgets.ListData

			if len(selectedIndexes) > 0 {
				data = listController.GetModel().GetItem(selectedIndexes[0].Row())
			} else {
				data = listController.GetModel().GetItem(0)
			}

			if data != nil {
				var entry models.DesktopEntry
				db := database.GetDatabase()
				db.First(&entry, "executable_path = ?", data.Executable)

				entry.Execute()
				listController.GetModel().Clear()
				inputController.GetInputModel().SetText("").Update()

				entry.TimesTriggered += 1
				db.Save(&entry)

				window.Hide()
			}
		}

		if eventHandler.IsOnlyKeyCombinationActive([]keylogger.GlobalKey{keylogger.KEY_SPACE, keylogger.KEY_CTRL}) {
			if window.IsVisible() {
				window.Hide()
			} else {
				window.Show()
				window.ActivateWindow()
				inputController := localWidgets.GetAppController().GetController("input-controller-1").(*localWidgets.InputController)
				inputController.GetInput().SetFocus2()
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
