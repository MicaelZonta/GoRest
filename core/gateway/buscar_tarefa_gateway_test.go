package gateway

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/entity"
	"GoRest/infra/repository/postgres/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuscarTarefaExecute(t *testing.T) {
	codigoBusca := int64(1)

	tarefaRepositoryMock := mock.TarefaRepositoryMock{}

	tarefaRepositoryMock.BuscarTarefaEntity = entity.TarefaEntity{
		Codigo:    codigoBusca,
		Titulo:    "Tarefa Teste",
		Descricao: "Descrição da tarefa",
		Completa:  false,
	}

	buscarGateway := BuscarTarefaGatewayImpl{
		uC:               loghelper.UserContextImpl{}.Init("Teste"),
		TarefaRepository: tarefaRepositoryMock,
	}

	tarefa, err := buscarGateway.BuscarTarefaExecute(codigoBusca)

	assert.Nil(t, err)
	assert.Equal(t, codigoBusca, tarefa.Codigo)
	assert.Equal(t, tarefaRepositoryMock.BuscarTarefaEntity.Titulo, tarefa.Titulo)
	assert.Equal(t, tarefaRepositoryMock.BuscarTarefaEntity.Descricao, tarefa.Descricao)
	assert.Equal(t, tarefaRepositoryMock.BuscarTarefaEntity.Completa, tarefa.Completa)

}
