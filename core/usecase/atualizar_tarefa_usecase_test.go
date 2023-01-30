package usecase

import (
	"GoRest/core/gateway/mock"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"testing"
)

func TestAtualizarTarefaExecute(t *testing.T) {
	atualizarTarefaUseCase := AtualizarTarefaUsecaseImpl{
		AtualizarTarefaGateway: mock.AtualizarTarefaGatewayMock{
			LinhasAtualizadas: 1,
			Err:               nil,
		},
	}
	atualizarTarefaUseCase.uC = loghelper.UserContextImpl{}.Init("Teste")

	tarefaModel := model.TarefaModel{
		Codigo:    int64(1),
		Titulo:    "Tarefa Teste",
		Descricao: "Descrição da tarefa",
		Completa:  false,
	}

	linhasAtualizadas, err := atualizarTarefaUseCase.AtualizarTarefaExecute(1, tarefaModel)

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	if 1 != linhasAtualizadas {
		t.Errorf("Expected linhas atualizadas to be %v, but got %v", 1, linhasAtualizadas)
	}
}
