// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopAllBatchesAndJobsForEnvironmentReader is a Reader for the StopAllBatchesAndJobsForEnvironment structure.
type StopAllBatchesAndJobsForEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAllBatchesAndJobsForEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopAllBatchesAndJobsForEnvironmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopAllBatchesAndJobsForEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopAllBatchesAndJobsForEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopAllBatchesAndJobsForEnvironmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopAllBatchesAndJobsForEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop] stopAllBatchesAndJobsForEnvironment", response, response.Code())
	}
}

// NewStopAllBatchesAndJobsForEnvironmentNoContent creates a StopAllBatchesAndJobsForEnvironmentNoContent with default headers values
func NewStopAllBatchesAndJobsForEnvironmentNoContent() *StopAllBatchesAndJobsForEnvironmentNoContent {
	return &StopAllBatchesAndJobsForEnvironmentNoContent{}
}

/*
StopAllBatchesAndJobsForEnvironmentNoContent describes a response with status code 204, with default header values.

Success
*/
type StopAllBatchesAndJobsForEnvironmentNoContent struct {
}

// IsSuccess returns true when this stop all batches and jobs for environment no content response has a 2xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop all batches and jobs for environment no content response has a 3xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all batches and jobs for environment no content response has a 4xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop all batches and jobs for environment no content response has a 5xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all batches and jobs for environment no content response a status code equal to that given
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the stop all batches and jobs for environment no content response
func (o *StopAllBatchesAndJobsForEnvironmentNoContent) Code() int {
	return 204
}

func (o *StopAllBatchesAndJobsForEnvironmentNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentNoContent", 204)
}

func (o *StopAllBatchesAndJobsForEnvironmentNoContent) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentNoContent", 204)
}

func (o *StopAllBatchesAndJobsForEnvironmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllBatchesAndJobsForEnvironmentBadRequest creates a StopAllBatchesAndJobsForEnvironmentBadRequest with default headers values
func NewStopAllBatchesAndJobsForEnvironmentBadRequest() *StopAllBatchesAndJobsForEnvironmentBadRequest {
	return &StopAllBatchesAndJobsForEnvironmentBadRequest{}
}

/*
StopAllBatchesAndJobsForEnvironmentBadRequest describes a response with status code 400, with default header values.

Invalid batch
*/
type StopAllBatchesAndJobsForEnvironmentBadRequest struct {
}

// IsSuccess returns true when this stop all batches and jobs for environment bad request response has a 2xx status code
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all batches and jobs for environment bad request response has a 3xx status code
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all batches and jobs for environment bad request response has a 4xx status code
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all batches and jobs for environment bad request response has a 5xx status code
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all batches and jobs for environment bad request response a status code equal to that given
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stop all batches and jobs for environment bad request response
func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) Code() int {
	return 400
}

func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentBadRequest", 400)
}

func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentBadRequest", 400)
}

func (o *StopAllBatchesAndJobsForEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllBatchesAndJobsForEnvironmentUnauthorized creates a StopAllBatchesAndJobsForEnvironmentUnauthorized with default headers values
func NewStopAllBatchesAndJobsForEnvironmentUnauthorized() *StopAllBatchesAndJobsForEnvironmentUnauthorized {
	return &StopAllBatchesAndJobsForEnvironmentUnauthorized{}
}

/*
StopAllBatchesAndJobsForEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopAllBatchesAndJobsForEnvironmentUnauthorized struct {
}

// IsSuccess returns true when this stop all batches and jobs for environment unauthorized response has a 2xx status code
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all batches and jobs for environment unauthorized response has a 3xx status code
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all batches and jobs for environment unauthorized response has a 4xx status code
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all batches and jobs for environment unauthorized response has a 5xx status code
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all batches and jobs for environment unauthorized response a status code equal to that given
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stop all batches and jobs for environment unauthorized response
func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) Code() int {
	return 401
}

func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentUnauthorized", 401)
}

func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentUnauthorized", 401)
}

func (o *StopAllBatchesAndJobsForEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllBatchesAndJobsForEnvironmentForbidden creates a StopAllBatchesAndJobsForEnvironmentForbidden with default headers values
func NewStopAllBatchesAndJobsForEnvironmentForbidden() *StopAllBatchesAndJobsForEnvironmentForbidden {
	return &StopAllBatchesAndJobsForEnvironmentForbidden{}
}

/*
StopAllBatchesAndJobsForEnvironmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopAllBatchesAndJobsForEnvironmentForbidden struct {
}

// IsSuccess returns true when this stop all batches and jobs for environment forbidden response has a 2xx status code
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all batches and jobs for environment forbidden response has a 3xx status code
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all batches and jobs for environment forbidden response has a 4xx status code
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all batches and jobs for environment forbidden response has a 5xx status code
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all batches and jobs for environment forbidden response a status code equal to that given
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop all batches and jobs for environment forbidden response
func (o *StopAllBatchesAndJobsForEnvironmentForbidden) Code() int {
	return 403
}

func (o *StopAllBatchesAndJobsForEnvironmentForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentForbidden", 403)
}

func (o *StopAllBatchesAndJobsForEnvironmentForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentForbidden", 403)
}

func (o *StopAllBatchesAndJobsForEnvironmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllBatchesAndJobsForEnvironmentNotFound creates a StopAllBatchesAndJobsForEnvironmentNotFound with default headers values
func NewStopAllBatchesAndJobsForEnvironmentNotFound() *StopAllBatchesAndJobsForEnvironmentNotFound {
	return &StopAllBatchesAndJobsForEnvironmentNotFound{}
}

/*
StopAllBatchesAndJobsForEnvironmentNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopAllBatchesAndJobsForEnvironmentNotFound struct {
}

// IsSuccess returns true when this stop all batches and jobs for environment not found response has a 2xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all batches and jobs for environment not found response has a 3xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all batches and jobs for environment not found response has a 4xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all batches and jobs for environment not found response has a 5xx status code
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all batches and jobs for environment not found response a status code equal to that given
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop all batches and jobs for environment not found response
func (o *StopAllBatchesAndJobsForEnvironmentNotFound) Code() int {
	return 404
}

func (o *StopAllBatchesAndJobsForEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentNotFound", 404)
}

func (o *StopAllBatchesAndJobsForEnvironmentNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/stop][%d] stopAllBatchesAndJobsForEnvironmentNotFound", 404)
}

func (o *StopAllBatchesAndJobsForEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
