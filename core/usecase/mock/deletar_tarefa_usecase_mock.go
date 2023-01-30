package mock

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/loghelper"
)

type DeletarTarefaUsecaseMock struct {
	LinhasDeletadas int64
	Err             error
}

func (_this DeletarTarefaUsecaseMock) Init(uC loghelper.UserContext) usecase.DeletarTarefaUsecase {
	return _this
}

func (_this DeletarTarefaUsecaseMock) DeletarTarefaExecute(codigo int64) (int64, error) {
	return _this.LinhasDeletadas, _this.Err
}
