// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated-client/models"
)

// GetApplicationEnvironmentDeploymentsReader is a Reader for the GetApplicationEnvironmentDeployments structure.
type GetApplicationEnvironmentDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationEnvironmentDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationEnvironmentDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationEnvironmentDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationEnvironmentDeploymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/deployments] getApplicationEnvironmentDeployments", response, response.Code())
	}
}

// NewGetApplicationEnvironmentDeploymentsOK creates a GetApplicationEnvironmentDeploymentsOK with default headers values
func NewGetApplicationEnvironmentDeploymentsOK() *GetApplicationEnvironmentDeploymentsOK {
	return &GetApplicationEnvironmentDeploymentsOK{}
}

/*
GetApplicationEnvironmentDeploymentsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetApplicationEnvironmentDeploymentsOK struct {
	Payload []*models.DeploymentSummary
}

// IsSuccess returns true when this get application environment deployments o k response has a 2xx status code
func (o *GetApplicationEnvironmentDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application environment deployments o k response has a 3xx status code
func (o *GetApplicationEnvironmentDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application environment deployments o k response has a 4xx status code
func (o *GetApplicationEnvironmentDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application environment deployments o k response has a 5xx status code
func (o *GetApplicationEnvironmentDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application environment deployments o k response a status code equal to that given
func (o *GetApplicationEnvironmentDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application environment deployments o k response
func (o *GetApplicationEnvironmentDeploymentsOK) Code() int {
	return 200
}

func (o *GetApplicationEnvironmentDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationEnvironmentDeploymentsOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationEnvironmentDeploymentsOK) GetPayload() []*models.DeploymentSummary {
	return o.Payload
}

func (o *GetApplicationEnvironmentDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationEnvironmentDeploymentsUnauthorized creates a GetApplicationEnvironmentDeploymentsUnauthorized with default headers values
func NewGetApplicationEnvironmentDeploymentsUnauthorized() *GetApplicationEnvironmentDeploymentsUnauthorized {
	return &GetApplicationEnvironmentDeploymentsUnauthorized{}
}

/*
GetApplicationEnvironmentDeploymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationEnvironmentDeploymentsUnauthorized struct {
}

// IsSuccess returns true when this get application environment deployments unauthorized response has a 2xx status code
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application environment deployments unauthorized response has a 3xx status code
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application environment deployments unauthorized response has a 4xx status code
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application environment deployments unauthorized response has a 5xx status code
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application environment deployments unauthorized response a status code equal to that given
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application environment deployments unauthorized response
func (o *GetApplicationEnvironmentDeploymentsUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationEnvironmentDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsUnauthorized ", 401)
}

func (o *GetApplicationEnvironmentDeploymentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsUnauthorized ", 401)
}

func (o *GetApplicationEnvironmentDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationEnvironmentDeploymentsNotFound creates a GetApplicationEnvironmentDeploymentsNotFound with default headers values
func NewGetApplicationEnvironmentDeploymentsNotFound() *GetApplicationEnvironmentDeploymentsNotFound {
	return &GetApplicationEnvironmentDeploymentsNotFound{}
}

/*
GetApplicationEnvironmentDeploymentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetApplicationEnvironmentDeploymentsNotFound struct {
}

// IsSuccess returns true when this get application environment deployments not found response has a 2xx status code
func (o *GetApplicationEnvironmentDeploymentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application environment deployments not found response has a 3xx status code
func (o *GetApplicationEnvironmentDeploymentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application environment deployments not found response has a 4xx status code
func (o *GetApplicationEnvironmentDeploymentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application environment deployments not found response has a 5xx status code
func (o *GetApplicationEnvironmentDeploymentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application environment deployments not found response a status code equal to that given
func (o *GetApplicationEnvironmentDeploymentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application environment deployments not found response
func (o *GetApplicationEnvironmentDeploymentsNotFound) Code() int {
	return 404
}

func (o *GetApplicationEnvironmentDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsNotFound ", 404)
}

func (o *GetApplicationEnvironmentDeploymentsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/deployments][%d] getApplicationEnvironmentDeploymentsNotFound ", 404)
}

func (o *GetApplicationEnvironmentDeploymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}