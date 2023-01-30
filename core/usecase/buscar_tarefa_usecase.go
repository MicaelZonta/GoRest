package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type BuscarTarefaUseCase interface {
	Init(uC loghelper.UserContext) BuscarTarefaUseCase
	BuscarTarefaExecute(codigo int64) (model.TarefaModel, error)
}

type BuscarTarefaUseCaseImpl struct {
	uC                  loghelper.UserContext
	BuscarTarefaGateway gateway.BuscarTarefaGateway
}

func (_this BuscarTarefaUseCaseImpl) Init(uC loghelper.UserContext) BuscarTarefaUseCase {
	_this.BuscarTarefaGateway = gateway.BuscarTarefaGatewayImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this BuscarTarefaUseCaseImpl) BuscarTarefaExecute(codigo int64) (model.TarefaModel, error) {

	_this.uC.LogDebug(loghelper.D00003)
	t, err := _this.BuscarTarefaGateway.BuscarTarefaExecute(codigo)
	_this.uC.LogDebug(loghelper.D00004)

	return t, err
}
