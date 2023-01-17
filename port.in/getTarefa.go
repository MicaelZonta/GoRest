package port_in

import (
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {

	uC := loghelper.CreateUserContext(r)

	loghelper.LogInfo(uC, loghelper.I00001_CODE, loghelper.I00001_STRING)
	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		loghelper.LogError(uC, loghelper.E00001_CODE, loghelper.E00001_STRING)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Paramêtro inválido", nil)
		return
	}

	//Chama UseCase
	tarefa, err := models.Get(int64(codigo))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			loghelper.LogWarn(uC, loghelper.W00001_CODE, loghelper.W00001_STRING)
			httphelper.CreateResponse(w, http.StatusUnprocessableEntity, "Código não encontrado", nil)
			return
		} else {
			loghelper.LogError(uC, loghelper.E00002_CODE, loghelper.E00002_STRING)
			httphelper.CreateResponse(w, http.StatusInternalServerError, "Banco de dados indisponível", nil)
			return
		}
	}

	loghelper.LogInfo(uC, loghelper.I00002_CODE, loghelper.I00002_STRING)
	httphelper.CreateResponse(w, http.StatusOK, "CREATED", []any{tarefa})
	return
}
