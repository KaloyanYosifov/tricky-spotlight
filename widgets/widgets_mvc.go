package widgets

import "github.com/therecipe/qt/widgets"

type Model interface {
	init()
}

type Controller interface {
	render()
	getModel() Model
	getView() widgets.QWidget_ITF
}
