package mock

import (
	"GoRest/core/model"
	"GoRest/core/usecase"
	"GoRest/crosscutting/loghelper"
)

type CriarTarefaUsecaseMock struct {
	Err error
}

func (_this CriarTarefaUsecaseMock) Init(uC loghelper.UserContext) usecase.CriarTarefaUseCase {
	return _this
}

func (_this CriarTarefaUsecaseMock) CriarTarefaExecute(tM model.TarefaModel) (model.TarefaModel, error) {
	return tM, _this.Err
}
