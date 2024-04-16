package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/logger"
	"github.com/yayxela/go-todo/internal/values"
)

type Middleware interface {
	Panic(c *gin.Context)
	Error(c *gin.Context)
}

type middleware struct {
	log logger.Logger
}

// Panic ...
// Обработка критический ошибок
func (m *middleware) Panic(c *gin.Context) {
	c.Next()
	if r := recover(); r != nil {
		var err error
		switch t := r.(type) {
		case string:
			err = values.NewHttpError(http.StatusInternalServerError, t)
		case error:
			err = values.NewHttpError(http.StatusInternalServerError, t.Error())
		default:
			err = values.NewHttpError(http.StatusInternalServerError, "api.unknown_err")
		}
		_ = c.Error(err)
	}
}

// Error ...
// Обработка ошибок
func (m *middleware) Error(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	var statusCode int
	errorResponse := make([]*dto.ErrorResponse, 0)
	for _, err := range c.Errors {
		var httpError *values.HttpError
		var validatorError validator.ValidationErrors

		if errors.As(err, &httpError) {
			// кастомные ошибки
			statusCode = httpError.Code()
			m.log.Error(err)
		} else if errors.As(err, &validatorError) {
			// ошибки валидатора
			for _, e := range validatorError {
				errorResponse = append(errorResponse, &dto.ErrorResponse{
					Message: e.Error(),
					Field:   e.Field(),
					Tag:     e.Tag(),
				})
			}
			statusCode = http.StatusBadRequest
		} else if errors.Is(err, mongo.ErrNoDocuments) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
			m.log.Error(err)
		}
	}
	if len(errorResponse) > 0 {
		c.AbortWithStatusJSON(statusCode, errorResponse)
	} else {
		c.AbortWithStatus(statusCode)
	}
}

func New(log logger.Logger) Middleware {
	return &middleware{
		log: log.Named("middleware"),
	}
}
