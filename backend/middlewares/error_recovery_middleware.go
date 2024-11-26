package middlewares

import (
	"log"
	"net/http"

	"github.com/sheershjindal/visitor-management-system/models"
	"github.com/sheershjindal/visitor-management-system/utils"
)

func ErrorRecoveryMiddleware(next func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v", rec)
				utils.SendResponse(w, r, "error", "Internal Server Error", nil, nil, http.StatusInternalServerError)
			}
		}()

		if err := next(w, r); err != nil {
			if httpErr, ok := err.(*models.HTTPError); ok {
				utils.SendResponseWithErrorCode(w, r, "error", httpErr.Message, httpErr.Errors, nil, httpErr.ErrorCode, httpErr.StatusCode)
			} else {
				utils.SendResponse(w, r, "error", "Internal Server Error", nil, nil, http.StatusInternalServerError)
			}
		}
	}
}
