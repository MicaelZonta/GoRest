package mock

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type AtualizarTarefaGatewayMock struct {
	LinhasAtualizadas int64
	Err               error
}

func (_this AtualizarTarefaGatewayMock) Init(uC loghelper.UserContext) gateway.AtualizarTarefaGateway {
	return _this
}

func (_this AtualizarTarefaGatewayMock) AtualizarTarefaExecute(codigo int64, tM model.TarefaModel) (int64, error) {
	return _this.LinhasAtualizadas, _this.Err
}
