package utils

import "github.com/Raman5837/kafka.go/base/types"

// Utility Function To Send API Response
func HttpResponse(message string, exception *string, success bool) types.HttpResponse {

	return types.HttpResponse{Data: struct {
		Error   *string `json:"error"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
	}{Message: message, Error: exception, Success: success}}
}
