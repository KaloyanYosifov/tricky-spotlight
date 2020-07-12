package widgets

import "github.com/therecipe/qt/widgets"

type inputView struct {
	*widgets.QLineEdit
	onSubmitEvents []func()
}

func NewInputView() *inputView {
	input := &inputView{widgets.NewQLineEdit(nil), make([]func(), 0)}
	input.SetPlaceholderText("Look for something nice")
	input.SetFixedHeight(50)
	input.SetStyleSheet("background-color: #1a2138; " +
		"color: #ffffff; padding: 10px;" +
		"font-size: 22px" +
		"")

	return input
}
