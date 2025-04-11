package handler

import (
	"calculator-api/internal/model"
	"encoding/json"
	"net/http"
)

func decodeRequest(request *http.Request, numbers *model.Numbers) error {
	err := json.NewDecoder(request.Body).Decode(&numbers)
	if err != nil {
		return err
	}
	return nil
}
