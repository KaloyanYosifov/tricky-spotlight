package widgets

import (
	"github.com/therecipe/qt/widgets"
)

type listController struct {
	model *listModel
	list  *listView
}

func NewListController(model *listModel, list *listView) *listController {
	list.SetModel(model.abstractListModel)

	listController := &listController{model, list}

	return listController
}

func NewListController2(modelId string) *listController {
	return NewListController(NewListModel(modelId), NewListView())
}

func (lc *listController) render() {
	if !lc.model.isDueForAnUpdate() {
		return
	}

	lc.model.Updated()
}

func (lc *listController) getModel() WidgetModel {
	return lc.model
}

func (lc *listController) getView() widgets.QWidget_ITF {
	return lc.list
}

func (lc *listController) Show() {
	lc.list.Show()
}

func (lc *listController) Hide() {
	lc.list.Hide()
}
