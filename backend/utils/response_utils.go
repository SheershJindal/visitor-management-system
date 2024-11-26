package utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/sheershjindal/visitor-management-system/models"
)

func SendResponse(w http.ResponseWriter, r *http.Request, status, message string, dataOrErrors interface{}, meta interface{}, statusCode int) error {
	return SendResponseWithErrorCode(w, r, status, message, dataOrErrors, meta, "", statusCode)
}

// SendResponseWithError sends both success and error responses
func SendResponseWithErrorCode(w http.ResponseWriter, r *http.Request, status, message string, dataOrErrors interface{}, meta interface{}, errorCode string, statusCode int) error {
	// Always include request ID and execution time
	requestID := r.Context().Value("requestID").(string)
	startTime := r.Context().Value("startTime").(time.Time)
	executionTime := time.Since(startTime).String()

	status = strings.ToUpper(status)

	// Prepare the response with meta containing the standard metadata
	response := models.Response{
		Status:  status,
		Message: message,
		Meta: map[string]interface{}{
			"requestID":     requestID,
			"executionTime": executionTime,
			// Include other static meta if needed (e.g., API version)
		},
	}

	// Add dynamic metadata (like pagination) if provided
	if meta != nil {
		if m, ok := meta.(map[string]interface{}); ok {
			for key, value := range m {
				// Make sure we modify the map inside Meta
				response.Meta.(map[string]interface{})[key] = value
			}
		}
	}

	// Set data or errors based on the response type (success or error)
	if status == "SUCCESS" {
		response.Data = dataOrErrors
	} else if status == "ERROR" {
		response.Errors = dataOrErrors
		response.ErrorCode = errorCode
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
