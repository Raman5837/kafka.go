package utils

import "github.com/Raman5837/kafka.go/base/types"

// Utility Function To Send API Response
func HttpResponseFail(data *interface{}, message string, exception string) types.HttpResponse {

	return types.HttpResponse{Data: struct {
		Data    *interface{} `json:"data"`
		Error   *string      `json:"error"`
		Success bool         `json:"success"`
		Message string       `json:"message"`
	}{Data: data, Message: message, Error: &exception, Success: false}}
}

// Utility Function To Send Successful API Response
func HttpResponseOK(data interface{}, message string) types.HttpResponse {
	return types.HttpResponse{Data: struct {
		Data    *interface{} `json:"data"`
		Error   *string      `json:"error"`
		Success bool         `json:"success"`
		Message string       `json:"message"`
	}{Data: &data, Message: message, Error: nil, Success: true}}
}
