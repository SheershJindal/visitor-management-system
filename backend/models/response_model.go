package models

// PaginationMeta is used for paginated responses.
// type PaginationMeta struct {
// 	Total       int `json:"total"`
// 	PerPage     int `json:"per_page"`
// 	CurrentPage int `json:"current_page"`
// 	TotalPages  int `json:"total_pages"`
// }

// Response represents a response structure.
type Response struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Meta      interface{} `json:"meta"`
	Errors    interface{} `json:"errors,omitempty"`
}
