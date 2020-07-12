package widgets

import "github.com/therecipe/qt/widgets"

type InputController struct {
	model *inputModel
	input *inputView
}

func NewInputController(model *inputModel, input *inputView) *InputController {
	inputController := &InputController{model, input}
	input.ConnectTextChanged(func(text string) {
		model.SetText(text)
	})

	return inputController
}

func NewInputController2(modelId string) *InputController {
	return NewInputController(NewInputModel(modelId), NewInputView())
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

func (ic *InputController) Show() {
	ic.input.Show()
}

func (ic *InputController) Hide() {
	ic.input.Hide()
}

func (ic *InputController) GetInput() *inputView {
	return ic.input
}

func (ic *InputController) GetInputModel() *inputModel {
	return ic.model
}
