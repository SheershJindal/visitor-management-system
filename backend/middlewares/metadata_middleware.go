package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// MetadataMiddleware adds standard metadata like request ID and execution time
func MetadataMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a unique request ID for tracing
		requestID := uuid.New().String()

		// Capture the start time for execution time
		startTime := time.Now()

		// Add metadata to context
		ctx := context.WithValue(r.Context(), "requestID", requestID)
		ctx = context.WithValue(ctx, "startTime", startTime)

		// Pass the context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
