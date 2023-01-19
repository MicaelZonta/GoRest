package adapter

import (
	"GoRest/core/model"
	"GoRest/port.input/transferobject/request"
)

func TarefaModelFromTarefaRequest(tR request.TarefaRequest) (tM model.TarefaModel) {

	tM.Codigo = tR.Codigo
	tM.Titulo = tR.Titulo
	tM.Descricao = tR.Titulo
	tM.Completa = tR.Completa

	return tM
}
