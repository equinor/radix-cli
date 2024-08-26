// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated-client/models"
)

// GetEnvironmentSummaryReader is a Reader for the GetEnvironmentSummary structure.
type GetEnvironmentSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmentSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnvironmentSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEnvironmentSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments] getEnvironmentSummary", response, response.Code())
	}
}

// NewGetEnvironmentSummaryOK creates a GetEnvironmentSummaryOK with default headers values
func NewGetEnvironmentSummaryOK() *GetEnvironmentSummaryOK {
	return &GetEnvironmentSummaryOK{}
}

/*
GetEnvironmentSummaryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetEnvironmentSummaryOK struct {
	Payload []*models.EnvironmentSummary
}

// IsSuccess returns true when this get environment summary o k response has a 2xx status code
func (o *GetEnvironmentSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get environment summary o k response has a 3xx status code
func (o *GetEnvironmentSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment summary o k response has a 4xx status code
func (o *GetEnvironmentSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get environment summary o k response has a 5xx status code
func (o *GetEnvironmentSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment summary o k response a status code equal to that given
func (o *GetEnvironmentSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get environment summary o k response
func (o *GetEnvironmentSummaryOK) Code() int {
	return 200
}

func (o *GetEnvironmentSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryOK %s", 200, payload)
}

func (o *GetEnvironmentSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryOK %s", 200, payload)
}

func (o *GetEnvironmentSummaryOK) GetPayload() []*models.EnvironmentSummary {
	return o.Payload
}

func (o *GetEnvironmentSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentSummaryUnauthorized creates a GetEnvironmentSummaryUnauthorized with default headers values
func NewGetEnvironmentSummaryUnauthorized() *GetEnvironmentSummaryUnauthorized {
	return &GetEnvironmentSummaryUnauthorized{}
}

/*
GetEnvironmentSummaryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEnvironmentSummaryUnauthorized struct {
}

// IsSuccess returns true when this get environment summary unauthorized response has a 2xx status code
func (o *GetEnvironmentSummaryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment summary unauthorized response has a 3xx status code
func (o *GetEnvironmentSummaryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment summary unauthorized response has a 4xx status code
func (o *GetEnvironmentSummaryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment summary unauthorized response has a 5xx status code
func (o *GetEnvironmentSummaryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment summary unauthorized response a status code equal to that given
func (o *GetEnvironmentSummaryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get environment summary unauthorized response
func (o *GetEnvironmentSummaryUnauthorized) Code() int {
	return 401
}

func (o *GetEnvironmentSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryUnauthorized", 401)
}

func (o *GetEnvironmentSummaryUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryUnauthorized", 401)
}

func (o *GetEnvironmentSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentSummaryNotFound creates a GetEnvironmentSummaryNotFound with default headers values
func NewGetEnvironmentSummaryNotFound() *GetEnvironmentSummaryNotFound {
	return &GetEnvironmentSummaryNotFound{}
}

/*
GetEnvironmentSummaryNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEnvironmentSummaryNotFound struct {
}

// IsSuccess returns true when this get environment summary not found response has a 2xx status code
func (o *GetEnvironmentSummaryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment summary not found response has a 3xx status code
func (o *GetEnvironmentSummaryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment summary not found response has a 4xx status code
func (o *GetEnvironmentSummaryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment summary not found response has a 5xx status code
func (o *GetEnvironmentSummaryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment summary not found response a status code equal to that given
func (o *GetEnvironmentSummaryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get environment summary not found response
func (o *GetEnvironmentSummaryNotFound) Code() int {
	return 404
}

func (o *GetEnvironmentSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryNotFound", 404)
}

func (o *GetEnvironmentSummaryNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryNotFound", 404)
}

func (o *GetEnvironmentSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
