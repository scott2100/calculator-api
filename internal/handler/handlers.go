package handler

import (
	"calculator-api/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func HandleSum(writer http.ResponseWriter, request *http.Request) {
	var sumNumbers model.SumNumbers

	err := json.NewDecoder(request.Body).Decode(&sumNumbers)
	if err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := 0

	for i := range sumNumbers.Values {
		result += sumNumbers.Values[i]
	}

	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "Received: Numbers=%d\n", sumNumbers.Values)
	fmt.Fprintf(writer, "Result: %d\n", result)
}

func HandleDivide(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers model.Numbers) (error, int) {
		if numbers.Number2 == 0 {
			return errors.New("Cannot divide by zero"), 0
		}
		return nil, numbers.Number1 / numbers.Number2
	})
}

func HandleMultiply(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers model.Numbers) (error, int) {
		return nil, numbers.Number1 * numbers.Number2
	})
}

func HandleSubtract(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers model.Numbers) (error, int) {
		return nil, numbers.Number1 - numbers.Number2
	})
}

func HandleAdd(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers model.Numbers) (error, int) {
		return nil, numbers.Number1 + numbers.Number2
	})
}

func HandleRoot(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func handleOperation(writer http.ResponseWriter, request *http.Request, operation model.OperationFunc) {
	var numbers model.Numbers

	err := decodeRequest(request, &numbers)
	if err != nil {
		logger.Error("Error decoding request body",
			slog.String("statusCode", strconv.Itoa(http.StatusBadRequest)),
			slog.String("error", err.Error()))
	}

	err, result := operation(numbers)
	if err != nil {
		http.Error(writer, "Error: "+err.Error(), http.StatusInternalServerError)
	}
	response := model.ResultResponse{Result: result}

	err = encodeResponse(writer, response)
	if err != nil {
		http.Error(writer, "Error encoding JSON"+err.Error(), http.StatusInternalServerError)
		return
	}
}
