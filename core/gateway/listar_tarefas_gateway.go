package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

func ListarTarefasGateway(uC *loghelper.Usercontext) ([]model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.D00005)
	tESlice, err := postgres.ListarTarefasRepository(uC)
	loghelper.LogDebug(uC, loghelper.D00006)

	return adapter.TarefaModelSliceFromTarefaEntitySlice(tESlice), err
}
