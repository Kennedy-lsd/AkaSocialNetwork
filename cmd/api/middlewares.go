package api

import (
	"log"
	"net/http"
	"time"
)

// LOGGER
const (
	ColorReset = "\033[0m"
	ColorCyan  = "\033[36m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
)

type responseWriter struct {
	writer     http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.writer.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	return rw.writer.Write(data)
}

func (rw *responseWriter) Header() http.Header {
	return rw.writer.Header()
}

func statusColor(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return ColorGreen
	case statusCode >= 400 && statusCode <= 500:
		return ColorRed
	default:
		return ColorReset
	}
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			writer:     w,
			statusCode: http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		color := statusColor(wrapped.statusCode)

		log.Printf("%s%s %s%s %s%d %s%s\n", ColorCyan, r.Method, ColorReset, r.URL.Path, color, wrapped.statusCode, ColorReset, time.Since(start))

	})
}
