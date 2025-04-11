package main

import (
	"calculator-api/internal/handler"
	"calculator-api/internal/middleware"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handler.HandleRoot)
	router.HandleFunc("POST /add", handler.HandleAdd)
	router.HandleFunc("POST /subtract", handler.HandleSubtract)
	router.HandleFunc("POST /multiply", handler.HandleMultiply)
	router.HandleFunc("POST /divide", handler.HandleDivide)
	router.HandleFunc("POST /sum", handler.HandleSum)

	server := http.Server{
		Addr:    ":8983",
		Handler: cors.AllowAll().Handler(middleware.Logging(logger, router)),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}

}
