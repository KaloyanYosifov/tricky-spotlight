package widgets

import (
	"github.com/therecipe/qt/widgets"
)

type ListController struct {
	model *listModel
	list  *listView
}

func NewListController(model *listModel, list *listView) *ListController {
	list.SetModel(model.abstractListModel)

	listController := &ListController{model, list}

	return listController
}

func NewListController2(modelId string) *ListController {
	return NewListController(NewListModel(modelId), NewListView())
}

func (lc *ListController) render() {
	if !lc.model.isDueForAnUpdate() {
		return
	}

	lc.model.Updated()
}

func (lc *ListController) getModel() WidgetModel {
	return lc.model
}

func (lc *ListController) getView() widgets.QWidget_ITF {
	return lc.list
}

func (lc *ListController) Show() {
	lc.list.Show()
}

func (lc *ListController) Hide() {
	lc.list.Hide()
}

func (lc *ListController) GetList() *listView {
	return lc.list
}

func (lc *ListController) GetModel() *listModel {
	return lc.model
}
