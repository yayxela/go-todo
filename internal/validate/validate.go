package validate

import (
	"regexp"
	"slices"

	"github.com/go-playground/validator/v10"

	"github.com/yayxela/go-todo/internal/values"
)

const (
	DateRule           = "date"
	TaskStatusDateRule = "task-status"
)

type Validator interface {
	Validate(data any) error
}
type customValidator struct {
	validate *validator.Validate
}

func (v *customValidator) Validate(data interface{}) error {
	return v.validate.Struct(data)
}

func New() Validator {
	v := validator.New()
	_ = v.RegisterValidation(DateRule, ValidateDate)
	_ = v.RegisterValidation(TaskStatusDateRule, ValidateTaskStatus)
	return &customValidator{
		validate: v,
	}
}

// ValidateDate ...
// Валидация даты задачи
func ValidateDate(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`, fl.Field().String())
	return ok
}

// ValidateTaskStatus ...
// Валидация статусов задачи
func ValidateTaskStatus(fl validator.FieldLevel) bool {
	return slices.Contains(values.TaskStatuses, values.TaskStatus(fl.Field().String()))
}
