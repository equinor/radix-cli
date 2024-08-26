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

	"github.com/equinor/radix-cli/generated-client/models"
)

// ChangeRegistrationDetailsReader is a Reader for the ChangeRegistrationDetails structure.
type ChangeRegistrationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeRegistrationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeRegistrationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeRegistrationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeRegistrationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeRegistrationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewChangeRegistrationDetailsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /applications/{appName}] changeRegistrationDetails", response, response.Code())
	}
}

// NewChangeRegistrationDetailsOK creates a ChangeRegistrationDetailsOK with default headers values
func NewChangeRegistrationDetailsOK() *ChangeRegistrationDetailsOK {
	return &ChangeRegistrationDetailsOK{}
}

/*
ChangeRegistrationDetailsOK describes a response with status code 200, with default header values.

Change registration operation result
*/
type ChangeRegistrationDetailsOK struct {
	Payload *models.ApplicationRegistrationUpsertResponse
}

// IsSuccess returns true when this change registration details o k response has a 2xx status code
func (o *ChangeRegistrationDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change registration details o k response has a 3xx status code
func (o *ChangeRegistrationDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change registration details o k response has a 4xx status code
func (o *ChangeRegistrationDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change registration details o k response has a 5xx status code
func (o *ChangeRegistrationDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change registration details o k response a status code equal to that given
func (o *ChangeRegistrationDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change registration details o k response
func (o *ChangeRegistrationDetailsOK) Code() int {
	return 200
}

func (o *ChangeRegistrationDetailsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsOK %s", 200, payload)
}

func (o *ChangeRegistrationDetailsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsOK %s", 200, payload)
}

func (o *ChangeRegistrationDetailsOK) GetPayload() *models.ApplicationRegistrationUpsertResponse {
	return o.Payload
}

func (o *ChangeRegistrationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationRegistrationUpsertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeRegistrationDetailsBadRequest creates a ChangeRegistrationDetailsBadRequest with default headers values
func NewChangeRegistrationDetailsBadRequest() *ChangeRegistrationDetailsBadRequest {
	return &ChangeRegistrationDetailsBadRequest{}
}

/*
ChangeRegistrationDetailsBadRequest describes a response with status code 400, with default header values.

Invalid application
*/
type ChangeRegistrationDetailsBadRequest struct {
}

// IsSuccess returns true when this change registration details bad request response has a 2xx status code
func (o *ChangeRegistrationDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change registration details bad request response has a 3xx status code
func (o *ChangeRegistrationDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change registration details bad request response has a 4xx status code
func (o *ChangeRegistrationDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this change registration details bad request response has a 5xx status code
func (o *ChangeRegistrationDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this change registration details bad request response a status code equal to that given
func (o *ChangeRegistrationDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the change registration details bad request response
func (o *ChangeRegistrationDetailsBadRequest) Code() int {
	return 400
}

func (o *ChangeRegistrationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsBadRequest", 400)
}

func (o *ChangeRegistrationDetailsBadRequest) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsBadRequest", 400)
}

func (o *ChangeRegistrationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsUnauthorized creates a ChangeRegistrationDetailsUnauthorized with default headers values
func NewChangeRegistrationDetailsUnauthorized() *ChangeRegistrationDetailsUnauthorized {
	return &ChangeRegistrationDetailsUnauthorized{}
}

/*
ChangeRegistrationDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangeRegistrationDetailsUnauthorized struct {
}

// IsSuccess returns true when this change registration details unauthorized response has a 2xx status code
func (o *ChangeRegistrationDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change registration details unauthorized response has a 3xx status code
func (o *ChangeRegistrationDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change registration details unauthorized response has a 4xx status code
func (o *ChangeRegistrationDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change registration details unauthorized response has a 5xx status code
func (o *ChangeRegistrationDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change registration details unauthorized response a status code equal to that given
func (o *ChangeRegistrationDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the change registration details unauthorized response
func (o *ChangeRegistrationDetailsUnauthorized) Code() int {
	return 401
}

func (o *ChangeRegistrationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsUnauthorized", 401)
}

func (o *ChangeRegistrationDetailsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsUnauthorized", 401)
}

func (o *ChangeRegistrationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsNotFound creates a ChangeRegistrationDetailsNotFound with default headers values
func NewChangeRegistrationDetailsNotFound() *ChangeRegistrationDetailsNotFound {
	return &ChangeRegistrationDetailsNotFound{}
}

/*
ChangeRegistrationDetailsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ChangeRegistrationDetailsNotFound struct {
}

// IsSuccess returns true when this change registration details not found response has a 2xx status code
func (o *ChangeRegistrationDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change registration details not found response has a 3xx status code
func (o *ChangeRegistrationDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change registration details not found response has a 4xx status code
func (o *ChangeRegistrationDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this change registration details not found response has a 5xx status code
func (o *ChangeRegistrationDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this change registration details not found response a status code equal to that given
func (o *ChangeRegistrationDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the change registration details not found response
func (o *ChangeRegistrationDetailsNotFound) Code() int {
	return 404
}

func (o *ChangeRegistrationDetailsNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsNotFound", 404)
}

func (o *ChangeRegistrationDetailsNotFound) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsNotFound", 404)
}

func (o *ChangeRegistrationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRegistrationDetailsConflict creates a ChangeRegistrationDetailsConflict with default headers values
func NewChangeRegistrationDetailsConflict() *ChangeRegistrationDetailsConflict {
	return &ChangeRegistrationDetailsConflict{}
}

/*
ChangeRegistrationDetailsConflict describes a response with status code 409, with default header values.

Conflict
*/
type ChangeRegistrationDetailsConflict struct {
}

// IsSuccess returns true when this change registration details conflict response has a 2xx status code
func (o *ChangeRegistrationDetailsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change registration details conflict response has a 3xx status code
func (o *ChangeRegistrationDetailsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change registration details conflict response has a 4xx status code
func (o *ChangeRegistrationDetailsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this change registration details conflict response has a 5xx status code
func (o *ChangeRegistrationDetailsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this change registration details conflict response a status code equal to that given
func (o *ChangeRegistrationDetailsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the change registration details conflict response
func (o *ChangeRegistrationDetailsConflict) Code() int {
	return 409
}

func (o *ChangeRegistrationDetailsConflict) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsConflict", 409)
}

func (o *ChangeRegistrationDetailsConflict) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}][%d] changeRegistrationDetailsConflict", 409)
}

func (o *ChangeRegistrationDetailsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
