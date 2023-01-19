package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

func SalvarTarefaGateway(uC *loghelper.Usercontext, tM model.TarefaModel) (model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.I00007)
	tE := adapter.TarefaEntityFromTarefaModel(tM)

	tE, err := postgres.InserirTarefaRepository(uC, tE)
	loghelper.LogDebug(uC, loghelper.I00008)

	return adapter.TarefaModelFromTarefaEntity(tE), err
}
