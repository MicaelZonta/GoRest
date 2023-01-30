package mock

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type ListarTarefaGatewayMock struct {
	ListarTarefasModels []model.TarefaModel
	Err                 error
}

func (_this ListarTarefaGatewayMock) Init(uC loghelper.UserContext) gateway.ListarTarefaGateway {
	return _this
}

func (_this ListarTarefaGatewayMock) ListarTarefasExecute() ([]model.TarefaModel, error) {
	return _this.ListarTarefasModels, _this.Err
}
