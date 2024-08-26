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

// UpdateEnvironmentAlertingConfigReader is a Reader for the UpdateEnvironmentAlertingConfig structure.
type UpdateEnvironmentAlertingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEnvironmentAlertingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEnvironmentAlertingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEnvironmentAlertingConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateEnvironmentAlertingConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEnvironmentAlertingConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEnvironmentAlertingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEnvironmentAlertingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /applications/{appName}/environments/{envName}/alerting] updateEnvironmentAlertingConfig", response, response.Code())
	}
}

// NewUpdateEnvironmentAlertingConfigOK creates a UpdateEnvironmentAlertingConfigOK with default headers values
func NewUpdateEnvironmentAlertingConfigOK() *UpdateEnvironmentAlertingConfigOK {
	return &UpdateEnvironmentAlertingConfigOK{}
}

/*
UpdateEnvironmentAlertingConfigOK describes a response with status code 200, with default header values.

Successful alerts config update
*/
type UpdateEnvironmentAlertingConfigOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this update environment alerting config o k response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update environment alerting config o k response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config o k response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update environment alerting config o k response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update environment alerting config o k response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update environment alerting config o k response
func (o *UpdateEnvironmentAlertingConfigOK) Code() int {
	return 200
}

func (o *UpdateEnvironmentAlertingConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigOK %s", 200, payload)
}

func (o *UpdateEnvironmentAlertingConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigOK %s", 200, payload)
}

func (o *UpdateEnvironmentAlertingConfigOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *UpdateEnvironmentAlertingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEnvironmentAlertingConfigBadRequest creates a UpdateEnvironmentAlertingConfigBadRequest with default headers values
func NewUpdateEnvironmentAlertingConfigBadRequest() *UpdateEnvironmentAlertingConfigBadRequest {
	return &UpdateEnvironmentAlertingConfigBadRequest{}
}

/*
UpdateEnvironmentAlertingConfigBadRequest describes a response with status code 400, with default header values.

Invalid configuration
*/
type UpdateEnvironmentAlertingConfigBadRequest struct {
}

// IsSuccess returns true when this update environment alerting config bad request response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update environment alerting config bad request response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config bad request response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update environment alerting config bad request response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update environment alerting config bad request response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update environment alerting config bad request response
func (o *UpdateEnvironmentAlertingConfigBadRequest) Code() int {
	return 400
}

func (o *UpdateEnvironmentAlertingConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigBadRequest", 400)
}

func (o *UpdateEnvironmentAlertingConfigBadRequest) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigBadRequest", 400)
}

func (o *UpdateEnvironmentAlertingConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentAlertingConfigUnauthorized creates a UpdateEnvironmentAlertingConfigUnauthorized with default headers values
func NewUpdateEnvironmentAlertingConfigUnauthorized() *UpdateEnvironmentAlertingConfigUnauthorized {
	return &UpdateEnvironmentAlertingConfigUnauthorized{}
}

/*
UpdateEnvironmentAlertingConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateEnvironmentAlertingConfigUnauthorized struct {
}

// IsSuccess returns true when this update environment alerting config unauthorized response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update environment alerting config unauthorized response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config unauthorized response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update environment alerting config unauthorized response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update environment alerting config unauthorized response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update environment alerting config unauthorized response
func (o *UpdateEnvironmentAlertingConfigUnauthorized) Code() int {
	return 401
}

func (o *UpdateEnvironmentAlertingConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigUnauthorized", 401)
}

func (o *UpdateEnvironmentAlertingConfigUnauthorized) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigUnauthorized", 401)
}

func (o *UpdateEnvironmentAlertingConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentAlertingConfigForbidden creates a UpdateEnvironmentAlertingConfigForbidden with default headers values
func NewUpdateEnvironmentAlertingConfigForbidden() *UpdateEnvironmentAlertingConfigForbidden {
	return &UpdateEnvironmentAlertingConfigForbidden{}
}

/*
UpdateEnvironmentAlertingConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateEnvironmentAlertingConfigForbidden struct {
}

// IsSuccess returns true when this update environment alerting config forbidden response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update environment alerting config forbidden response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config forbidden response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update environment alerting config forbidden response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update environment alerting config forbidden response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update environment alerting config forbidden response
func (o *UpdateEnvironmentAlertingConfigForbidden) Code() int {
	return 403
}

func (o *UpdateEnvironmentAlertingConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigForbidden", 403)
}

func (o *UpdateEnvironmentAlertingConfigForbidden) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigForbidden", 403)
}

func (o *UpdateEnvironmentAlertingConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentAlertingConfigNotFound creates a UpdateEnvironmentAlertingConfigNotFound with default headers values
func NewUpdateEnvironmentAlertingConfigNotFound() *UpdateEnvironmentAlertingConfigNotFound {
	return &UpdateEnvironmentAlertingConfigNotFound{}
}

/*
UpdateEnvironmentAlertingConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateEnvironmentAlertingConfigNotFound struct {
}

// IsSuccess returns true when this update environment alerting config not found response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update environment alerting config not found response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config not found response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update environment alerting config not found response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update environment alerting config not found response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update environment alerting config not found response
func (o *UpdateEnvironmentAlertingConfigNotFound) Code() int {
	return 404
}

func (o *UpdateEnvironmentAlertingConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigNotFound", 404)
}

func (o *UpdateEnvironmentAlertingConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigNotFound", 404)
}

func (o *UpdateEnvironmentAlertingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentAlertingConfigInternalServerError creates a UpdateEnvironmentAlertingConfigInternalServerError with default headers values
func NewUpdateEnvironmentAlertingConfigInternalServerError() *UpdateEnvironmentAlertingConfigInternalServerError {
	return &UpdateEnvironmentAlertingConfigInternalServerError{}
}

/*
UpdateEnvironmentAlertingConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateEnvironmentAlertingConfigInternalServerError struct {
}

// IsSuccess returns true when this update environment alerting config internal server error response has a 2xx status code
func (o *UpdateEnvironmentAlertingConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update environment alerting config internal server error response has a 3xx status code
func (o *UpdateEnvironmentAlertingConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update environment alerting config internal server error response has a 4xx status code
func (o *UpdateEnvironmentAlertingConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update environment alerting config internal server error response has a 5xx status code
func (o *UpdateEnvironmentAlertingConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update environment alerting config internal server error response a status code equal to that given
func (o *UpdateEnvironmentAlertingConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update environment alerting config internal server error response
func (o *UpdateEnvironmentAlertingConfigInternalServerError) Code() int {
	return 500
}

func (o *UpdateEnvironmentAlertingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigInternalServerError", 500)
}

func (o *UpdateEnvironmentAlertingConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/alerting][%d] updateEnvironmentAlertingConfigInternalServerError", 500)
}

func (o *UpdateEnvironmentAlertingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
