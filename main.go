package main

import (
	"calculator-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleRoot(w, r)
	})
	router.Post("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleAdd(w, r)
	})
	router.Post("/subtract", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSubtract(w, r)
	})
	router.Post("/multiply", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleMultiply(w, r)
	})
	router.Post("/divide", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleDivide(w, r)
	})

	err := http.ListenAndServe(":8983", router)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
