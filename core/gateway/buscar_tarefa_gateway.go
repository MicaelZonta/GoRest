package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

type BuscarTarefaGateway interface {
	Init(uC loghelper.UserContext) BuscarTarefaGateway
	BuscarTarefaExecute(codigo int64) (model.TarefaModel, error)
}

type BuscarTarefaGatewayImpl struct {
	TarefaRepository postgres.TarefaRepository
	uC               loghelper.UserContext
}

func (_this BuscarTarefaGatewayImpl) Init(uC loghelper.UserContext) BuscarTarefaGateway {
	_this.TarefaRepository = postgres.TarefaRepositoryImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this BuscarTarefaGatewayImpl) BuscarTarefaExecute(codigo int64) (model.TarefaModel, error) {

	_this.uC.LogDebug(loghelper.D00005)
	tE, err := _this.TarefaRepository.BuscarTarefaExecute(codigo)
	_this.uC.LogDebug(loghelper.D00006)

	return adapter.TarefaModelFromTarefaEntity(tE), err
}
