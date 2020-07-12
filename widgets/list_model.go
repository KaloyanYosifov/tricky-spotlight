package widgets

import "github.com/therecipe/qt/core"

type ListData struct {
	Icon       string
	Name       string
	Executable string
}

type listModel struct {
	*BaseModel
	modelData         []ListData
	abstractListModel *core.QAbstractListModel
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
		BaseModel:         &BaseModel{},
		modelData:         make([]ListData, 0, 0),
		abstractListModel: core.NewQAbstractListModel(nil),
	}

	listModel.abstractListModel.ConnectRowCount(listModel.rowCount)
	listModel.abstractListModel.ConnectData(listModel.data)

	GetListModelManager().addModel(modelId, listModel)

	return listModel
}

func (im *listModel) rowCount(*core.QModelIndex) int {
	return len(im.modelData)
}

func (im *listModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := im.modelData[index.Row()]
	return core.NewQVariant1(item.Name)
}

func (im *listModel) Add(item ListData) {
	im.abstractListModel.BeginInsertRows(core.NewQModelIndex(), len(im.modelData), len(im.modelData))
	im.modelData = append(im.modelData, item)
	im.abstractListModel.EndInsertRows()
}

func (im *listModel) RemoveLastItem() {
	if len(im.modelData) == 0 {
		return
	}
	im.abstractListModel.BeginRemoveRows(core.NewQModelIndex(), len(im.modelData)-1, len(im.modelData)-1)
	im.modelData = im.modelData[:len(im.modelData)-1]
	im.abstractListModel.EndRemoveRows()
}

func (im *listModel) GetData() []ListData {
	if len(im.modelData) <= 0 {
		return im.modelData
	}

	copiedListData := make([]ListData, 0, len(im.modelData))

	for _, listData := range im.modelData {
		copiedListData = append(copiedListData, listData)
	}

	return im.modelData
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
