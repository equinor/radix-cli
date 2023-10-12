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

// GetJobComponentDeploymentsReader is a Reader for the GetJobComponentDeployments structure.
type GetJobComponentDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobComponentDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobComponentDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetJobComponentDeploymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments] GetJobComponentDeployments", response, response.Code())
	}
}

// NewGetJobComponentDeploymentsOK creates a GetJobComponentDeploymentsOK with default headers values
func NewGetJobComponentDeploymentsOK() *GetJobComponentDeploymentsOK {
	return &GetJobComponentDeploymentsOK{}
}

/*
GetJobComponentDeploymentsOK describes a response with status code 200, with default header values.

Radix deployments
*/
type GetJobComponentDeploymentsOK struct {
	Payload []*models.DeploymentItem
}

// IsSuccess returns true when this get job component deployments o k response has a 2xx status code
func (o *GetJobComponentDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get job component deployments o k response has a 3xx status code
func (o *GetJobComponentDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job component deployments o k response has a 4xx status code
func (o *GetJobComponentDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job component deployments o k response has a 5xx status code
func (o *GetJobComponentDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get job component deployments o k response a status code equal to that given
func (o *GetJobComponentDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get job component deployments o k response
func (o *GetJobComponentDeploymentsOK) Code() int {
	return 200
}

func (o *GetJobComponentDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments][%d] getJobComponentDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetJobComponentDeploymentsOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments][%d] getJobComponentDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetJobComponentDeploymentsOK) GetPayload() []*models.DeploymentItem {
	return o.Payload
}

func (o *GetJobComponentDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobComponentDeploymentsNotFound creates a GetJobComponentDeploymentsNotFound with default headers values
func NewGetJobComponentDeploymentsNotFound() *GetJobComponentDeploymentsNotFound {
	return &GetJobComponentDeploymentsNotFound{}
}

/*
GetJobComponentDeploymentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetJobComponentDeploymentsNotFound struct {
}

// IsSuccess returns true when this get job component deployments not found response has a 2xx status code
func (o *GetJobComponentDeploymentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job component deployments not found response has a 3xx status code
func (o *GetJobComponentDeploymentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job component deployments not found response has a 4xx status code
func (o *GetJobComponentDeploymentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job component deployments not found response has a 5xx status code
func (o *GetJobComponentDeploymentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get job component deployments not found response a status code equal to that given
func (o *GetJobComponentDeploymentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get job component deployments not found response
func (o *GetJobComponentDeploymentsNotFound) Code() int {
	return 404
}

func (o *GetJobComponentDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments][%d] getJobComponentDeploymentsNotFound ", 404)
}

func (o *GetJobComponentDeploymentsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments][%d] getJobComponentDeploymentsNotFound ", 404)
}

func (o *GetJobComponentDeploymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
