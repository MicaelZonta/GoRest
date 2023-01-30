package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/adapter"
)

type AtualizarTarefaGateway interface {
	Init(uC loghelper.UserContext) AtualizarTarefaGateway
	AtualizarTarefaExecute(codigo int64, tM model.TarefaModel) (int64, error)
}

type AtualizarTarefaGatewayImpl struct {
	TarefaRepository postgres.TarefaRepository
	uC               loghelper.UserContext
}

func (_this AtualizarTarefaGatewayImpl) Init(uC loghelper.UserContext) AtualizarTarefaGateway {
	_this.TarefaRepository = postgres.TarefaRepositoryImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this AtualizarTarefaGatewayImpl) AtualizarTarefaExecute(codigo int64, tM model.TarefaModel) (int64, error) {

	_this.uC.LogDebug(loghelper.D00005)
	tE := adapter.TarefaEntityFromTarefaModel(tM)
	linhasAtualizadas, err := _this.TarefaRepository.AtualizarTarefaExecute(codigo, tE)
	_this.uC.LogDebug(loghelper.D00006)

	return linhasAtualizadas, err
}
