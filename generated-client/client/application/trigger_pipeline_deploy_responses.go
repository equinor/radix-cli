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

// TriggerPipelineDeployReader is a Reader for the TriggerPipelineDeploy structure.
type TriggerPipelineDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerPipelineDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerPipelineDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTriggerPipelineDeployForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTriggerPipelineDeployNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/pipelines/deploy] triggerPipelineDeploy", response, response.Code())
	}
}

// NewTriggerPipelineDeployOK creates a TriggerPipelineDeployOK with default headers values
func NewTriggerPipelineDeployOK() *TriggerPipelineDeployOK {
	return &TriggerPipelineDeployOK{}
}

/*
TriggerPipelineDeployOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelineDeployOK struct {
	Payload *models.JobSummary
}

// IsSuccess returns true when this trigger pipeline deploy o k response has a 2xx status code
func (o *TriggerPipelineDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger pipeline deploy o k response has a 3xx status code
func (o *TriggerPipelineDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline deploy o k response has a 4xx status code
func (o *TriggerPipelineDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger pipeline deploy o k response has a 5xx status code
func (o *TriggerPipelineDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline deploy o k response a status code equal to that given
func (o *TriggerPipelineDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger pipeline deploy o k response
func (o *TriggerPipelineDeployOK) Code() int {
	return 200
}

func (o *TriggerPipelineDeployOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployOK  %+v", 200, o.Payload)
}

func (o *TriggerPipelineDeployOK) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployOK  %+v", 200, o.Payload)
}

func (o *TriggerPipelineDeployOK) GetPayload() *models.JobSummary {
	return o.Payload
}

func (o *TriggerPipelineDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerPipelineDeployForbidden creates a TriggerPipelineDeployForbidden with default headers values
func NewTriggerPipelineDeployForbidden() *TriggerPipelineDeployForbidden {
	return &TriggerPipelineDeployForbidden{}
}

/*
TriggerPipelineDeployForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerPipelineDeployForbidden struct {
}

// IsSuccess returns true when this trigger pipeline deploy forbidden response has a 2xx status code
func (o *TriggerPipelineDeployForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline deploy forbidden response has a 3xx status code
func (o *TriggerPipelineDeployForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline deploy forbidden response has a 4xx status code
func (o *TriggerPipelineDeployForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline deploy forbidden response has a 5xx status code
func (o *TriggerPipelineDeployForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline deploy forbidden response a status code equal to that given
func (o *TriggerPipelineDeployForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trigger pipeline deploy forbidden response
func (o *TriggerPipelineDeployForbidden) Code() int {
	return 403
}

func (o *TriggerPipelineDeployForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployForbidden ", 403)
}

func (o *TriggerPipelineDeployForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployForbidden ", 403)
}

func (o *TriggerPipelineDeployForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerPipelineDeployNotFound creates a TriggerPipelineDeployNotFound with default headers values
func NewTriggerPipelineDeployNotFound() *TriggerPipelineDeployNotFound {
	return &TriggerPipelineDeployNotFound{}
}

/*
TriggerPipelineDeployNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelineDeployNotFound struct {
}

// IsSuccess returns true when this trigger pipeline deploy not found response has a 2xx status code
func (o *TriggerPipelineDeployNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline deploy not found response has a 3xx status code
func (o *TriggerPipelineDeployNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline deploy not found response has a 4xx status code
func (o *TriggerPipelineDeployNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline deploy not found response has a 5xx status code
func (o *TriggerPipelineDeployNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline deploy not found response a status code equal to that given
func (o *TriggerPipelineDeployNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger pipeline deploy not found response
func (o *TriggerPipelineDeployNotFound) Code() int {
	return 404
}

func (o *TriggerPipelineDeployNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployNotFound ", 404)
}

func (o *TriggerPipelineDeployNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/deploy][%d] triggerPipelineDeployNotFound ", 404)
}

func (o *TriggerPipelineDeployNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
