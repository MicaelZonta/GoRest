package gateway

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
)

func DeletarTarefaGateway(uC *loghelper.Usercontext, codigo int64) (int64, error) {

	loghelper.LogDebug(uC, loghelper.D00006)
	res, err := postgres.DeletarTarefaRepository(uC, codigo)
	loghelper.LogDebug(uC, loghelper.D00006)

	return res, err
}
