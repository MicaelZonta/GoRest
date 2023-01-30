package usecase

import (
	"GoRest/core/gateway/mock"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"testing"
)

func TestBuscarTarefaExecute(t *testing.T) {
	tarefaModel := model.TarefaModel{
		Codigo:    int64(1),
		Titulo:    "Tarefa Teste",
		Descricao: "Descrição da tarefa",
		Completa:  false,
	}

	buscarTarefaUseCase := BuscarTarefaUseCaseImpl{
		BuscarTarefaGateway: mock.BuscarTarefaGatewayMock{
			BuscarTarefaModel: tarefaModel,
			Err:               nil,
		},
	}
	buscarTarefaUseCase.uC = loghelper.UserContextImpl{}.Init("Teste")

	tarefaReturn, err := buscarTarefaUseCase.BuscarTarefaExecute(1)

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	if tarefaModel.Codigo != tarefaReturn.Codigo {
		t.Errorf("Expected codigo tarefa to be %v, but got %v", tarefaModel.Codigo, tarefaReturn.Codigo)
	}
}
