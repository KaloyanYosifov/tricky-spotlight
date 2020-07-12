package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type AppController struct {
	widget      *widgets.QWidget
	layout      *widgets.QVBoxLayout
	controllers []WidgetController
}

var appController *AppController

func InitAppController() *AppController {
	if appController != nil {
		return GetAppController()
	}

	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	appController = &AppController{widget, layout, make([]WidgetController, 0, 5)}

	return appController
}

func GetAppController() *AppController {
	if appController == nil {
		return InitAppController()
	}

	return appController
}

func (ac *AppController) AddController(controller WidgetController) *AppController {
	ac.controllers = append(ac.controllers, controller)
	ac.widget.Layout().AddWidget(controller.getView())

	return ac
}

func (ac *AppController) AddController2(controller WidgetController, stretch int, alignement core.Qt__AlignmentFlag) *AppController {
	ac.controllers = append(ac.controllers, controller)
	ac.layout.AddWidget(controller.getView(), stretch, alignement)

	return ac
}

func (ac *AppController) Render() *AppController {
	for _, controller := range ac.controllers {
		if controller == nil {
			continue
		}

		controller.render()
	}

	return ac
}

func (ac *AppController) GetCentralWidget() widgets.QWidget_ITF {
	return ac.widget
}
