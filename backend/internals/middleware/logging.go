package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

var logPath = "/app/logs/access.log"

type contextKey string

const (
	contextKeyRequestID contextKey = "requestID"
)

// LoggingMiddleware logs incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Generate a unique request ID
		requestID := uuid.New().String()

		// Log to file
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Error opening log file: %v", err)
		} else {
			defer logFile.Close()
			logger := log.New(logFile, "", 0)
			logger.Printf("[%s] [%s] %s %s %s", requestID, r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		}

		// Set requestID in request context for traceability
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextKeyRequestID, requestID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
