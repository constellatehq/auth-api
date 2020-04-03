package model

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string `json:"message"`               // Message to be display to the end user without debugging information
	Type       string `json:"type"`                  // Error type
	StatusCode int    `json:"status_code,omitempty"` // The http status code
	params     map[string]interface{}
}

type ErrorResponse struct {
	Error Error `json:"error"` // Error Object
}

func NewErrorResponse(errorType string, message string, params map[string]interface{}) *ErrorResponse {
	errorResponse := &ErrorResponse{}
	err := &Error{}
	err.Message = message
	err.Type = errorType
	err.Message = message
	errorResponse.Error = *err
	return errorResponse
}

func CreateErrorResponse(w http.ResponseWriter, statusCode int, errorType string, message string, params map[string]interface{}) {
	err := NewErrorResponse(errorType, message, params)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
