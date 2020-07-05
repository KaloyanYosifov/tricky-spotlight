package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func InitMainWidgets() *widgets.QWidget {
	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetStyleSheet("background-color: #222b45;")

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Look for something nice")
	input.SetFixedHeight(50)
	input.SetStyleSheet("background-color: #1a2138; color: #ffffff; padding: 10px;")
	layout.AddWidget(input, 0, core.Qt__AlignTop)

	widget.SetLayout(layout)

	return widget
}
