package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

func CriarTarefaUsecase(uC *loghelper.Usercontext, t model.TarefaModel) (model.TarefaModel, error) {

	loghelper.LogDebug(uC, loghelper.D00004)
	t, err := gateway.SalvarTarefaGateway(uC, t)
	loghelper.LogDebug(uC, loghelper.D00004)

	return t, err
}
