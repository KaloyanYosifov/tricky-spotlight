package widgets

import "github.com/therecipe/qt/widgets"

type WidgetModel interface {
	Update()
	Updated()
	isDueForAnUpdate() bool
}

type WidgetController interface {
	render()
	getModel() WidgetModel
	getView() widgets.QWidget_ITF
}

type BaseModel struct {
	shouldUpdate bool
}

func (bm *BaseModel) Update() {
	bm.shouldUpdate = true
}

func (bm *BaseModel) Updated() {
	bm.shouldUpdate = false
}

func (bm *BaseModel) isDueForAnUpdate() bool {
	return bm.shouldUpdate
}
