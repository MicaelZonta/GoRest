package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

func AtualizarTarefaGateway(uC *loghelper.Usercontext, codigo int64, tM model.TarefaModel) (int64, error) {

	loghelper.LogDebug(uC, loghelper.D00005)
	tE := adapter.TarefaEntityFromTarefaModel(tM)

	linhasAtualizadas, err := postgres.AtualizarTarefaRepository(uC, codigo, tE)
	loghelper.LogDebug(uC, loghelper.D00006)

	return linhasAtualizadas, err
}
