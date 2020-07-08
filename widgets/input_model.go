package widgets

type inputModel struct {
	*BaseModel
	text string
}

type inputModelManager struct {
	models map[string]*inputModel
}

var modelManager *inputModelManager

func GetModelManager() *inputModelManager {
	if modelManager == nil {
		modelManager = &inputModelManager{make(map[string]*inputModel, 0)}
	}

	return modelManager
}

func NewInputModel(modelId string) *inputModel {
	inputModel := &inputModel{
		&BaseModel{},
		"",
	}
	GetModelManager().addModel(modelId, inputModel)

	return inputModel
}

func (im *inputModel) SetText(text string) *inputModel {
	im.text = text

	return im
}

func (im *inputModel) GetText() string {
	return im.text
}

func (imm *inputModelManager) addModel(key string, model *inputModel) *inputModelManager {
	_, ok := imm.models[key]

	if ok {
		panic("input model manager: there is already a model with such key" + key + "!")
	}

	imm.models[key] = model

	return imm
}

func (imm *inputModelManager) GetModel(key string) *inputModel {
	return imm.models[key]
}
