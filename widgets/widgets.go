package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Layout struct {
	*widgets.QVBoxLayout
}

type Widget struct {
	*widgets.QWidget
}

func (w *Widget) AddController(controller Controller) *Widget {
	w.Layout().AddWidget(controller.getView())

	return w
}

func (l *Layout) AddController(controller Controller, stretch int, alignement core.Qt__AlignmentFlag) *Layout {
	l.AddWidget(controller.getView(), stretch, alignement)

	return l
}

func InitMainWidgets() *Widget {
	layout := &Layout{widgets.NewQVBoxLayout()}
	widget := &Widget{widgets.NewQWidget(nil, 0)}
	widget.SetStyleSheet("background-color: #222b45;")

	model := &InputModel{}
	input := NewInputView()
	inputController := NewInputController(model, input)
	layout.AddController(inputController, 0, core.Qt__AlignTop)

	widget.SetLayout(layout)

	inputController.render()

	model.SetText("tetesd")

	inputController.render()

	return widget
}
