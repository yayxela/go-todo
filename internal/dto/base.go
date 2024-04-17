package dto

// ErrorResponse ...
// Кастомное сообщение об ошибке
type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Tag     string `json:"tag,omitempty"`
}
