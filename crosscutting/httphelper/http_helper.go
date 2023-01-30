package httphelper

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Status  int
	Message string
	Data    []any
}

type HttpWriter interface {
	Init() HttpWriter
	CreateResponse(w http.ResponseWriter, httpStatus int, msg string, obj []any)
}

type HttpWriterImpl struct {
}

func (_this HttpWriterImpl) Init() HttpWriter {
	return _this
}

func (_this HttpWriterImpl) CreateResponse(w http.ResponseWriter, httpStatus int, msg string, obj []any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpStatus)

	r := httpResponse{
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
