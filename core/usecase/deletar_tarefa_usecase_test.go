package usecase

import (
	"GoRest/core/gateway/mock"
	"GoRest/crosscutting/loghelper"
	"testing"
)

func TestDeletarTarefaExecute(t *testing.T) {
	deletarTarefaUseCase := DeletarTarefaUsecaseImpl{
		DeletarTarefaGateway: mock.DeletarTarefaGatewayMock{
			LinhasDeletadas: 1,
			Err:             nil,
		},
	}
	deletarTarefaUseCase.uC = loghelper.UserContextImpl{}.Init("Teste")

	linhasDeletadas, err := deletarTarefaUseCase.DeletarTarefaExecute(1)

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	if 1 != linhasDeletadas {
		t.Errorf("Expected linhas deletadas to be %v, but got %v", 1, linhasDeletadas)
	}
}
