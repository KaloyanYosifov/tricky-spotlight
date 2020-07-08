package widgets

type InputModel struct {
	*BaseModel
	text string
}

func (im *InputModel) SetText(text string) *InputModel {
	im.text = text

	return im
}

func (im *InputModel) GetText() string {
	return im.text
}

func (im *InputModel) Update() {
	im.BaseModel.Update()
}
