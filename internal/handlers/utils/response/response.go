package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_Undefined/internal/handlers/dto"
)

func SendError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp, err := json.Marshal(dto.ErrorDTO{Message: message})
	if err != nil {

		log.Printf("failed to marshal response: %s", err.Error())
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Printf("failed to write response: %s", err.Error())
	}
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, body any) {
	if body == nil {
		w.WriteHeader(statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to marshal response %s", err.Error())
		return
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(resp); err != nil {
		log.Printf("failed to write response %s", err.Error())
	}
}
