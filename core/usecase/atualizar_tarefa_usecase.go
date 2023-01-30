package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type AtualizarTarefaUsecase interface {
	Init(uC loghelper.UserContext) AtualizarTarefaUsecase
	AtualizarTarefaExecute(codigo int64, tarefaModel model.TarefaModel) (int64, error)
}

type AtualizarTarefaUsecaseImpl struct {
	uC                     loghelper.UserContext
	AtualizarTarefaGateway gateway.AtualizarTarefaGateway
}

func (_this AtualizarTarefaUsecaseImpl) Init(uC loghelper.UserContext) AtualizarTarefaUsecase {
	_this.AtualizarTarefaGateway = gateway.AtualizarTarefaGatewayImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this AtualizarTarefaUsecaseImpl) AtualizarTarefaExecute(codigo int64, tarefaModel model.TarefaModel) (int64, error) {

	_this.uC.LogDebug(loghelper.D00003)
	linhasAtualizadas, err := _this.AtualizarTarefaGateway.AtualizarTarefaExecute(codigo, tarefaModel)
	_this.uC.LogDebug(loghelper.D00004)

	return linhasAtualizadas, err
}
