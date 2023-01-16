package handlers

import (
	"GoRest/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	//Pega Body
	var t models.Tarefa
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("Erro durante decode de Json, %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Usecase
	codigo, err := models.Insert(t)

	//Monta Resp
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   "true",
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   "false",
			"Message": fmt.Sprintf("Sucesso ao inserir: %v", codigo),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
