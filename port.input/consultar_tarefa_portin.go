package port_in

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/port.input/adapter"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func ConsultarTarefaController(w http.ResponseWriter, r *http.Request) {

	//Contexto
	uC := loghelper.CreateUserContext(r)
	loghelper.LogInfo(uC, loghelper.I00001)

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		loghelper.LogError(uC, loghelper.E00001, err)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Paramêtro inválido", nil)
		return
	}

	//Chama UseCase
	tModel, err := usecase.BuscarTarefaUsecase(uC, int64(codigo))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			loghelper.LogWarn(uC, loghelper.W00001)
			httphelper.CreateResponse(w, http.StatusUnprocessableEntity, "Código não encontrado", nil)
			return
		} else {
			loghelper.LogError(uC, loghelper.E00002, err)
			httphelper.CreateResponse(w, http.StatusInternalServerError, "Banco de dados indisponível", nil)
			return
		}
	}

	//Sucesso
	loghelper.LogInfo(uC, loghelper.I00002)
	httphelper.CreateResponse(w, http.StatusOK, "CREATED", []any{adapter.TarefaResponseFromTarefaModel(tModel)})
	return
}
