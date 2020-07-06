package widgets

import "github.com/therecipe/qt/widgets"

type InputView struct {
	*widgets.QLineEdit
}

func NewInputView() *InputView {
	input := &InputView{widgets.NewQLineEdit(nil)}
	input.SetPlaceholderText("Look for something nice")
	input.SetFixedHeight(50)
	input.SetStyleSheet("background-color: #1a2138; " +
		"color: #ffffff; padding: 10px;" +
		"font-size: 22px" +
		"")

	return input
}
