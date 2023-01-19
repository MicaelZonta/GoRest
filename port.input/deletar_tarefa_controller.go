package port_in

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/port.input/transferobject/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func DeletarTabelaController(w http.ResponseWriter, r *http.Request) {

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

	//Chama UseCase
	linhasDeletadas, err := usecase.DeletarTarefaUsecase(uC, int64(codigo))

	if err != nil {
		loghelper.LogError(uC, loghelper.E00005, err)
		httphelper.CreateResponse(w, http.StatusInternalServerError, "Erro ao realizar operação de deletar.", []any{response.TarefaResponse{Codigo: int64(codigo)}})
		return
	}

	if linhasDeletadas == 0 {
		loghelper.LogWarn(uC, loghelper.W00001)
		httphelper.CreateResponse(w, http.StatusUnprocessableEntity, "Registro não encontrado", []any{response.TarefaResponse{Codigo: int64(codigo)}})
		return
	}

	loghelper.LogInfo(uC, loghelper.I00002)
	httphelper.CreateResponse(w, http.StatusOK, "Registro deletado", []any{response.TarefaResponse{Codigo: int64(codigo)}})
}
