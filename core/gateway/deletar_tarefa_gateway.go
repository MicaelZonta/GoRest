package gateway

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
)

type DeletarTarefaGateway interface {
	Init(uC loghelper.UserContext) DeletarTarefaGateway
	DeletarTarefaExecute(codigo int64) (int64, error)
}

type DeletarTarefaGatewayImpl struct {
	TarefaRepository postgres.TarefaRepository
	uC               loghelper.UserContext
}

func (_this DeletarTarefaGatewayImpl) Init(uC loghelper.UserContext) DeletarTarefaGateway {
	_this.TarefaRepository = postgres.TarefaRepositoryImpl{}.Init(uC)
	_this.uC = uC
	return _this
}

func (_this DeletarTarefaGatewayImpl) DeletarTarefaExecute(c int64) (int64, error) {

	_this.uC.LogDebug(loghelper.D00006)
	res, err := _this.TarefaRepository.DeletarTarefaExecute(c)
	_this.uC.LogDebug(loghelper.D00006)

	return res, err
}
