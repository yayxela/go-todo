package validate

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/values"
)

func TestValidate_ValidateDate(t *testing.T) {
	validator := New(WithDateRule())
	cases := []struct {
		name   string
		input  dto.TaskRequest
		result bool
	}{
		{
			name: "success request",
			input: dto.TaskRequest{
				Title:    "test",
				ActiveAt: time.DateOnly,
			},
			result: true,
		}, {
			name: "wrong ActiveAt",
			input: dto.TaskRequest{
				Title:    "test",
				ActiveAt: time.DateTime,
			},
			result: false,
		}, {
			name: "success request. max title",
			input: dto.TaskRequest{
				Title:    strings.Repeat("0", 200),
				ActiveAt: time.DateOnly,
			},
			result: true,
		}, {
			name: "wrong title",
			input: dto.TaskRequest{
				Title:    strings.Repeat("0", 201),
				ActiveAt: time.DateOnly,
			},
			result: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Validate(tc.input)
			require.Equal(t, tc.result, err == nil)
		})
	}
}

func TestValidate_ValidateTaskStatus(t *testing.T) {
	validator := New(WithTaskStatusRule())
	cases := []struct {
		name   string
		input  dto.ListRequest
		result bool
	}{
		{
			name: "success request",
			input: dto.ListRequest{
				Status: string(values.Done),
			},
			result: true,
		}, {
			name: "success request",
			input: dto.ListRequest{
				Status: string(values.Active),
			},
			result: true,
		}, {
			name: "not existing status",
			input: dto.ListRequest{
				Status: "new",
			},
			result: false,
		}, {
			name:   "empty status",
			input:  dto.ListRequest{},
			result: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Validate(tc.input)
			require.Equal(t, tc.result, err == nil)
		})
	}
}

func TestValidate_ValidateObjectID(t *testing.T) {
	validator := New(WithObjectIDRule())
	cases := []struct {
		name   string
		input  dto.GetByID
		result bool
	}{
		{
			name: "success request",
			input: dto.GetByID{
				ID: primitive.NewObjectID().Hex(),
			},
			result: true,
		}, {
			name: "success request",
			input: dto.GetByID{
				ID: "661f6efdb6ecc3459c7e9600",
			},
			result: true,
		}, {
			name: "not valid id",
			input: dto.GetByID{
				ID: "test",
			},
			result: false,
		}, {
			name:   "without id",
			input:  dto.GetByID{},
			result: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Validate(tc.input)
			require.Equal(t, tc.result, err == nil)
		})
	}
}
