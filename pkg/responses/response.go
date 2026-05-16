package responses

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func JSON(w http.ResponseWriter, status int, message string, data interface{}) {
	respond(w, status, message, data, nil, nil)
}

func Error(w http.ResponseWriter, status int, message string) {
	respond(w, status, message, nil, nil, nil)
}

func ValidationError(w http.ResponseWriter, message string, errors interface{}) {
	respond(w, http.StatusUnprocessableEntity, message, nil, errors, nil)
}

func PaginatedJSON(w http.ResponseWriter, status int, message string, data interface{}, meta interface{}) {
	respond(w, status, message, data, nil, meta)
}

func respond(w http.ResponseWriter, status int, message string, data interface{}, errors interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
		Meta:    meta,
	}
	json.NewEncoder(w).Encode(response)
}
