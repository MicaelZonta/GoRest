package port_in

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/port.input/adapter"
	"GoRest/port.input/transferobject/request"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func AtualizarTarefaController(w http.ResponseWriter, r *http.Request) {

	//Contexto
	var uC = loghelper.UserContextImpl{}.Init(r.Context().Value("X-Correlation-Id").(string))
	uC.LogInfo(loghelper.I00009)

	//HttpWriter
	writer := httphelper.HttpWriterImpl{}.Init()

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		uC.LogError(loghelper.E00004, err)
		writer.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Pega Body
	var tRequest request.TarefaRequest
	err = json.NewDecoder(r.Body).Decode(&tRequest)
	if err != nil {
		uC.LogError(loghelper.E00004, err)
		writer.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Chama UseCase
	atualizarUsecase := usecase.AtualizarTarefaUsecaseImpl{}.Init(uC)
	linhasAtualizadas, err := atualizarUsecase.AtualizarTarefaExecute(int64(codigo), adapter.TarefaModelFromTarefaRequest(tRequest))
	tResp := adapter.TarefaResponseFromTarefaRequest(tRequest)
	tResp.Codigo = int64(codigo)

	//Monta Resposta
	if err != nil {
		uC.LogError(loghelper.E00005, err)
		writer.CreateResponse(w, http.StatusInternalServerError, "Erro ao realizar operação de atualizar.", []any{tResp})
		return
	}

	if linhasAtualizadas == 0 {
		uC.LogWarn(loghelper.W00001)
		writer.CreateResponse(w, http.StatusUnprocessableEntity, "Registro não encontrado", []any{tResp})
		return
	}

	//Sucesso
	uC.LogInfo(loghelper.I00002)
	writer.CreateResponse(w, http.StatusOK, "UPDATED", []any{tResp})
	return
}
