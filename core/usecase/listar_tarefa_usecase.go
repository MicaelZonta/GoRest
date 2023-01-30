package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type ListarTarefaUseCase interface {
	Init(uC loghelper.UserContext) ListarTarefaUseCase
	ListarTarefasExecute() ([]model.TarefaModel, error)
}

type ListarTarefaCaseImpl struct {
	uC                  loghelper.UserContext
	ListarTarefaGateway gateway.ListarTarefaGateway
}

func (_this ListarTarefaCaseImpl) Init(uC loghelper.UserContext) ListarTarefaUseCase {
	_this.ListarTarefaGateway = gateway.ListarTarefaGatewayImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this ListarTarefaCaseImpl) ListarTarefasExecute() ([]model.TarefaModel, error) {
	_this.uC.SendValue("2", "2")
	_this.uC.LogDebug(loghelper.D00003)
	tMSlice, err := _this.ListarTarefaGateway.ListarTarefasExecute()
	_this.uC.LogDebug(loghelper.D00004)

	return tMSlice, err
}
