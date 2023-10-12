// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated-client/models"
)

// GetJobReader is a Reader for the GetJob structure.
type GetJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}] getJob", response, response.Code())
	}
}

// NewGetJobOK creates a GetJobOK with default headers values
func NewGetJobOK() *GetJobOK {
	return &GetJobOK{}
}

/*
GetJobOK describes a response with status code 200, with default header values.

scheduled job
*/
type GetJobOK struct {
	Payload *models.ScheduledJobSummary
}

// IsSuccess returns true when this get job o k response has a 2xx status code
func (o *GetJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get job o k response has a 3xx status code
func (o *GetJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job o k response has a 4xx status code
func (o *GetJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job o k response has a 5xx status code
func (o *GetJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get job o k response a status code equal to that given
func (o *GetJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get job o k response
func (o *GetJobOK) Code() int {
	return 200
}

func (o *GetJobOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}][%d] getJobOK  %+v", 200, o.Payload)
}

func (o *GetJobOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}][%d] getJobOK  %+v", 200, o.Payload)
}

func (o *GetJobOK) GetPayload() *models.ScheduledJobSummary {
	return o.Payload
}

func (o *GetJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledJobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobNotFound creates a GetJobNotFound with default headers values
func NewGetJobNotFound() *GetJobNotFound {
	return &GetJobNotFound{}
}

/*
GetJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetJobNotFound struct {
}

// IsSuccess returns true when this get job not found response has a 2xx status code
func (o *GetJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job not found response has a 3xx status code
func (o *GetJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job not found response has a 4xx status code
func (o *GetJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job not found response has a 5xx status code
func (o *GetJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get job not found response a status code equal to that given
func (o *GetJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get job not found response
func (o *GetJobNotFound) Code() int {
	return 404
}

func (o *GetJobNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}][%d] getJobNotFound ", 404)
}

func (o *GetJobNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}][%d] getJobNotFound ", 404)
}

func (o *GetJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
