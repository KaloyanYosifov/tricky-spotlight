package widgets

import (
	"github.com/therecipe/qt/widgets"
)

func InitMainWidgets() *widgets.QWidget {
	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetStyleSheet("background-color: #222b45;")

	//layout.AddWidget(input, 0, core.Qt__AlignTop)

	widget.SetLayout(layout)

	return widget
}
