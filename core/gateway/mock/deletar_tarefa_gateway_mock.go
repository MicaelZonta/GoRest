package mock

import (
	"GoRest/core/gateway"
	"GoRest/crosscutting/loghelper"
)

type DeletarTarefaGatewayMock struct {
	LinhasDeletadas int64
	Err             error
}

func (_this DeletarTarefaGatewayMock) Init(uC loghelper.UserContext) gateway.DeletarTarefaGateway {
	return _this
}

func (_this DeletarTarefaGatewayMock) DeletarTarefaExecute(codigo int64) (int64, error) {
	return _this.LinhasDeletadas, _this.Err
}
