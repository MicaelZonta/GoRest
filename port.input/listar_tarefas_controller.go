package port_in

import (
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"GoRest/port.input/adapter"
	"net/http"
)

func ListarTarefasController(w http.ResponseWriter, r *http.Request) {

	//Contexto
	var uC = loghelper.UserContextImpl{}.Init(r.Context().Value("X-Correlation-Id").(string))
	uC.LogInfo(loghelper.I00009)
	uC.SendValue("1", "1")

	//HttpWriter
	writer := httphelper.HttpWriterImpl{}.Init()

	//Chama UseCase
	listarUsecase := usecase.ListarTarefaCaseImpl{}.Init(uC)
	tarefaModelSlice, err := listarUsecase.ListarTarefasExecute()

	if err != nil {
		uC.LogError(loghelper.E00004, err)
		writer.CreateResponse(w, http.StatusInternalServerError, "Erro durante listagem", nil)
		return
	}
	uC.LogInfo(loghelper.I00002)

	writer.CreateResponse(w, http.StatusOK, "Listagem de Tarefas Realizada", []any{adapter.TarefaResponseSliceFromTarefaModelSlice(tarefaModelSlice)})
}
