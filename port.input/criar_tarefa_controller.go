package port_in

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/port.input/adapter"
	"GoRest/port.input/transferobject/request"
	"encoding/json"
	"net/http"
)

func CriarTarefaController(w http.ResponseWriter, r *http.Request) {

	//Contexto
	var uC = loghelper.UserContextImpl{}.Init(r.Context().Value("X-Correlation-Id").(string))
	uC.LogInfo(loghelper.I00009)

	//HttpWriter
	writer := httphelper.HttpWriterImpl{}.Init()

	//Pega Body
	var tRequest request.TarefaRequest
	err := json.NewDecoder(r.Body).Decode(&tRequest)
	if err != nil {
		uC.LogError(loghelper.E00004, err)
		writer.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Usecase
	criarUsecase := usecase.CriarTarefaUseCaseImpl{}.Init(uC)
	tModel, err := criarUsecase.CriarTarefaExecute(adapter.TarefaModelFromTarefaRequest(tRequest))
	tResponse := adapter.TarefaResponseFromTarefaModel(tModel)

	if err != nil {
		uC.LogError(loghelper.E00005, err)
		writer.CreateResponse(w, http.StatusInternalServerError, "Erro durante processo de salvar Tarefa", []any{tResponse})
		return
	}

	//Monta Resp
	uC.LogInfo(loghelper.I00010)
	writer.CreateResponse(w, http.StatusCreated, "CREATED", []any{tResponse})
}
