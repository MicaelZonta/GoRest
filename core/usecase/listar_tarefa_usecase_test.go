package usecase

import (
	"GoRest/core/gateway/mock"
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListarTarefasExecute(t *testing.T) {

	listaTarefas := montarListaTarefasModels()

	listarTarefaUseCase := ListarTarefaCaseImpl{ListarTarefaGateway: mock.ListarTarefaGatewayMock{
		ListarTarefasModels: listaTarefas,
		Err:                 nil,
	}}
	listarTarefaUseCase.uC = loghelper.UserContextImpl{}.Init("Teste")

	actualTarefaModel, err := listarTarefaUseCase.ListarTarefasExecute()

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	assert.Nil(t, err)

	assert.Equal(t, listaTarefas[0].Codigo, actualTarefaModel[0].Codigo)
	assert.Equal(t, listaTarefas[0].Titulo, actualTarefaModel[0].Titulo)
	assert.Equal(t, listaTarefas[0].Descricao, actualTarefaModel[0].Descricao)
	assert.Equal(t, listaTarefas[0].Completa, actualTarefaModel[0].Completa)

	assert.Equal(t, listaTarefas[1].Codigo, actualTarefaModel[1].Codigo)
	assert.Equal(t, listaTarefas[1].Titulo, actualTarefaModel[1].Titulo)
	assert.Equal(t, listaTarefas[1].Descricao, actualTarefaModel[1].Descricao)
	assert.Equal(t, listaTarefas[1].Completa, actualTarefaModel[1].Completa)

	assert.Equal(t, listaTarefas[2].Codigo, actualTarefaModel[2].Codigo)
	assert.Equal(t, listaTarefas[2].Titulo, actualTarefaModel[2].Titulo)
	assert.Equal(t, listaTarefas[2].Descricao, actualTarefaModel[2].Descricao)
	assert.Equal(t, listaTarefas[2].Completa, actualTarefaModel[2].Completa)
}

func montarListaTarefasModels() []model.TarefaModel {
	tarefaModel1 := model.TarefaModel{
		Codigo:    int64(1),
		Titulo:    "Tarefa Teste1",
		Descricao: "Descrição da tarefa1",
		Completa:  false,
	}
	tarefaModel2 := model.TarefaModel{
		Codigo:    int64(2),
		Titulo:    "Tarefa Teste2",
		Descricao: "Descrição da tarefa3",
		Completa:  true,
	}
	tarefaModel3 := model.TarefaModel{
		Codigo:    int64(3),
		Titulo:    "Tarefa Teste2",
		Descricao: "Descrição da tarefa3",
		Completa:  false,
	}

	listaTarefas := []model.TarefaModel{
		tarefaModel1, tarefaModel2, tarefaModel3,
	}

	return listaTarefas
}
