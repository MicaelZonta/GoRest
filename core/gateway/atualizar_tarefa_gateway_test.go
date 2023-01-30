package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtualizarTarefaExecute(t *testing.T) {
	tarefaRepositoryMock := mock.TarefaRepositoryMock{}
	tarefaRepositoryMock.LinhasRetornadasAtualizar = 1

	tarefaModel := model.TarefaModel{
		Codigo:    int64(1),
		Titulo:    "Tarefa Teste",
		Descricao: "Descrição da tarefa",
		Completa:  false,
	}

	atualizarGateway := AtualizarTarefaGatewayImpl{
		uC:               loghelper.UserContextImpl{}.Init("Teste"),
		TarefaRepository: tarefaRepositoryMock,
	}

	linhasAtualizadas, err := atualizarGateway.AtualizarTarefaExecute(1, tarefaModel)

	assert.Nil(t, err)
	assert.Equal(t, tarefaRepositoryMock.LinhasRetornadasAtualizar, linhasAtualizadas)
}
