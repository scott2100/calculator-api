package handlers

import (
	"encoding/json"
	"net/http"
)

func encodeResponse(writer http.ResponseWriter, result calcResult) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(result)
	if err != nil {
		return err
	}
	return nil
}

func decodeRequest(request *http.Request, numbers *numbers) error {
	err := json.NewDecoder(request.Body).Decode(&numbers)
	if err != nil {
		return err
	}
	return nil
}
