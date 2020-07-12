package widgets

import "github.com/therecipe/qt/widgets"

type inputController struct {
	model *inputModel
	input *inputView
}

func NewInputController(model *inputModel, input *inputView) *inputController {
	inputController := &inputController{model, input}
	input.ConnectTextChanged(func(text string) {
		model.SetText(text)
	})

	return inputController
}

func NewInputController2(modelId string) *inputController {
	return NewInputController(NewInputModel(modelId), NewInputView())
}

func (ic *inputController) render() {
	if !ic.model.isDueForAnUpdate() {
		return
	}

	ic.input.SetText(ic.model.text)
	ic.model.Updated()
}

func (ic *inputController) getModel() WidgetModel {
	return ic.model
}

func (ic *inputController) getView() widgets.QWidget_ITF {
	return ic.input
}

func (ic *inputController) Show() {
	ic.input.Show()
}

func (ic *inputController) Hide() {
	ic.input.Hide()
}

func (ic *inputController) GetInput() *inputView {
	return ic.input
}

func (ic *inputController) GetInputModel() *inputModel {
	return ic.model
}
