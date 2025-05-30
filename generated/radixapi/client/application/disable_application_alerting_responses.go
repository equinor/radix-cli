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

// DisableApplicationAlertingReader is a Reader for the DisableApplicationAlerting structure.
type DisableApplicationAlertingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableApplicationAlertingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableApplicationAlertingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisableApplicationAlertingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDisableApplicationAlertingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisableApplicationAlertingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableApplicationAlertingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDisableApplicationAlertingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/alerting/disable] disableApplicationAlerting", response, response.Code())
	}
}

// NewDisableApplicationAlertingOK creates a DisableApplicationAlertingOK with default headers values
func NewDisableApplicationAlertingOK() *DisableApplicationAlertingOK {
	return &DisableApplicationAlertingOK{}
}

/*
DisableApplicationAlertingOK describes a response with status code 200, with default header values.

Successful disable alerting
*/
type DisableApplicationAlertingOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this disable application alerting o k response has a 2xx status code
func (o *DisableApplicationAlertingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable application alerting o k response has a 3xx status code
func (o *DisableApplicationAlertingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting o k response has a 4xx status code
func (o *DisableApplicationAlertingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable application alerting o k response has a 5xx status code
func (o *DisableApplicationAlertingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable application alerting o k response a status code equal to that given
func (o *DisableApplicationAlertingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disable application alerting o k response
func (o *DisableApplicationAlertingOK) Code() int {
	return 200
}

func (o *DisableApplicationAlertingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingOK %s", 200, payload)
}

func (o *DisableApplicationAlertingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingOK %s", 200, payload)
}

func (o *DisableApplicationAlertingOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *DisableApplicationAlertingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableApplicationAlertingBadRequest creates a DisableApplicationAlertingBadRequest with default headers values
func NewDisableApplicationAlertingBadRequest() *DisableApplicationAlertingBadRequest {
	return &DisableApplicationAlertingBadRequest{}
}

/*
DisableApplicationAlertingBadRequest describes a response with status code 400, with default header values.

Alerting already enabled
*/
type DisableApplicationAlertingBadRequest struct {
}

// IsSuccess returns true when this disable application alerting bad request response has a 2xx status code
func (o *DisableApplicationAlertingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable application alerting bad request response has a 3xx status code
func (o *DisableApplicationAlertingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting bad request response has a 4xx status code
func (o *DisableApplicationAlertingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable application alerting bad request response has a 5xx status code
func (o *DisableApplicationAlertingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this disable application alerting bad request response a status code equal to that given
func (o *DisableApplicationAlertingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the disable application alerting bad request response
func (o *DisableApplicationAlertingBadRequest) Code() int {
	return 400
}

func (o *DisableApplicationAlertingBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingBadRequest", 400)
}

func (o *DisableApplicationAlertingBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingBadRequest", 400)
}

func (o *DisableApplicationAlertingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableApplicationAlertingUnauthorized creates a DisableApplicationAlertingUnauthorized with default headers values
func NewDisableApplicationAlertingUnauthorized() *DisableApplicationAlertingUnauthorized {
	return &DisableApplicationAlertingUnauthorized{}
}

/*
DisableApplicationAlertingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DisableApplicationAlertingUnauthorized struct {
}

// IsSuccess returns true when this disable application alerting unauthorized response has a 2xx status code
func (o *DisableApplicationAlertingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable application alerting unauthorized response has a 3xx status code
func (o *DisableApplicationAlertingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting unauthorized response has a 4xx status code
func (o *DisableApplicationAlertingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable application alerting unauthorized response has a 5xx status code
func (o *DisableApplicationAlertingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this disable application alerting unauthorized response a status code equal to that given
func (o *DisableApplicationAlertingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the disable application alerting unauthorized response
func (o *DisableApplicationAlertingUnauthorized) Code() int {
	return 401
}

func (o *DisableApplicationAlertingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingUnauthorized", 401)
}

func (o *DisableApplicationAlertingUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingUnauthorized", 401)
}

func (o *DisableApplicationAlertingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableApplicationAlertingForbidden creates a DisableApplicationAlertingForbidden with default headers values
func NewDisableApplicationAlertingForbidden() *DisableApplicationAlertingForbidden {
	return &DisableApplicationAlertingForbidden{}
}

/*
DisableApplicationAlertingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DisableApplicationAlertingForbidden struct {
}

// IsSuccess returns true when this disable application alerting forbidden response has a 2xx status code
func (o *DisableApplicationAlertingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable application alerting forbidden response has a 3xx status code
func (o *DisableApplicationAlertingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting forbidden response has a 4xx status code
func (o *DisableApplicationAlertingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable application alerting forbidden response has a 5xx status code
func (o *DisableApplicationAlertingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this disable application alerting forbidden response a status code equal to that given
func (o *DisableApplicationAlertingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the disable application alerting forbidden response
func (o *DisableApplicationAlertingForbidden) Code() int {
	return 403
}

func (o *DisableApplicationAlertingForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingForbidden", 403)
}

func (o *DisableApplicationAlertingForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingForbidden", 403)
}

func (o *DisableApplicationAlertingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableApplicationAlertingNotFound creates a DisableApplicationAlertingNotFound with default headers values
func NewDisableApplicationAlertingNotFound() *DisableApplicationAlertingNotFound {
	return &DisableApplicationAlertingNotFound{}
}

/*
DisableApplicationAlertingNotFound describes a response with status code 404, with default header values.

Not found
*/
type DisableApplicationAlertingNotFound struct {
}

// IsSuccess returns true when this disable application alerting not found response has a 2xx status code
func (o *DisableApplicationAlertingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable application alerting not found response has a 3xx status code
func (o *DisableApplicationAlertingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting not found response has a 4xx status code
func (o *DisableApplicationAlertingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable application alerting not found response has a 5xx status code
func (o *DisableApplicationAlertingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this disable application alerting not found response a status code equal to that given
func (o *DisableApplicationAlertingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the disable application alerting not found response
func (o *DisableApplicationAlertingNotFound) Code() int {
	return 404
}

func (o *DisableApplicationAlertingNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingNotFound", 404)
}

func (o *DisableApplicationAlertingNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingNotFound", 404)
}

func (o *DisableApplicationAlertingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableApplicationAlertingInternalServerError creates a DisableApplicationAlertingInternalServerError with default headers values
func NewDisableApplicationAlertingInternalServerError() *DisableApplicationAlertingInternalServerError {
	return &DisableApplicationAlertingInternalServerError{}
}

/*
DisableApplicationAlertingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DisableApplicationAlertingInternalServerError struct {
}

// IsSuccess returns true when this disable application alerting internal server error response has a 2xx status code
func (o *DisableApplicationAlertingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable application alerting internal server error response has a 3xx status code
func (o *DisableApplicationAlertingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable application alerting internal server error response has a 4xx status code
func (o *DisableApplicationAlertingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable application alerting internal server error response has a 5xx status code
func (o *DisableApplicationAlertingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this disable application alerting internal server error response a status code equal to that given
func (o *DisableApplicationAlertingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the disable application alerting internal server error response
func (o *DisableApplicationAlertingInternalServerError) Code() int {
	return 500
}

func (o *DisableApplicationAlertingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingInternalServerError", 500)
}

func (o *DisableApplicationAlertingInternalServerError) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/alerting/disable][%d] disableApplicationAlertingInternalServerError", 500)
}

func (o *DisableApplicationAlertingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
