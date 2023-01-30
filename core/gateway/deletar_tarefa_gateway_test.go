package gateway

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeletarTarefaExecute(t *testing.T) {
	codigoBusca := int64(1)

	tarefaRepositoryMock := mock.TarefaRepositoryMock{}
	tarefaRepositoryMock.LinhasRetornadasDeletar *= 1

	deletarGateway := DeletarTarefaGatewayImpl{
		uC:               loghelper.UserContextImpl{}.Init("Teste"),
		TarefaRepository: tarefaRepositoryMock,
	}

	linhasDeletadas, err := deletarGateway.DeletarTarefaExecute(codigoBusca)

	assert.Nil(t, err)
	assert.Equal(t, tarefaRepositoryMock.LinhasRetornadasAtualizar, linhasDeletadas)

}
