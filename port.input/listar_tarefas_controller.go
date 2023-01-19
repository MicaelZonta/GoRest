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
	uC := loghelper.CreateUserContext(r)
	loghelper.LogInfo(uC, loghelper.I00002)

	//Chama UseCase
	tarefaModelSlice, err := usecase.ListarTarefaUsecase(uC)

	if err != nil {
		loghelper.LogError(uC, loghelper.E00004, err)
		httphelper.CreateResponse(w, http.StatusInternalServerError, "Erro durante listagem", nil)
		return
	}
	loghelper.LogInfo(uC, loghelper.I00002)

	httphelper.CreateResponse(w, http.StatusOK, "Listagem de Tarefas Realizada", []any{adapter.TarefaResponseSliceFromTarefaModelSlice(tarefaModelSlice)})
}
