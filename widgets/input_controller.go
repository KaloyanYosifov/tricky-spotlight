package widgets

import "github.com/therecipe/qt/widgets"

type InputController struct {
	model *InputModel
	input *InputView
}

func NewInputController(model *InputModel, input *InputView) *InputController {
	inputController := &InputController{model, input}

	return inputController
}

func NewInputController2() *InputController {
	return NewInputController(&InputModel{}, NewInputView())
}

func (ic *InputController) render() {
	if !ic.model.isDueForAnUpdate() {
		return
	}

	ic.input.SetText(ic.model.text)
	ic.model.Updated()
}

func (ic *InputController) getModel() WidgetModel {
	return ic.model
}

func (ic *InputController) getView() widgets.QWidget_ITF {
	return ic.input
}