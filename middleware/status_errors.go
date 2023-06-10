package middleware

import (
	"encoding/json"
	"net/http"

	"travelis/m/v2/model"
)

func StatusOK(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(response)
}

func StatusInternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(model.Error{
		Message: "Internal Server Error",
	})
}

func StatusBadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(model.Error{
		Message: message,
	})
}

func StatusUnprocessableEntity(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(model.Error{
		Message: message,
	})
}
