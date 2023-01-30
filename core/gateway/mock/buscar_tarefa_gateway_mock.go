package mock

import (
	"GoRest/core/gateway"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
)

type BuscarTarefaGatewayMock struct {
	BuscarTarefaModel model.TarefaModel
	Err               error
}

func (_this BuscarTarefaGatewayMock) Init(uC loghelper.UserContext) gateway.BuscarTarefaGateway {
	return _this
}

func (_this BuscarTarefaGatewayMock) BuscarTarefaExecute(codigo int64) (model.TarefaModel, error) {
	return _this.BuscarTarefaModel, _this.Err
}
