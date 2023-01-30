package mock

import (
	"GoRest/core/model"
	"GoRest/core/usecase"
	"GoRest/crosscutting/loghelper"
)

type AtualizarTarefaUsecaseMock struct {
	LinhasAtualizadas int64
	Err               error
}

func (_this AtualizarTarefaUsecaseMock) Init(uC loghelper.UserContext) usecase.AtualizarTarefaUsecase {
	return _this
}

func (_this AtualizarTarefaUsecaseMock) AtualizarTarefaExecute(codigo int64, tarefaModel model.TarefaModel) (int64, error) {
	return _this.LinhasAtualizadas, _this.Err
}
