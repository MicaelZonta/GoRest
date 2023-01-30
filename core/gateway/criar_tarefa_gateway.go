package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

type CriarTarefaGateway interface {
	Init(uC loghelper.UserContext) CriarTarefaGateway
	CriarTarefaExecute(tM model.TarefaModel) (model.TarefaModel, error)
}

type CriarTarefaGatewayImpl struct {
	TarefaRepository postgres.TarefaRepository
	uC               loghelper.UserContext
}

func (_this CriarTarefaGatewayImpl) Init(uC loghelper.UserContext) CriarTarefaGateway {
	_this.TarefaRepository = postgres.TarefaRepositoryImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this CriarTarefaGatewayImpl) CriarTarefaExecute(tM model.TarefaModel) (model.TarefaModel, error) {

	_this.uC.LogDebug(loghelper.I00007)
	tE := adapter.TarefaEntityFromTarefaModel(tM)

	tE, err := _this.TarefaRepository.InserirTarefaExecute(tE)
	_this.uC.LogDebug(loghelper.I00008)

	return adapter.TarefaModelFromTarefaEntity(tE), err
}
