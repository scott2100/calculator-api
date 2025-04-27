package main

import (
	"calculator-api/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type RateLimiter struct {
	requestCount  int
	mu            sync.Mutex
	resetInterval time.Duration
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		resetInterval: interval,
	}
}

func (rl *RateLimiter) Reset() {
	rl.mu.Lock()
	rl.requestCount = 0
	rl.mu.Unlock()
}

func (rl *RateLimiter) RequestCount() int {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.requestCount
}

func (rl *RateLimiter) StartResetTimer(done <-chan bool) {
	ticker := time.NewTicker(rl.resetInterval)
	fmt.Println("Starting reset timer with interval ", rl.resetInterval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at ", t)
			rl.Reset()
			fmt.Println("Rate set to zero.")
		case <-done:
			return
		}
	}
}

func (rl *RateLimiter) IncrementCount() {
	rl.mu.Lock()
	rl.requestCount++
	rl.mu.Unlock()
}

func main() {

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	done := make(chan bool)
	defer close(done)

	rl := NewRateLimiter(time.Second * 5)
	go rl.StartResetTimer(done)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleRoot(w, r)
	})
	router.Post("/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle add")
		rl.IncrementCount()
		fmt.Println("Rate value: ", rl.requestCount)
		if rl.RequestCount() < 10 {
			handlers.HandleAdd(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
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
