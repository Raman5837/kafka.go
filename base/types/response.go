package types

// Http Response Structure
type HttpResponse struct {
	Data struct {
		Error   *string `json:"error"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
	}
}
