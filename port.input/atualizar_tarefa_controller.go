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

func AtualizarTabelasController(w http.ResponseWriter, r *http.Request) {

	//Contexto
	uC := loghelper.CreateUserContext(r)

	loghelper.LogInfo(uC, loghelper.I00001)

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		loghelper.LogError(uC, loghelper.E00004, err)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Pega Body
	var tRequest request.TarefaRequest
	err = json.NewDecoder(r.Body).Decode(&tRequest)
	if err != nil {
		loghelper.LogError(uC, loghelper.E00004, err)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Chama UseCase
	linhasAtualizadas, err := usecase.AtualizarTarefaUsecase(uC, int64(codigo), adapter.TarefaModelFromTarefaRequest(tRequest))
	tResp := adapter.TarefaResponseFromTarefaRequest(tRequest)
	tResp.Codigo = int64(codigo)

	//Monta Resposta
	if err != nil {
		loghelper.LogError(uC, loghelper.E00005, err)
		httphelper.CreateResponse(w, http.StatusInternalServerError, "Erro ao realizar operação de atualizar.", []any{tResp})
		return
	}

	if linhasAtualizadas == 0 {
		loghelper.LogWarn(uC, loghelper.W00001)
		httphelper.CreateResponse(w, http.StatusUnprocessableEntity, "Registro não encontrado", []any{tResp})
		return
	}

	//Sucesso
	loghelper.LogInfo(uC, loghelper.I00002)
	httphelper.CreateResponse(w, http.StatusOK, "UPDATED", []any{tResp})
	return
}
