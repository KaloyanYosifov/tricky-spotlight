package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"sync"
)

type AppController struct {
	widget      *widgets.QWidget
	layout      *widgets.QVBoxLayout
	controllers *sync.Map
}

var appController *AppController

func InitAppController() *AppController {
	if appController != nil {
		return GetAppController()
	}

	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	appController = &AppController{widget, layout, &sync.Map{}}

	return appController
}

func GetAppController() *AppController {
	if appController == nil {
		return InitAppController()
	}

	return appController
}

func (ac *AppController) AddController(id string, controller WidgetController) *AppController {
	_, ok := ac.controllers.Load(id)

	if ok {
		panic("There is already a controller with such id")
	}

	ac.controllers.Store(id, controller)
	ac.widget.Layout().AddWidget(controller.getView())

	return ac
}

func (ac *AppController) AddController2(id string, controller WidgetController, stretch int, alignement core.Qt__AlignmentFlag) *AppController {
	_, ok := ac.controllers.Load(id)

	if ok {
		panic("There is already a controller with such id")
	}

	ac.controllers.Store(id, controller)
	ac.layout.AddWidget(controller.getView(), stretch, alignement)

	return ac
}

func (ac *AppController) Render() *AppController {
	ac.controllers.Range(func(_ interface{}, controller interface{}) bool {
		if controller == nil {
			return false
		}

		controller.(WidgetController).render()

		return true
	})

	return ac
}

func (ac *AppController) GetCentralWidget() widgets.QWidget_ITF {
	return ac.widget
}

func (ac *AppController) GetController(id string) WidgetController {
	controller, ok := ac.controllers.Load(id)

	if !ok {
		panic("app controller: no controller is found with such id - " + id)
	}

	return controller.(WidgetController)
}
