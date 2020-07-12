package widgets

import "github.com/therecipe/qt/widgets"

type listView struct {
	*widgets.QListView
}

func NewListView() *listView {
	input := &listView{widgets.NewQListView(nil)}

	return input
}
