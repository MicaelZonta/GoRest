package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/adapter"
	"GoRest/infra/repository/postgres/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCriarTarefaExecute(t *testing.T) {
	tarefaRepositoryMock := mock.TarefaRepositoryMock{}

	tarefaModelParameter := model.TarefaModel{
		Titulo:    "Tarefa Teste",
		Descricao: "Descrição da tarefa",
		Completa:  false,
	}

	tarefaRepositoryMock.CriarTarefaEntity = adapter.TarefaEntityFromTarefaModel(tarefaModelParameter)
	tarefaRepositoryMock.CriarTarefaEntity.Codigo *= 10

	criarGateway := CriarTarefaGatewayImpl{
		uC:               loghelper.UserContextImpl{}.Init("Teste"),
		TarefaRepository: tarefaRepositoryMock,
	}

	tarefaReturn, err := criarGateway.CriarTarefaExecute(tarefaModelParameter)

	assert.Nil(t, err)
	assert.Equal(t, tarefaRepositoryMock.BuscarTarefaEntity.Codigo, tarefaReturn.Codigo)
	assert.Equal(t, tarefaModelParameter.Titulo, tarefaReturn.Titulo)
	assert.Equal(t, tarefaModelParameter.Descricao, tarefaReturn.Descricao)
	assert.Equal(t, tarefaModelParameter.Completa, tarefaReturn.Completa)

}
