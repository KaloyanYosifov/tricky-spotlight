package widgets

type listModel struct {
	*BaseModel
	text string
}

type listModelManager struct {
	models map[string]*listModel
}

var listManager *listModelManager

func GetListModelManager() *listModelManager {
	if listManager == nil {
		listManager = &listModelManager{make(map[string]*listModel, 0)}
	}

	return listManager
}

func NewListModel(modelId string) *listModel {
	listModel := &listModel{
		&BaseModel{},
		"",
	}
	GetListModelManager().addModel(modelId, listModel)

	return listModel
}

func (im *listModel) SetText(text string) *listModel {
	im.text = text

	return im
}

func (im *listModel) GetText() string {
	return im.text
}

func (lmm *listModelManager) addModel(key string, model *listModel) *listModelManager {
	_, ok := lmm.models[key]

	if ok {
		panic("input model manager: there is already a model with such key" + key + "!")
	}

	lmm.models[key] = model

	return lmm
}

func (lmm *listModelManager) GetModel(key string) *listModel {
	return lmm.models[key]
}
