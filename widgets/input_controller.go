package widgets

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
