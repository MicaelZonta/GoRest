package handlers

import (
	"GoRest/models"
	"encoding/json"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	//Chama UseCase
	tarefaSlice, err := models.GetAll()
	if err != nil {
		//log.Println("Erro buscar lista")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefaSlice)
}
