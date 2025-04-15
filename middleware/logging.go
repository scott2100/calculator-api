package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Info("Incoming Request",
			slog.String("ip_address", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("path", r.RequestURI),
			slog.String("url", r.URL.String()),
		)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		logger.Info("Request Completed",
			slog.String("method", r.Method), // Ensure method is logged on completion
			slog.String("path", r.RequestURI),
			slog.Duration("duration", duration), // Use slog.Duration for time duration
		)

	})
}
