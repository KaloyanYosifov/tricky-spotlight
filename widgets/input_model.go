package widgets

type InputModel struct {
	text string
}

func (im *InputModel) SetText(text string) *InputModel {
	im.text = text

	return im
}

func (im *InputModel) GetText() string {
	return im.text
}
