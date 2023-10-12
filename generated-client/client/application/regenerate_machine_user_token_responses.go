// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated-client/models"
)

// RegenerateMachineUserTokenReader is a Reader for the RegenerateMachineUserToken structure.
type RegenerateMachineUserTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateMachineUserTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenerateMachineUserTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegenerateMachineUserTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegenerateMachineUserTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegenerateMachineUserTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegenerateMachineUserTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegenerateMachineUserTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/regenerate-machine-user-token] regenerateMachineUserToken", response, response.Code())
	}
}

// NewRegenerateMachineUserTokenOK creates a RegenerateMachineUserTokenOK with default headers values
func NewRegenerateMachineUserTokenOK() *RegenerateMachineUserTokenOK {
	return &RegenerateMachineUserTokenOK{}
}

/*
RegenerateMachineUserTokenOK describes a response with status code 200, with default header values.

Successful regenerate machine-user token
*/
type RegenerateMachineUserTokenOK struct {
	Payload *models.MachineUser
}

// IsSuccess returns true when this regenerate machine user token o k response has a 2xx status code
func (o *RegenerateMachineUserTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this regenerate machine user token o k response has a 3xx status code
func (o *RegenerateMachineUserTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token o k response has a 4xx status code
func (o *RegenerateMachineUserTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this regenerate machine user token o k response has a 5xx status code
func (o *RegenerateMachineUserTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate machine user token o k response a status code equal to that given
func (o *RegenerateMachineUserTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the regenerate machine user token o k response
func (o *RegenerateMachineUserTokenOK) Code() int {
	return 200
}

func (o *RegenerateMachineUserTokenOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenOK  %+v", 200, o.Payload)
}

func (o *RegenerateMachineUserTokenOK) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenOK  %+v", 200, o.Payload)
}

func (o *RegenerateMachineUserTokenOK) GetPayload() *models.MachineUser {
	return o.Payload
}

func (o *RegenerateMachineUserTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateMachineUserTokenUnauthorized creates a RegenerateMachineUserTokenUnauthorized with default headers values
func NewRegenerateMachineUserTokenUnauthorized() *RegenerateMachineUserTokenUnauthorized {
	return &RegenerateMachineUserTokenUnauthorized{}
}

/*
RegenerateMachineUserTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegenerateMachineUserTokenUnauthorized struct {
}

// IsSuccess returns true when this regenerate machine user token unauthorized response has a 2xx status code
func (o *RegenerateMachineUserTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate machine user token unauthorized response has a 3xx status code
func (o *RegenerateMachineUserTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token unauthorized response has a 4xx status code
func (o *RegenerateMachineUserTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate machine user token unauthorized response has a 5xx status code
func (o *RegenerateMachineUserTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate machine user token unauthorized response a status code equal to that given
func (o *RegenerateMachineUserTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the regenerate machine user token unauthorized response
func (o *RegenerateMachineUserTokenUnauthorized) Code() int {
	return 401
}

func (o *RegenerateMachineUserTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenUnauthorized ", 401)
}

func (o *RegenerateMachineUserTokenUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenUnauthorized ", 401)
}

func (o *RegenerateMachineUserTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateMachineUserTokenForbidden creates a RegenerateMachineUserTokenForbidden with default headers values
func NewRegenerateMachineUserTokenForbidden() *RegenerateMachineUserTokenForbidden {
	return &RegenerateMachineUserTokenForbidden{}
}

/*
RegenerateMachineUserTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RegenerateMachineUserTokenForbidden struct {
}

// IsSuccess returns true when this regenerate machine user token forbidden response has a 2xx status code
func (o *RegenerateMachineUserTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate machine user token forbidden response has a 3xx status code
func (o *RegenerateMachineUserTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token forbidden response has a 4xx status code
func (o *RegenerateMachineUserTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate machine user token forbidden response has a 5xx status code
func (o *RegenerateMachineUserTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate machine user token forbidden response a status code equal to that given
func (o *RegenerateMachineUserTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the regenerate machine user token forbidden response
func (o *RegenerateMachineUserTokenForbidden) Code() int {
	return 403
}

func (o *RegenerateMachineUserTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenForbidden ", 403)
}

func (o *RegenerateMachineUserTokenForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenForbidden ", 403)
}

func (o *RegenerateMachineUserTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateMachineUserTokenNotFound creates a RegenerateMachineUserTokenNotFound with default headers values
func NewRegenerateMachineUserTokenNotFound() *RegenerateMachineUserTokenNotFound {
	return &RegenerateMachineUserTokenNotFound{}
}

/*
RegenerateMachineUserTokenNotFound describes a response with status code 404, with default header values.

Not found
*/
type RegenerateMachineUserTokenNotFound struct {
}

// IsSuccess returns true when this regenerate machine user token not found response has a 2xx status code
func (o *RegenerateMachineUserTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate machine user token not found response has a 3xx status code
func (o *RegenerateMachineUserTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token not found response has a 4xx status code
func (o *RegenerateMachineUserTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate machine user token not found response has a 5xx status code
func (o *RegenerateMachineUserTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate machine user token not found response a status code equal to that given
func (o *RegenerateMachineUserTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the regenerate machine user token not found response
func (o *RegenerateMachineUserTokenNotFound) Code() int {
	return 404
}

func (o *RegenerateMachineUserTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenNotFound ", 404)
}

func (o *RegenerateMachineUserTokenNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenNotFound ", 404)
}

func (o *RegenerateMachineUserTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateMachineUserTokenConflict creates a RegenerateMachineUserTokenConflict with default headers values
func NewRegenerateMachineUserTokenConflict() *RegenerateMachineUserTokenConflict {
	return &RegenerateMachineUserTokenConflict{}
}

/*
RegenerateMachineUserTokenConflict describes a response with status code 409, with default header values.

Conflict
*/
type RegenerateMachineUserTokenConflict struct {
}

// IsSuccess returns true when this regenerate machine user token conflict response has a 2xx status code
func (o *RegenerateMachineUserTokenConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate machine user token conflict response has a 3xx status code
func (o *RegenerateMachineUserTokenConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token conflict response has a 4xx status code
func (o *RegenerateMachineUserTokenConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate machine user token conflict response has a 5xx status code
func (o *RegenerateMachineUserTokenConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate machine user token conflict response a status code equal to that given
func (o *RegenerateMachineUserTokenConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the regenerate machine user token conflict response
func (o *RegenerateMachineUserTokenConflict) Code() int {
	return 409
}

func (o *RegenerateMachineUserTokenConflict) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenConflict ", 409)
}

func (o *RegenerateMachineUserTokenConflict) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenConflict ", 409)
}

func (o *RegenerateMachineUserTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateMachineUserTokenInternalServerError creates a RegenerateMachineUserTokenInternalServerError with default headers values
func NewRegenerateMachineUserTokenInternalServerError() *RegenerateMachineUserTokenInternalServerError {
	return &RegenerateMachineUserTokenInternalServerError{}
}

/*
RegenerateMachineUserTokenInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RegenerateMachineUserTokenInternalServerError struct {
}

// IsSuccess returns true when this regenerate machine user token internal server error response has a 2xx status code
func (o *RegenerateMachineUserTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate machine user token internal server error response has a 3xx status code
func (o *RegenerateMachineUserTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate machine user token internal server error response has a 4xx status code
func (o *RegenerateMachineUserTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this regenerate machine user token internal server error response has a 5xx status code
func (o *RegenerateMachineUserTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this regenerate machine user token internal server error response a status code equal to that given
func (o *RegenerateMachineUserTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the regenerate machine user token internal server error response
func (o *RegenerateMachineUserTokenInternalServerError) Code() int {
	return 500
}

func (o *RegenerateMachineUserTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenInternalServerError ", 500)
}

func (o *RegenerateMachineUserTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenInternalServerError ", 500)
}

func (o *RegenerateMachineUserTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
