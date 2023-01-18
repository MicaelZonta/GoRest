package port_in

import (
	"GoRest/core/model"
	"GoRest/core/usecase"
	"GoRest/crosscutting/httphelper"
	"GoRest/crosscutting/loghelper"
	"encoding/json"
	"net/http"
)

func CriarTarefaPortIn(w http.ResponseWriter, r *http.Request) {

	//Contexto
	uC := loghelper.CreateUserContext(r)

	//Pega Body
	var t model.Tarefa
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		loghelper.LogError(uC, loghelper.E00004, err)
		httphelper.CreateResponse(w, http.StatusBadRequest, "Entrada inválida", nil)
		return
	}

	//Usecase
	t, err = usecase.CriarTarefaUsecase(uC, t)

	if err != nil {
		loghelper.LogError(uC, loghelper.E00005, err)
		httphelper.CreateResponse(w, http.StatusInternalServerError, "Erro durante processo de salvar Tarefa", []any{t})
		return
	}

	//Monta Resp

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
