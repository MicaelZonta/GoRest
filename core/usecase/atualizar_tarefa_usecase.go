package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

func AtualizarTarefaUsecase(uC *loghelper.Usercontext, codigo int64, tarefaModel model.TarefaModel) (int64, error) {

	loghelper.LogDebug(uC, loghelper.D00003)
	linhasAtualizadas, err := gateway.AtualizarTarefaGateway(uC, codigo, tarefaModel)
	loghelper.LogDebug(uC, loghelper.D00004)

	return linhasAtualizadas, err
}
