package mock

import (
	"GoRest/core/model"
	"GoRest/core/usecase"
	"GoRest/crosscutting/loghelper"
)

type ListarTarefaUsecaseMock struct {
	ListarTarefasModels []model.TarefaModel
	Err                 error
}

func (_this ListarTarefaUsecaseMock) Init(uC loghelper.UserContext) usecase.ListarTarefaUseCase {
	return _this
}

func (_this ListarTarefaUsecaseMock) ListarTarefasExecute() ([]model.TarefaModel, error) {
	return _this.ListarTarefasModels, _this.Err
}
