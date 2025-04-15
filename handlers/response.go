package handlers

import (
	"encoding/json"
	"net/http"
)

func encodeResponse(writer http.ResponseWriter, result resultResponse) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(result)
	if err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return err
	}
	return nil
}
