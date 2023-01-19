package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

func BuscarTarefaGateway(uC *loghelper.Usercontext, codigo int64) (model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.D00005)
	tE, err := postgres.BuscarTarefaRepository(uC, codigo)
	loghelper.LogDebug(uC, loghelper.D00006)

	return adapter.TarefaModelFromTarefaEntity(tE), err
}
