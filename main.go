package main

import (
	"calculator-api/middleware"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

type Numbers struct {
	Number1 int `json:"a"`
	Number2 int `json:"b"`
}

type SumNumbers struct {
	Values []int `json:"values"`
}

type OperationFunc func(int, int) int

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handleRoot)
	router.HandleFunc("POST /add", handleAdd)
	router.HandleFunc("POST /subtract", handleSubtract)
	router.HandleFunc("POST /multiply", handleMultiply)
	router.HandleFunc("POST /divide", handleDivide)
	router.HandleFunc("POST /sum", handleSum)

	server := http.Server{
		Addr:    ":8983",
		Handler: middleware.Logging(logger, router),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}

}

func handleSum(writer http.ResponseWriter, request *http.Request) {
	var sumNumbers SumNumbers

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

func handleDivide(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(a, b int) int {
		return a / b
	})
}

func handleMultiply(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(a, b int) int {
		return a * b
	})
}

func handleSubtract(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(a, b int) int {
		return a - b
	})
}

func handleAdd(writer http.ResponseWriter, request *http.Request) {
	handleOperation(writer, request, func(a, b int) int {
		return a + b
	})
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleOperation(writer http.ResponseWriter, request *http.Request, operation OperationFunc) {
	var numbers Numbers

	// Decode the JSON body into the Numbers struct
	err := json.NewDecoder(request.Body).Decode(&numbers)
	if err != nil {
		logger.Error("Error decoding request body",
			slog.String("statusCode", strconv.Itoa(http.StatusBadRequest)),
			slog.String("error", err.Error()))
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Perform the operation
	result := operation(numbers.Number1, numbers.Number2)

	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "Received: Number1=%d, Number2=%d\n", numbers.Number1, numbers.Number2)
	fmt.Fprintf(writer, "Result: %d\n", result)
}
