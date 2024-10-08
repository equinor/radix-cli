// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ResetScaledComponentReader is a Reader for the ResetScaledComponent structure.
type ResetScaledComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetScaledComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetScaledComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResetScaledComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetScaledComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale] resetScaledComponent", response, response.Code())
	}
}

// NewResetScaledComponentOK creates a ResetScaledComponentOK with default headers values
func NewResetScaledComponentOK() *ResetScaledComponentOK {
	return &ResetScaledComponentOK{}
}

/*
ResetScaledComponentOK describes a response with status code 200, with default header values.

Component started ok
*/
type ResetScaledComponentOK struct {
}

// IsSuccess returns true when this reset scaled component o k response has a 2xx status code
func (o *ResetScaledComponentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reset scaled component o k response has a 3xx status code
func (o *ResetScaledComponentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset scaled component o k response has a 4xx status code
func (o *ResetScaledComponentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset scaled component o k response has a 5xx status code
func (o *ResetScaledComponentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reset scaled component o k response a status code equal to that given
func (o *ResetScaledComponentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reset scaled component o k response
func (o *ResetScaledComponentOK) Code() int {
	return 200
}

func (o *ResetScaledComponentOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentOK", 200)
}

func (o *ResetScaledComponentOK) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentOK", 200)
}

func (o *ResetScaledComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetScaledComponentUnauthorized creates a ResetScaledComponentUnauthorized with default headers values
func NewResetScaledComponentUnauthorized() *ResetScaledComponentUnauthorized {
	return &ResetScaledComponentUnauthorized{}
}

/*
ResetScaledComponentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ResetScaledComponentUnauthorized struct {
}

// IsSuccess returns true when this reset scaled component unauthorized response has a 2xx status code
func (o *ResetScaledComponentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset scaled component unauthorized response has a 3xx status code
func (o *ResetScaledComponentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset scaled component unauthorized response has a 4xx status code
func (o *ResetScaledComponentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset scaled component unauthorized response has a 5xx status code
func (o *ResetScaledComponentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this reset scaled component unauthorized response a status code equal to that given
func (o *ResetScaledComponentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the reset scaled component unauthorized response
func (o *ResetScaledComponentUnauthorized) Code() int {
	return 401
}

func (o *ResetScaledComponentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentUnauthorized", 401)
}

func (o *ResetScaledComponentUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentUnauthorized", 401)
}

func (o *ResetScaledComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetScaledComponentNotFound creates a ResetScaledComponentNotFound with default headers values
func NewResetScaledComponentNotFound() *ResetScaledComponentNotFound {
	return &ResetScaledComponentNotFound{}
}

/*
ResetScaledComponentNotFound describes a response with status code 404, with default header values.

Not found
*/
type ResetScaledComponentNotFound struct {
}

// IsSuccess returns true when this reset scaled component not found response has a 2xx status code
func (o *ResetScaledComponentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset scaled component not found response has a 3xx status code
func (o *ResetScaledComponentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset scaled component not found response has a 4xx status code
func (o *ResetScaledComponentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset scaled component not found response has a 5xx status code
func (o *ResetScaledComponentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reset scaled component not found response a status code equal to that given
func (o *ResetScaledComponentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reset scaled component not found response
func (o *ResetScaledComponentNotFound) Code() int {
	return 404
}

func (o *ResetScaledComponentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentNotFound", 404)
}

func (o *ResetScaledComponentNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/reset-scale][%d] resetScaledComponentNotFound", 404)
}

func (o *ResetScaledComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
