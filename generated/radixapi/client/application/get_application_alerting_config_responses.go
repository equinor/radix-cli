// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated/radixapi/models"
)

// GetApplicationAlertingConfigReader is a Reader for the GetApplicationAlertingConfig structure.
type GetApplicationAlertingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationAlertingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationAlertingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationAlertingConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationAlertingConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationAlertingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApplicationAlertingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/alerting] getApplicationAlertingConfig", response, response.Code())
	}
}

// NewGetApplicationAlertingConfigOK creates a GetApplicationAlertingConfigOK with default headers values
func NewGetApplicationAlertingConfigOK() *GetApplicationAlertingConfigOK {
	return &GetApplicationAlertingConfigOK{}
}

/*
GetApplicationAlertingConfigOK describes a response with status code 200, with default header values.

Successful get alerts config
*/
type GetApplicationAlertingConfigOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this get application alerting config o k response has a 2xx status code
func (o *GetApplicationAlertingConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application alerting config o k response has a 3xx status code
func (o *GetApplicationAlertingConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application alerting config o k response has a 4xx status code
func (o *GetApplicationAlertingConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application alerting config o k response has a 5xx status code
func (o *GetApplicationAlertingConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application alerting config o k response a status code equal to that given
func (o *GetApplicationAlertingConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application alerting config o k response
func (o *GetApplicationAlertingConfigOK) Code() int {
	return 200
}

func (o *GetApplicationAlertingConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigOK %s", 200, payload)
}

func (o *GetApplicationAlertingConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigOK %s", 200, payload)
}

func (o *GetApplicationAlertingConfigOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *GetApplicationAlertingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationAlertingConfigUnauthorized creates a GetApplicationAlertingConfigUnauthorized with default headers values
func NewGetApplicationAlertingConfigUnauthorized() *GetApplicationAlertingConfigUnauthorized {
	return &GetApplicationAlertingConfigUnauthorized{}
}

/*
GetApplicationAlertingConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationAlertingConfigUnauthorized struct {
}

// IsSuccess returns true when this get application alerting config unauthorized response has a 2xx status code
func (o *GetApplicationAlertingConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application alerting config unauthorized response has a 3xx status code
func (o *GetApplicationAlertingConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application alerting config unauthorized response has a 4xx status code
func (o *GetApplicationAlertingConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application alerting config unauthorized response has a 5xx status code
func (o *GetApplicationAlertingConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application alerting config unauthorized response a status code equal to that given
func (o *GetApplicationAlertingConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application alerting config unauthorized response
func (o *GetApplicationAlertingConfigUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationAlertingConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigUnauthorized", 401)
}

func (o *GetApplicationAlertingConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigUnauthorized", 401)
}

func (o *GetApplicationAlertingConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationAlertingConfigForbidden creates a GetApplicationAlertingConfigForbidden with default headers values
func NewGetApplicationAlertingConfigForbidden() *GetApplicationAlertingConfigForbidden {
	return &GetApplicationAlertingConfigForbidden{}
}

/*
GetApplicationAlertingConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationAlertingConfigForbidden struct {
}

// IsSuccess returns true when this get application alerting config forbidden response has a 2xx status code
func (o *GetApplicationAlertingConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application alerting config forbidden response has a 3xx status code
func (o *GetApplicationAlertingConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application alerting config forbidden response has a 4xx status code
func (o *GetApplicationAlertingConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application alerting config forbidden response has a 5xx status code
func (o *GetApplicationAlertingConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get application alerting config forbidden response a status code equal to that given
func (o *GetApplicationAlertingConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get application alerting config forbidden response
func (o *GetApplicationAlertingConfigForbidden) Code() int {
	return 403
}

func (o *GetApplicationAlertingConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigForbidden", 403)
}

func (o *GetApplicationAlertingConfigForbidden) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigForbidden", 403)
}

func (o *GetApplicationAlertingConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationAlertingConfigNotFound creates a GetApplicationAlertingConfigNotFound with default headers values
func NewGetApplicationAlertingConfigNotFound() *GetApplicationAlertingConfigNotFound {
	return &GetApplicationAlertingConfigNotFound{}
}

/*
GetApplicationAlertingConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetApplicationAlertingConfigNotFound struct {
}

// IsSuccess returns true when this get application alerting config not found response has a 2xx status code
func (o *GetApplicationAlertingConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application alerting config not found response has a 3xx status code
func (o *GetApplicationAlertingConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application alerting config not found response has a 4xx status code
func (o *GetApplicationAlertingConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application alerting config not found response has a 5xx status code
func (o *GetApplicationAlertingConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application alerting config not found response a status code equal to that given
func (o *GetApplicationAlertingConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application alerting config not found response
func (o *GetApplicationAlertingConfigNotFound) Code() int {
	return 404
}

func (o *GetApplicationAlertingConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigNotFound", 404)
}

func (o *GetApplicationAlertingConfigNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigNotFound", 404)
}

func (o *GetApplicationAlertingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationAlertingConfigInternalServerError creates a GetApplicationAlertingConfigInternalServerError with default headers values
func NewGetApplicationAlertingConfigInternalServerError() *GetApplicationAlertingConfigInternalServerError {
	return &GetApplicationAlertingConfigInternalServerError{}
}

/*
GetApplicationAlertingConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetApplicationAlertingConfigInternalServerError struct {
}

// IsSuccess returns true when this get application alerting config internal server error response has a 2xx status code
func (o *GetApplicationAlertingConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application alerting config internal server error response has a 3xx status code
func (o *GetApplicationAlertingConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application alerting config internal server error response has a 4xx status code
func (o *GetApplicationAlertingConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application alerting config internal server error response has a 5xx status code
func (o *GetApplicationAlertingConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get application alerting config internal server error response a status code equal to that given
func (o *GetApplicationAlertingConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get application alerting config internal server error response
func (o *GetApplicationAlertingConfigInternalServerError) Code() int {
	return 500
}

func (o *GetApplicationAlertingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigInternalServerError", 500)
}

func (o *GetApplicationAlertingConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/alerting][%d] getApplicationAlertingConfigInternalServerError", 500)
}

func (o *GetApplicationAlertingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
