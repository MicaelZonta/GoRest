package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/adapter"
	"GoRest/infra/repository/postgres/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListarTarefasExecute(t *testing.T) {
	tarefaRepositoryMock := mock.TarefaRepositoryMock{}

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

	tarefaRepositoryMock.ListarTarefaEntity = append(tarefaRepositoryMock.ListarTarefaEntity, adapter.TarefaEntityFromTarefaModel(tarefaModel1))
	tarefaRepositoryMock.ListarTarefaEntity = append(tarefaRepositoryMock.ListarTarefaEntity, adapter.TarefaEntityFromTarefaModel(tarefaModel2))
	tarefaRepositoryMock.ListarTarefaEntity = append(tarefaRepositoryMock.ListarTarefaEntity, adapter.TarefaEntityFromTarefaModel(tarefaModel3))

	listarGateway := ListarTarefaGatewayImpl{
		uC:               loghelper.UserContextImpl{}.Init("Teste"),
		TarefaRepository: tarefaRepositoryMock,
	}

	listaTarefaModel, err := listarGateway.ListarTarefasExecute()

	assert.Nil(t, err)

	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[0].Codigo, listaTarefaModel[0].Codigo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[0].Titulo, listaTarefaModel[0].Titulo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[0].Descricao, listaTarefaModel[0].Descricao)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[0].Completa, listaTarefaModel[0].Completa)

	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[1].Codigo, listaTarefaModel[1].Codigo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[1].Titulo, listaTarefaModel[1].Titulo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[1].Descricao, listaTarefaModel[1].Descricao)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[1].Completa, listaTarefaModel[1].Completa)

	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[2].Codigo, listaTarefaModel[2].Codigo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[2].Titulo, listaTarefaModel[2].Titulo)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[2].Descricao, listaTarefaModel[2].Descricao)
	assert.Equal(t, tarefaRepositoryMock.ListarTarefaEntity[2].Completa, listaTarefaModel[2].Completa)

}
