package widgets

import "github.com/therecipe/qt/widgets"

type listController struct {
	model *listModel
	list  *listView
}

func NewListController(model *listModel, list *listView) *listController {
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

	lc.list.SetText(lc.model.text)
	lc.model.Updated()
}

func (lc *listController) getModel() WidgetModel {
	return lc.model
}

func (lc *listController) getView() widgets.QWidget_ITF {
	return lc.list
}
