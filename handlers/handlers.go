package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type operationFunc func(numbers numbers) (error, int)

type numbers struct {
	Number1 int `json:"a"`
	Number2 int `json:"b"`
}

type calcResult struct {
	Result int `json:"result"`
}

func HandleDivide(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers numbers) (error, int) {
		if numbers.Number2 == 0 {
			return errors.New("Cannot divide by zero"), 0
		}
		return nil, numbers.Number1 / numbers.Number2
	})
}

func HandleMultiply(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers numbers) (error, int) {
		return nil, numbers.Number1 * numbers.Number2
	})
}

func HandleSubtract(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers numbers) (error, int) {
		return nil, numbers.Number1 - numbers.Number2
	})
}

func HandleAdd(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(numbers numbers) (error, int) {
		return nil, numbers.Number1 + numbers.Number2
	})
}

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err := fmt.Fprint(writer, "Welcome to my Calculator API!")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleOperation(writer http.ResponseWriter, request *http.Request, operation operationFunc) {

	var numbers numbers

	err := json.NewDecoder(request.Body).Decode(&numbers)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err, result := operation(numbers)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := calcResult{Result: result}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

/*
type sumNumbers struct {
	Values []int `json:"values"`
}

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

*/
