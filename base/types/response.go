package types

// Http Response Structure
type HttpResponse struct {
	Data struct {
		Data    *interface{} `json:"data"`
		Error   *string      `json:"error"`
		Success bool         `json:"success"`
		Message string       `json:"message"`
	}
}
