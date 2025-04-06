package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Initialize with default OK status
		}

		logger.Info("Incoming Request",
			slog.String("ip_address", r.RemoteAddr),
			slog.String("method", r.Method), // Corrected: Use r.Method for HTTP method
			slog.String("path", r.RequestURI),
			slog.String("url", r.URL.String()),
		)

		// Call the next handler with the wrapped writer
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		logger.Info("Request Completed",
			slog.Int("statusCode", wrapped.statusCode),
			slog.String("method", r.Method), // Ensure method is logged on completion
			slog.String("path", r.RequestURI),
			slog.Duration("duration", duration), // Use slog.Duration for time duration
		)

	})
}
