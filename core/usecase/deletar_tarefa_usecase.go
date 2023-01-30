package usecase

import (
	"GoRest/core/gateway"
	"GoRest/crosscutting/loghelper"
)

type DeletarTarefaUsecase interface {
	Init(uC loghelper.UserContext) DeletarTarefaUsecase
	DeletarTarefaExecute(codigo int64) (int64, error)
}

type DeletarTarefaUsecaseImpl struct {
	uC                   loghelper.UserContext
	DeletarTarefaGateway gateway.DeletarTarefaGateway
}

func (_this DeletarTarefaUsecaseImpl) Init(uC loghelper.UserContext) DeletarTarefaUsecase {
	_this.DeletarTarefaGateway = gateway.DeletarTarefaGatewayImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this DeletarTarefaUsecaseImpl) DeletarTarefaExecute(codigo int64) (int64, error) {

	_this.uC.LogDebug(loghelper.D00003)
	t, err := _this.DeletarTarefaGateway.DeletarTarefaExecute(codigo)
	_this.uC.LogDebug(loghelper.D00004)

	return t, err
}
