package widgets

import "github.com/therecipe/qt/widgets"

type WidgetModel interface {
	Update()
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

func (bm *BaseModel) isDueForAnUpdate() bool {
	return bm.shouldUpdate
}
