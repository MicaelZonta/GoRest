package handlers

import (
	"GoRest/core/model"
	"GoRest/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		//log.Println("Erro buscar param")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Pega Body
	var t model.Tarefa
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		//log.Println("Erro durante decode de Json")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Chama UseCase
	updateRows, err := models.Update(int64(codigo), t)

	//Monta Resposta
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   "true",
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   "false",
			"Message": fmt.Sprintf("Quantidade de linhas atualizzadas ao inserir: %v", updateRows),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
