package middlewares

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sheershjindal/visitor-management-system/config"
)

// LoggingMiddleware logs HTTP requests and their execution time
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime, ok := r.Context().Value("startTime").(time.Time)
		if !ok {
			startTime = time.Now()
		}
		requestID := r.Context().Value("requestID")

		// Log incoming request details
		log.Printf("Request [%s]: %s %s %s", requestID, r.Method, r.URL.Path, r.RemoteAddr)

		// Log body or URL data if present and feature flag is enabled
		if config.IsDetailedLoggingEnabled() {
			logRequestData(r)
		}

		// Wrap the ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		// Log response details
		duration := time.Since(startTime)
		log.Printf("Response [%s]: %d %s (took %s)", requestID, lrw.statusCode, http.StatusText(lrw.statusCode), duration)
	})
}

// loggingResponseWriter wraps http.ResponseWriter to capture status codes
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// logRequestData logs the request data in a formatted manner based on its type.
func logRequestData(r *http.Request) {
	requestID, _ := r.Context().Value("requestID").(string)

	// Log query parameters
	if len(r.URL.Query()) > 0 {
		log.Printf("Request [%s]: Query Parameters:\n%s", requestID, "  "+formatQueryParameters(r.URL.Query()))
	}

	// Log body
	if r.Body != nil && r.ContentLength > 0 {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			formattedBody := formatBody(body, r.Header.Get("Content-Type"))
			log.Printf("Request [%s]: Body:\n%s", requestID, "  "+formattedBody)

			// Restore the body so it can be read downstream
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		} else {
			log.Printf("Request [%s]: Failed to read body: %v", requestID, err)
		}
	}
}

func formatQueryParameters(params url.Values) string {
	var formatted strings.Builder
	for key, values := range params {
		formatted.WriteString(key + ": " + strings.Join(values, ", ") + "\n")
	}
	return formatted.String()
}

// formatBody formats the body based on the content type.
func formatBody(body []byte, contentType string) string {
	contentType = strings.ToLower(contentType)

	switch {
	case strings.Contains(contentType, "application/json"):
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, body, "  ", "  "); err == nil {
			return prettyJSON.String()
		}
		return string(body)

	case strings.Contains(contentType, "application/xml"), strings.Contains(contentType, "text/xml"):
		var prettyXML bytes.Buffer
		if err := formatXML(body, &prettyXML); err == nil {
			return prettyXML.String()
		}
		return "[Invalid XML Format]"

	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		values, err := url.ParseQuery(string(body))
		if err == nil {
			return formatQueryParameters(values)
		}
		return string(body)

	case strings.Contains(contentType, "multipart/form-data"):
		return "[Multipart formdata body skipped]"
		// return formatMultipartFormData(r)

	case strings.Contains(contentType, "text/plain"):
		return string(body)

	default:
		return string(body)
	}
}

func formatMultipartFormData(r *http.Request) string {
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		return "[Failed to parse multipart form: " + err.Error() + "]"
	}
	var formatted strings.Builder

	formatted.WriteString("[Multipart form data]\n")

	for key, values := range r.MultipartForm.Value {
		formatted.WriteString("  " + key + ": " + strings.Join(values, ", ") + "\n")
	}
	for key := range r.MultipartForm.File {
		formatted.WriteString("  " + key + ": [File Present]\n")
	}
	return formatted.String()
}

// formatXML indents and formats an XML body.
func formatXML(input []byte, output *bytes.Buffer) error {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	encoder := xml.NewEncoder(output)
	encoder.Indent("  ", "  ")

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Skip unnecessary whitespace tokens
		if _, ok := token.(xml.CharData); ok {
			// Trim leading/trailing whitespace from text nodes
			token = xml.CharData(bytes.TrimSpace([]byte(token.(xml.CharData))))
		}

		if err := encoder.EncodeToken(token); err != nil {
			return err
		}
	}

	return encoder.Flush()
}
