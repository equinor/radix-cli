// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated-client/models"
)

// GetSearchApplicationsReader is a Reader for the GetSearchApplications structure.
type GetSearchApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSearchApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSearchApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSearchApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetSearchApplicationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSearchApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/_search] getSearchApplications", response, response.Code())
	}
}

// NewGetSearchApplicationsOK creates a GetSearchApplicationsOK with default headers values
func NewGetSearchApplicationsOK() *GetSearchApplicationsOK {
	return &GetSearchApplicationsOK{}
}

/*
GetSearchApplicationsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSearchApplicationsOK struct {
	Payload []*models.ApplicationSummary
}

// IsSuccess returns true when this get search applications o k response has a 2xx status code
func (o *GetSearchApplicationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get search applications o k response has a 3xx status code
func (o *GetSearchApplicationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications o k response has a 4xx status code
func (o *GetSearchApplicationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search applications o k response has a 5xx status code
func (o *GetSearchApplicationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get search applications o k response a status code equal to that given
func (o *GetSearchApplicationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get search applications o k response
func (o *GetSearchApplicationsOK) Code() int {
	return 200
}

func (o *GetSearchApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetSearchApplicationsOK) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetSearchApplicationsOK) GetPayload() []*models.ApplicationSummary {
	return o.Payload
}

func (o *GetSearchApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchApplicationsUnauthorized creates a GetSearchApplicationsUnauthorized with default headers values
func NewGetSearchApplicationsUnauthorized() *GetSearchApplicationsUnauthorized {
	return &GetSearchApplicationsUnauthorized{}
}

/*
GetSearchApplicationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSearchApplicationsUnauthorized struct {
}

// IsSuccess returns true when this get search applications unauthorized response has a 2xx status code
func (o *GetSearchApplicationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search applications unauthorized response has a 3xx status code
func (o *GetSearchApplicationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications unauthorized response has a 4xx status code
func (o *GetSearchApplicationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search applications unauthorized response has a 5xx status code
func (o *GetSearchApplicationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get search applications unauthorized response a status code equal to that given
func (o *GetSearchApplicationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get search applications unauthorized response
func (o *GetSearchApplicationsUnauthorized) Code() int {
	return 401
}

func (o *GetSearchApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsUnauthorized ", 401)
}

func (o *GetSearchApplicationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsUnauthorized ", 401)
}

func (o *GetSearchApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchApplicationsForbidden creates a GetSearchApplicationsForbidden with default headers values
func NewGetSearchApplicationsForbidden() *GetSearchApplicationsForbidden {
	return &GetSearchApplicationsForbidden{}
}

/*
GetSearchApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSearchApplicationsForbidden struct {
}

// IsSuccess returns true when this get search applications forbidden response has a 2xx status code
func (o *GetSearchApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search applications forbidden response has a 3xx status code
func (o *GetSearchApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications forbidden response has a 4xx status code
func (o *GetSearchApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search applications forbidden response has a 5xx status code
func (o *GetSearchApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get search applications forbidden response a status code equal to that given
func (o *GetSearchApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get search applications forbidden response
func (o *GetSearchApplicationsForbidden) Code() int {
	return 403
}

func (o *GetSearchApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsForbidden ", 403)
}

func (o *GetSearchApplicationsForbidden) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsForbidden ", 403)
}

func (o *GetSearchApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchApplicationsNotFound creates a GetSearchApplicationsNotFound with default headers values
func NewGetSearchApplicationsNotFound() *GetSearchApplicationsNotFound {
	return &GetSearchApplicationsNotFound{}
}

/*
GetSearchApplicationsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSearchApplicationsNotFound struct {
}

// IsSuccess returns true when this get search applications not found response has a 2xx status code
func (o *GetSearchApplicationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search applications not found response has a 3xx status code
func (o *GetSearchApplicationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications not found response has a 4xx status code
func (o *GetSearchApplicationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search applications not found response has a 5xx status code
func (o *GetSearchApplicationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get search applications not found response a status code equal to that given
func (o *GetSearchApplicationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get search applications not found response
func (o *GetSearchApplicationsNotFound) Code() int {
	return 404
}

func (o *GetSearchApplicationsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsNotFound ", 404)
}

func (o *GetSearchApplicationsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsNotFound ", 404)
}

func (o *GetSearchApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchApplicationsConflict creates a GetSearchApplicationsConflict with default headers values
func NewGetSearchApplicationsConflict() *GetSearchApplicationsConflict {
	return &GetSearchApplicationsConflict{}
}

/*
GetSearchApplicationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetSearchApplicationsConflict struct {
}

// IsSuccess returns true when this get search applications conflict response has a 2xx status code
func (o *GetSearchApplicationsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search applications conflict response has a 3xx status code
func (o *GetSearchApplicationsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications conflict response has a 4xx status code
func (o *GetSearchApplicationsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search applications conflict response has a 5xx status code
func (o *GetSearchApplicationsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get search applications conflict response a status code equal to that given
func (o *GetSearchApplicationsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get search applications conflict response
func (o *GetSearchApplicationsConflict) Code() int {
	return 409
}

func (o *GetSearchApplicationsConflict) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsConflict ", 409)
}

func (o *GetSearchApplicationsConflict) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsConflict ", 409)
}

func (o *GetSearchApplicationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchApplicationsInternalServerError creates a GetSearchApplicationsInternalServerError with default headers values
func NewGetSearchApplicationsInternalServerError() *GetSearchApplicationsInternalServerError {
	return &GetSearchApplicationsInternalServerError{}
}

/*
GetSearchApplicationsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSearchApplicationsInternalServerError struct {
}

// IsSuccess returns true when this get search applications internal server error response has a 2xx status code
func (o *GetSearchApplicationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search applications internal server error response has a 3xx status code
func (o *GetSearchApplicationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search applications internal server error response has a 4xx status code
func (o *GetSearchApplicationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search applications internal server error response has a 5xx status code
func (o *GetSearchApplicationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get search applications internal server error response a status code equal to that given
func (o *GetSearchApplicationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get search applications internal server error response
func (o *GetSearchApplicationsInternalServerError) Code() int {
	return 500
}

func (o *GetSearchApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsInternalServerError ", 500)
}

func (o *GetSearchApplicationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications/_search][%d] getSearchApplicationsInternalServerError ", 500)
}

func (o *GetSearchApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
