package dto

// ErrorResponse ...
// Кастоиное сообщение об ошибки
type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Tag     string `json:"tag,omitempty"`
}
