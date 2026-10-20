package client

import (
	"fmt"
	"strings"

	"github.com/equinor/radix-cli/generated/radixapi/models"
)

// ErrorResponseMessage returns the user-facing text from a Radix API error body, preferring Message
// and falling back to the underlying error details. Returns "" for a nil or empty body.
func ErrorResponseMessage(body *models.ErrorResponseBody) string {
	if body == nil {
		return ""
	}
	if body.Message != nil && strings.TrimSpace(*body.Message) != "" {
		return strings.TrimSpace(*body.Message)
	}
	return strings.TrimSpace(body.Err)
}

// ErrorResponseDetails formats a Radix API error body together with its status code, e.g.
// "configuration item is required (status code 400)". Degrades to "status code 400" when the body is empty.
func ErrorResponseDetails(body *models.ErrorResponseBody, statusCode int) string {
	if message := ErrorResponseMessage(body); message != "" {
		return fmt.Sprintf("%s (status code %d)", message, statusCode)
	}
	return fmt.Sprintf("status code %d", statusCode)
}
