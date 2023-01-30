package mock

import (
	"GoRest/core/model"
	"GoRest/core/usecase"
	"GoRest/crosscutting/loghelper"
)

type BuscarTarefaUsecaseMock struct {
	BuscarTarefaModel model.TarefaModel
	Err               error
}

func (_this BuscarTarefaUsecaseMock) Init(uC loghelper.UserContext) usecase.BuscarTarefaUseCase {
	return _this
}

func (_this BuscarTarefaUsecaseMock) BuscarTarefaExecute(codigo int64) (model.TarefaModel, error) {
	return _this.BuscarTarefaModel, _this.Err
}
