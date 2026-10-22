package cmd

import (
	"errors"
	"testing"

	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/stretchr/testify/assert"
)

func errorBody(errType, message, err string) *models.ErrorResponseBody {
	body := &models.ErrorResponseBody{Type: &errType, Err: err}
	if message != "" {
		body.Message = &message
	}
	return body
}

func Test_handleRegisterApplicationError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "400 with message",
			err:      &platform.RegisterApplicationBadRequest{Payload: errorBody(models.ErrorResponseBodyTypeUser, "configuration item is required", "validation failed")},
			expected: "configuration item is required",
		},
		{
			name:     "400 falls back to error details when message is missing",
			err:      &platform.RegisterApplicationBadRequest{Payload: errorBody(models.ErrorResponseBodyTypeUser, "", "validation failed")},
			expected: "validation failed",
		},
		{
			name:     "401",
			err:      &platform.RegisterApplicationUnauthorized{Payload: errorBody(models.ErrorResponseBodyTypeUser, "token expired", "")},
			expected: "token expired",
		},
		{
			name:     "403",
			err:      &platform.RegisterApplicationForbidden{Payload: errorBody(models.ErrorResponseBodyTypeForbidden, "user is not member of an admin group", "")},
			expected: "user is not member of an admin group",
		},
		{
			name:     "409",
			err:      &platform.RegisterApplicationConflict{Payload: errorBody(models.ErrorResponseBodyTypeUser, "application my-app already exists", "")},
			expected: "application my-app already exists",
		},
		{
			name:     "the Error prefix is stripped from the message",
			err:      &platform.RegisterApplicationBadRequest{Payload: errorBody(models.ErrorResponseBodyTypeUser, "Error: invalid repository url", "")},
			expected: "invalid repository url",
		},
		{
			name:     "other errors are returned unchanged",
			err:      errors.New("dial tcp: connection refused"),
			expected: "dial tcp: connection refused",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, client.HandleErrorResponseBody(tt.err), tt.expected)
		})
	}
}
