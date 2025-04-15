package handlers

import (
	"encoding/json"
	"net/http"
)

func decodeRequest(request *http.Request, numbers *numbers) error {
	err := json.NewDecoder(request.Body).Decode(&numbers)
	if err != nil {
		return err
	}
	return nil
}
