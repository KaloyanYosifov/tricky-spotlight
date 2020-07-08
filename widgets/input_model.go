package widgets

type InputModel struct {
	*BaseModel
	text string
}

func NewInputModel() *InputModel {
	return &InputModel{
		&BaseModel{},
		"",
	}
}

func (im *InputModel) SetText(text string) *InputModel {
	im.text = text

	return im
}

func (im *InputModel) GetText() string {
	return im.text
}
