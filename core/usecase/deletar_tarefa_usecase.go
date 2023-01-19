package usecase

import (
	"GoRest/core/gateway"
	"GoRest/crosscutting/loghelper"
)

func DeletarTarefaUsecase(uC *loghelper.Usercontext, codigo int64) (int64, error) {

	loghelper.LogDebug(uC, loghelper.D00003)
	t, err := gateway.DeletarTarefaGateway(uC, codigo)
	loghelper.LogDebug(uC, loghelper.D00004)

	return t, err
}
