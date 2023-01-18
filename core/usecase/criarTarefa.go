package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

func CriarTarefaUsecase(uC *loghelper.Usercontext, t model.Tarefa) (model.Tarefa, error) {

	loghelper.LogDebug(uC, loghelper.I00003)
	t, err := gateway.SalvarTarefaGateway(uC, t)
	loghelper.LogDebug(uC, loghelper.I00004)

	return t, err
}
