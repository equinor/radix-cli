package model

import (
	"bytes"
	"github.com/go-openapi/runtime"
)

// ApiRequestResult Interface for API request result
type ApiRequestResult interface {
	ResponseReaderClientOption(operation *runtime.ClientOperation)
	GetResponseBody() string
	GetResponseCode() int
	GetResponseMessage() string
}
type apiRequestResult struct {
	baseReader      runtime.ClientResponseReader
	responseBody    string
	responseCode    int
	responseMessage string
}

// NewApiRequestResult creates a new ApiRequestResult
func NewApiRequestResult() ApiRequestResult {
	return &apiRequestResult{}
}

// ReadResponse reads a server response
func (reader *apiRequestResult) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	body := response.Body()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(body); err != nil {
		return nil, err
	}
	reader.responseBody = buf.String()
	reader.responseCode = response.Code()
	reader.responseMessage = response.Message()
	return reader.baseReader.ReadResponse(response, consumer)
}

// GetResponseBody returns the response body
func (reader *apiRequestResult) GetResponseBody() string {
	return reader.responseBody
}

// GetResponseCode returns the response code
func (reader *apiRequestResult) GetResponseCode() int {
	return reader.responseCode
}

// GetResponseMessage returns the response message
func (reader *apiRequestResult) GetResponseMessage() string {
	return reader.responseMessage
}

// ResponseReaderClientOption sets the response reader client option
func (reader *apiRequestResult) ResponseReaderClientOption(operation *runtime.ClientOperation) {
	reader.baseReader = operation.Reader
	operation.Reader = reader
}
