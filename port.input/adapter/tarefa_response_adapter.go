package adapter

import (
	"GoRest/core/model"
	"GoRest/port.input/transferobject/request"
	"GoRest/port.input/transferobject/response"
)

func TarefaResponseFromTarefaModel(tM model.TarefaModel) (tR response.TarefaResponse) {

	tR.Codigo = tM.Codigo
	tR.Titulo = tM.Titulo
	tR.Descricao = tM.Titulo
	tR.Completa = tM.Completa

	return tR
}

func TarefaResponseSliceFromTarefaModelSlice(tMSlice []model.TarefaModel) (tRSlice []response.TarefaResponse) {

	for _, tM := range tMSlice {
		tRSlice = append(tRSlice, TarefaResponseFromTarefaModel(tM))
	}

	return tRSlice
}

func TarefaResponseFromTarefaRequest(tRequest request.TarefaRequest) (tResp response.TarefaResponse) {

	tResp.Codigo = tRequest.Codigo
	tResp.Titulo = tRequest.Titulo
	tResp.Descricao = tRequest.Titulo
	tResp.Completa = tRequest.Completa

	return tResp
}
