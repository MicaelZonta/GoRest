package usecase

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type CriarTarefaUseCase interface {
	Init(uC loghelper.UserContext) CriarTarefaUseCase
	CriarTarefaExecute(t model.TarefaModel) (model.TarefaModel, error)
}

type CriarTarefaUseCaseImpl struct {
	uC                 loghelper.UserContext
	CriarTarefaGateway gateway.CriarTarefaGateway
}

func (_this CriarTarefaUseCaseImpl) Init(uC loghelper.UserContext) CriarTarefaUseCase {
	_this.CriarTarefaGateway = gateway.CriarTarefaGatewayImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this CriarTarefaUseCaseImpl) CriarTarefaExecute(t model.TarefaModel) (model.TarefaModel, error) {

	_this.uC.LogDebug(loghelper.D00004)
	t, err := _this.CriarTarefaGateway.CriarTarefaExecute(t)
	_this.uC.LogDebug(loghelper.D00004)

	return t, err
}
