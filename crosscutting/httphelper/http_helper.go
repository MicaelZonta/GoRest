package httphelper

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  int
	Message string
	Data    []any
}

func CreateResponse(w http.ResponseWriter, httpStatus int, msg string, obj []any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpStatus)

	r := response{
		Status:  httpStatus,
		Message: msg,
		Data:    obj,
	}

	jsonResp, err := json.Marshal(r)
	if err != nil {
		w.Write([]byte{})
	}
	w.Write(jsonResp)
}
