package main

import (
	"calculator-api/handlers"
	"calculator-api/middleware"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HandleRoot)
	router.HandleFunc("POST /add", handlers.HandleAdd)
	router.HandleFunc("POST /subtract", handlers.HandleSubtract)
	router.HandleFunc("POST /multiply", handlers.HandleMultiply)
	router.HandleFunc("POST /divide", handlers.HandleDivide)
	//router.HandleFunc("POST /sum", handlers.HandleSum)

	server := http.Server{
		Addr:    ":8983",
		Handler: cors.AllowAll().Handler(middleware.Logging(logger, router)),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}

}
