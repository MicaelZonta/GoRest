package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

type ListarTarefaGateway interface {
	Init(uC loghelper.UserContext) ListarTarefaGateway
	ListarTarefasExecute() ([]model.TarefaModel, error)
}

type ListarTarefaGatewayImpl struct {
	TarefaRepository postgres.TarefaRepository
	uC               loghelper.UserContext
}

func (_this ListarTarefaGatewayImpl) Init(uC loghelper.UserContext) ListarTarefaGateway {
	_this.TarefaRepository = postgres.TarefaRepositoryImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this ListarTarefaGatewayImpl) ListarTarefasExecute() ([]model.TarefaModel, error) {
	_this.uC.SendValue("3", "3")
	_this.uC.LogDebug(loghelper.D00005)
	tESlice, err := _this.TarefaRepository.ListarTarefasExecute()
	_this.uC.LogDebug(loghelper.D00006)

	return adapter.TarefaModelSliceFromTarefaEntitySlice(tESlice), err
}
