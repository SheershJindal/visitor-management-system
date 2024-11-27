package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/sheershjindal/visitor-management-system/config"
	"github.com/sheershjindal/visitor-management-system/models"
)

// CORSMiddleware applies the global CORS configuration to HTTP responses
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsConfig := config.CORSConfig
		origin := r.Header.Get("Origin")

		// Validate origin
		if !isOriginAllowed(origin, corsConfig.AllowedOrigins) {
			http.Error(w, "CORS origin not allowed", http.StatusForbidden)
			return
		}

		// Apply CORS headers
		applyCORSHeaders(w, origin, corsConfig)

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}

// isOriginAllowed checks if the request origin is allowed by the CORS config
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	if len(allowedOrigins) == 0 || allowedOrigins[0] == "*" {
		return true // Default to allow all origins
	}
	for _, o := range allowedOrigins {
		if strings.EqualFold(o, origin) {
			return true
		}
	}
	return false
}

// applyCORSHeaders sets the appropriate CORS headers based on the config
func applyCORSHeaders(w http.ResponseWriter, origin string, corsConfig models.CORSConfig) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", joinOrDefault(corsConfig.AllowedMethods, "GET, POST, PUT, DELETE, OPTIONS"))
	w.Header().Set("Access-Control-Allow-Headers", joinOrDefault(corsConfig.AllowedHeaders, "Content-Type, Authorization"))
	w.Header().Set("Access-Control-Expose-Headers", joinOrDefault(corsConfig.ExposedHeaders, ""))
	if corsConfig.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	if corsConfig.MaxAge > 0 {
		w.Header().Set("Access-Control-Max-Age", strconv.Itoa(corsConfig.MaxAge))
	}
}

// joinOrDefault joins a slice with a comma or returns a default value if empty
func joinOrDefault(values []string, defaultValue string) string {
	if len(values) == 0 {
		return defaultValue
	}
	return strings.Join(values, ", ")
}
