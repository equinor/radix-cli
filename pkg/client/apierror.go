package client

import (
	"errors"
	"strings"

	"github.com/equinor/radix-cli/generated/radixapi/models"
)

type errorResponsePayloader interface {
	GetPayload() *models.ErrorResponseBody
}

// HandleErrorResponseBody extracts and returns a human-readable error message from the API's ErrorResponseBody.
// If the error does not contain an ErrorResponseBody, it returns the original error.
func HandleErrorResponseBody(err error) error {
	payloader, ok := err.(errorResponsePayloader)
	if !ok {
		return err
	}

	msg := payloader.GetPayload().Message
	if msg != nil {
		errStr, _ := strings.CutPrefix(*msg, "Error: ")
		return errors.New(errStr)
	}

	errMsg := payloader.GetPayload().Err
	errStr, _ := strings.CutPrefix(errMsg, "Error: ")
	return errors.New(errStr)
}
