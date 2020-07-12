package widgets

import "github.com/therecipe/qt/widgets"

type listView struct {
	*widgets.QLineEdit
}

func NewListView() *listView {
	input := &listView{widgets.NewQLineEdit(nil)}
	input.SetPlaceholderText("Look for something nice")
	input.SetFixedHeight(50)
	input.SetStyleSheet("background-color: #1a2138; " +
		"color: #ffffff; padding: 10px;" +
		"font-size: 22px" +
		"")

	return input
}
