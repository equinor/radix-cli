package cmd

import (
	"errors"
	"testing"

	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/stretchr/testify/assert"
)

func errorBody(errType, message, err string) *models.ErrorResponseBody {
	return &models.ErrorResponseBody{Type: &errType, Message: &message, Err: err}
}

func Test_registerApplicationError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "400 with message",
			err:      &platform.RegisterApplicationBadRequest{Payload: errorBody(models.ErrorResponseBodyTypeUser, "configuration item is required", "validation failed")},
			expected: "invalid application registration: configuration item is required (status code 400)",
		},
		{
			name:     "400 falls back to error details when message is empty",
			err:      &platform.RegisterApplicationBadRequest{Payload: errorBody(models.ErrorResponseBodyTypeUser, "", "validation failed")},
			expected: "invalid application registration: validation failed (status code 400)",
		},
		{
			name:     "400 with empty body",
			err:      &platform.RegisterApplicationBadRequest{},
			expected: "invalid application registration: status code 400",
		},
		{
			name:     "401",
			err:      &platform.RegisterApplicationUnauthorized{Payload: errorBody(models.ErrorResponseBodyTypeUser, "token expired", "")},
			expected: "unauthorized, run `rx login` or provide a valid token: token expired (status code 401)",
		},
		{
			name:     "403",
			err:      &platform.RegisterApplicationForbidden{Payload: errorBody(models.ErrorResponseBodyTypeForbidden, "user is not member of an admin group", "")},
			expected: "you are not allowed to register this application: user is not member of an admin group (status code 403)",
		},
		{
			name:     "409",
			err:      &platform.RegisterApplicationConflict{Payload: errorBody(models.ErrorResponseBodyTypeUser, "application my-app already exists", "")},
			expected: "application registration conflicts with an existing application: application my-app already exists (status code 409)",
		},
		{
			name:     "other errors are returned unchanged",
			err:      errors.New("dial tcp: connection refused"),
			expected: "dial tcp: connection refused",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, registerApplicationError(tt.err), tt.expected)
		})
	}
}
