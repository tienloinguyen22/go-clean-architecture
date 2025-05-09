package httputils

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ResponseWithError(w http.ResponseWriter, statusCode int, err error) {
	ResonseWithJSON(w, statusCode, HttpResponse{
		Status:  "error",
		Message: err.Error(),
	})
}

func ResonseWithJSON(w http.ResponseWriter, statusCode int, data any) {
	var response interface{}
	if _, ok := data.(HttpResponse); !ok {
		response = HttpResponse{
			Status: "success",
			Data:   data,
		}
	} else {
		response = data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
