package validate

import (
	"reflect"
	"slices"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yayxela/go-todo/internal/values"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DateRule           = "date"
	TaskStatusDateRule = "task-status"
	ObjectID           = "object-id"
)

type Validator interface {
	Validate(data any) error
}
type customValidator struct {
	validate *validator.Validate
}

func (v *customValidator) Validate(data any) error {
	return v.validate.Struct(data)
}

// Option ..
// Опциональные параметры для валидатора
type Option func(v *customValidator)

func WithDateRule() func(v *customValidator) {
	return func(v *customValidator) {
		_ = v.validate.RegisterValidation(DateRule, validateDate)
	}
}

func WithTaskStatusRule() func(v *customValidator) {
	return func(v *customValidator) {
		_ = v.validate.RegisterValidation(TaskStatusDateRule, validateTaskStatus)
	}
}

func WithObjectIDRule() func(v *customValidator) {
	return func(v *customValidator) {
		_ = v.validate.RegisterValidation(ObjectID, validateObjectID)
	}
}

func WithNamedFields() func(v *customValidator) {
	return func(v *customValidator) {
		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Name
			switch {
			case field.Tag.Get("name") != "":
				name = field.Tag.Get("name")
			case field.Tag.Get("form") != "":
				name = field.Tag.Get("form")
			case field.Tag.Get("json") != "":
				name = field.Tag.Get("json")
			}
			return name
		})
	}
}

func New(opts ...Option) Validator {
	v := validator.New()
	cv := &customValidator{
		validate: v,
	}

	for _, opt := range opts {
		opt(cv)
	}
	return cv
}

func Default() Validator {
	return New(
		WithDateRule(),
		WithTaskStatusRule(),
		WithObjectIDRule(),
		WithNamedFields(),
	)
}

// validateDate ...
// Валидация даты задачи
func validateDate(fl validator.FieldLevel) bool {
	// ok, _ := regexp.MatchString(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`, fl.Field().String())
	_, err := time.Parse(time.DateOnly, fl.Field().String())
	return err == nil
}

// validateTaskStatus ...
// Валидация статусов задачи
func validateTaskStatus(fl validator.FieldLevel) bool {
	return slices.Contains(values.TaskStatuses, values.TaskStatus(fl.Field().String()))
}

// validateObjectID ...
// Валидация ObjectID
func validateObjectID(fl validator.FieldLevel) bool {
	_, err := primitive.ObjectIDFromHex(fl.Field().String())
	return err == nil
}
