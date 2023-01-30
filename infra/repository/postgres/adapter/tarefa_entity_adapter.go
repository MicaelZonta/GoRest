package adapter

import (
	"GoRest/core/model"
	"GoRest/infra/repository/postgres/entity"
)

func TarefaModelFromTarefaEntity(tE entity.TarefaEntity) (tM model.TarefaModel) {

	tM.Codigo = tE.Codigo
	tM.Titulo = tE.Titulo
	tM.Descricao = tE.Descricao
	tM.Completa = tE.Completa

	return tM
}

func TarefaEntityFromTarefaModel(tM model.TarefaModel) (tE entity.TarefaEntity) {
	tE.Codigo = tM.Codigo
	tE.Titulo = tM.Titulo
	tE.Descricao = tM.Descricao
	tE.Completa = tM.Completa

	return tE
}

func TarefaModelSliceFromTarefaEntitySlice(tESlice []entity.TarefaEntity) (tMSlice []model.TarefaModel) {

	for _, tE := range tESlice {
		tMSlice = append(tMSlice, TarefaModelFromTarefaEntity(tE))
	}

	return tMSlice
}
