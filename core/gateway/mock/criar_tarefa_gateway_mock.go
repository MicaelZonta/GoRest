package mock

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type CriarTarefaGatewayMock struct {
	Err error
}

func (_this CriarTarefaGatewayMock) Init(uC loghelper.UserContext) gateway.CriarTarefaGateway {
	return _this
}

func (_this CriarTarefaGatewayMock) CriarTarefaExecute(tM model.TarefaModel) (model.TarefaModel, error) {
	return tM, _this.Err
}
