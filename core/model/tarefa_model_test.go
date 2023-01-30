package model

import "testing"

func TestTarefaModel(t *testing.T) {
	// Arrange
	expectedCode := int64(1)
	expectedTitle := "Test Title"
	expectedDescription := "Test Description"
	expectedComplete := true

	// Act
	tarefaModel := TarefaModel{
		Codigo:    expectedCode,
		Titulo:    expectedTitle,
		Descricao: expectedDescription,
		Completa:  expectedComplete,
	}

	// Assert
	if tarefaModel.Codigo != expectedCode {
		t.Errorf("Expected codigo to be %d but got %d", expectedCode, tarefaModel.Codigo)
	}

	if tarefaModel.Titulo != expectedTitle {
		t.Errorf("Expected titulo to be %s but got %s", expectedTitle, tarefaModel.Titulo)
	}

	if tarefaModel.Descricao != expectedDescription {
		t.Errorf("Expected descricao to be %s but got %s", expectedDescription, tarefaModel.Descricao)
	}

	if tarefaModel.Completa != expectedComplete {
		t.Errorf("Expected completa to be %t but got %t", expectedComplete, tarefaModel.Completa)
	}
}
