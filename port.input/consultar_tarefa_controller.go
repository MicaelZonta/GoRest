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
	var uC = loghelper.UserContextImpl{}.Init(r.Context().Value("X-Correlation-Id").(string))
	uC.LogInfo(loghelper.I00009)

	//HttpWriter
	writer := httphelper.HttpWriterImpl{}.Init()

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		uC.LogError(loghelper.E00001, err)
		writer.CreateResponse(w, http.StatusBadRequest, "Paramêtro inválido", nil)
		return
	}

	//Chama UseCase
	consultarUsecase := usecase.BuscarTarefaUseCaseImpl{}.Init(uC)
	tModel, err := consultarUsecase.BuscarTarefaExecute(int64(codigo))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			uC.LogWarn(loghelper.W00001)
			writer.CreateResponse(w, http.StatusUnprocessableEntity, "Código não encontrado", nil)
			return
		} else {
			uC.LogError(loghelper.E00002, err)
			writer.CreateResponse(w, http.StatusInternalServerError, "Banco de dados indisponível", nil)
			return
		}
	}

	//Sucesso
	uC.LogInfo(loghelper.I00002)
	writer.CreateResponse(w, http.StatusOK, "CREATED", []any{adapter.TarefaResponseFromTarefaModel(tModel)})
	return
}
