package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

func ListarTarefaUsecase(uC *loghelper.Usercontext) ([]model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.D00003)
	tMSlice, err := gateway.ListarTarefasGateway(uC)
	loghelper.LogDebug(uC, loghelper.D00004)

	return tMSlice, err
}
