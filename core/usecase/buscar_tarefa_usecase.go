package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

func BuscarTarefaUsecase(uC *loghelper.Usercontext, codigo int64) (model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.D00003)
	t, err := gateway.BuscarTarefaGateway(uC, codigo)
	loghelper.LogDebug(uC, loghelper.D00004)

	return t, err
}
