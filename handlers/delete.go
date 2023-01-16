package handlers

import (
	"GoRest/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	//Pega Param
	codigo, err := strconv.Atoi(chi.URLParam(r, "codigo"))
	if err != nil {
		log.Println("Erro buscar param")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Chama UseCase
	updateRows, err := models.Delete(int64(codigo))

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
			"Message": fmt.Sprintf("Quantidade de linhas deletadas: %v", updateRows),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
