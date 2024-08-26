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

// TriggerPipelinePromoteReader is a Reader for the TriggerPipelinePromote structure.
type TriggerPipelinePromoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerPipelinePromoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerPipelinePromoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTriggerPipelinePromoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/pipelines/promote] triggerPipelinePromote", response, response.Code())
	}
}

// NewTriggerPipelinePromoteOK creates a TriggerPipelinePromoteOK with default headers values
func NewTriggerPipelinePromoteOK() *TriggerPipelinePromoteOK {
	return &TriggerPipelinePromoteOK{}
}

/*
TriggerPipelinePromoteOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelinePromoteOK struct {
	Payload *models.JobSummary
}

// IsSuccess returns true when this trigger pipeline promote o k response has a 2xx status code
func (o *TriggerPipelinePromoteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger pipeline promote o k response has a 3xx status code
func (o *TriggerPipelinePromoteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline promote o k response has a 4xx status code
func (o *TriggerPipelinePromoteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger pipeline promote o k response has a 5xx status code
func (o *TriggerPipelinePromoteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline promote o k response a status code equal to that given
func (o *TriggerPipelinePromoteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger pipeline promote o k response
func (o *TriggerPipelinePromoteOK) Code() int {
	return 200
}

func (o *TriggerPipelinePromoteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/promote][%d] triggerPipelinePromoteOK %s", 200, payload)
}

func (o *TriggerPipelinePromoteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/promote][%d] triggerPipelinePromoteOK %s", 200, payload)
}

func (o *TriggerPipelinePromoteOK) GetPayload() *models.JobSummary {
	return o.Payload
}

func (o *TriggerPipelinePromoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerPipelinePromoteNotFound creates a TriggerPipelinePromoteNotFound with default headers values
func NewTriggerPipelinePromoteNotFound() *TriggerPipelinePromoteNotFound {
	return &TriggerPipelinePromoteNotFound{}
}

/*
TriggerPipelinePromoteNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelinePromoteNotFound struct {
}

// IsSuccess returns true when this trigger pipeline promote not found response has a 2xx status code
func (o *TriggerPipelinePromoteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline promote not found response has a 3xx status code
func (o *TriggerPipelinePromoteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline promote not found response has a 4xx status code
func (o *TriggerPipelinePromoteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline promote not found response has a 5xx status code
func (o *TriggerPipelinePromoteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline promote not found response a status code equal to that given
func (o *TriggerPipelinePromoteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger pipeline promote not found response
func (o *TriggerPipelinePromoteNotFound) Code() int {
	return 404
}

func (o *TriggerPipelinePromoteNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/promote][%d] triggerPipelinePromoteNotFound", 404)
}

func (o *TriggerPipelinePromoteNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/promote][%d] triggerPipelinePromoteNotFound", 404)
}

func (o *TriggerPipelinePromoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
