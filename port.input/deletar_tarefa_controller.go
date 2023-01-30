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

func DeletarTarefaController(w http.ResponseWriter, r *http.Request) {

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

	//Chama UseCase
	deletarUsecase := usecase.DeletarTarefaUsecaseImpl{}.Init(uC)
	linhasDeletadas, err := deletarUsecase.DeletarTarefaExecute(int64(codigo))

	if err != nil {
		uC.LogError(loghelper.E00005, err)
		writer.CreateResponse(w, http.StatusInternalServerError, "Erro ao realizar operação de deletar.", []any{response.TarefaResponse{Codigo: int64(codigo)}})
		return
	}

	if linhasDeletadas == 0 {
		uC.LogWarn(loghelper.W00001)
		writer.CreateResponse(w, http.StatusUnprocessableEntity, "Registro não encontrado", []any{response.TarefaResponse{Codigo: int64(codigo)}})
		return
	}

	uC.LogInfo(loghelper.I00002)
	writer.CreateResponse(w, http.StatusOK, "Registro deletado", []any{response.TarefaResponse{Codigo: int64(codigo)}})
}
