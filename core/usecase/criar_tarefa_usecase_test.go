package usecase

import (
	"GoRest/core/gateway/mock"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"testing"
)

func TestCriarTarefaExecute(t *testing.T) {
	criarTarefaUseCase := CriarTarefaUseCaseImpl{CriarTarefaGateway: mock.CriarTarefaGatewayMock{}}
	criarTarefaUseCase.uC = loghelper.UserContextImpl{}.Init("Teste")

	tarefaModel := model.TarefaModel{Titulo: "Test Tarefa Title", Descricao: "Test Tarefa Description", Completa: false}

	expectedTarefaModel := tarefaModel
	actualTarefaModel, err := criarTarefaUseCase.CriarTarefaExecute(tarefaModel)

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	if expectedTarefaModel != actualTarefaModel {
		t.Errorf("Expected tarefa model to be %v, but got %v", expectedTarefaModel, actualTarefaModel)
	}
}
