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
	uC := loghelper.CreateUserContext(r)
	loghelper.LogInfo(uC, loghelper.I00009)

	//Pega Body
	var tRequest request.TarefaRequest
	err := json.NewDecoder(r.Body).Decode(&tRequest)
	if err != nil {
		loghelper.LogError(uC, loghelper.E00004, err)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Usecase
	tModel, err := usecase.CriarTarefaUsecase(uC, adapter.TarefaModelFromTarefaRequest(tRequest))
	tResponse := adapter.TarefaResponseFromTarefaModel(tModel)

	if err != nil {
		loghelper.LogError(uC, loghelper.E00005, err)
		httphelper.CreateResponse(w, http.StatusInternalServerError, "Erro durante processo de salvar Tarefa", []any{tResponse})
		return
	}

	//Monta Resp
	loghelper.LogInfo(uC, loghelper.I00010)
	httphelper.CreateResponse(w, http.StatusCreated, "CREATED", []any{tResponse})
}
