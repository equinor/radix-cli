// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cli/generated/radixapi/models"
)

// NewTriggerPipelineDeployParams creates a new TriggerPipelineDeployParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerPipelineDeployParams() *TriggerPipelineDeployParams {
	return &TriggerPipelineDeployParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerPipelineDeployParamsWithTimeout creates a new TriggerPipelineDeployParams object
// with the ability to set a timeout on a request.
func NewTriggerPipelineDeployParamsWithTimeout(timeout time.Duration) *TriggerPipelineDeployParams {
	return &TriggerPipelineDeployParams{
		timeout: timeout,
	}
}

// NewTriggerPipelineDeployParamsWithContext creates a new TriggerPipelineDeployParams object
// with the ability to set a context for a request.
func NewTriggerPipelineDeployParamsWithContext(ctx context.Context) *TriggerPipelineDeployParams {
	return &TriggerPipelineDeployParams{
		Context: ctx,
	}
}

// NewTriggerPipelineDeployParamsWithHTTPClient creates a new TriggerPipelineDeployParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerPipelineDeployParamsWithHTTPClient(client *http.Client) *TriggerPipelineDeployParams {
	return &TriggerPipelineDeployParams{
		HTTPClient: client,
	}
}

/*
TriggerPipelineDeployParams contains all the parameters to send to the API endpoint

	for the trigger pipeline deploy operation.

	Typically these are written to a http.Request.
*/
type TriggerPipelineDeployParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* PipelineParametersDeploy.

	   Pipeline parameters
	*/
	PipelineParametersDeploy *models.PipelineParametersDeploy

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger pipeline deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineDeployParams) WithDefaults() *TriggerPipelineDeployParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger pipeline deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineDeployParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithTimeout(timeout time.Duration) *TriggerPipelineDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithContext(ctx context.Context) *TriggerPipelineDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithHTTPClient(client *http.Client) *TriggerPipelineDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithImpersonateGroup(impersonateGroup *string) *TriggerPipelineDeployParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithImpersonateUser(impersonateUser *string) *TriggerPipelineDeployParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithPipelineParametersDeploy adds the pipelineParametersDeploy to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithPipelineParametersDeploy(pipelineParametersDeploy *models.PipelineParametersDeploy) *TriggerPipelineDeployParams {
	o.SetPipelineParametersDeploy(pipelineParametersDeploy)
	return o
}

// SetPipelineParametersDeploy adds the pipelineParametersDeploy to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetPipelineParametersDeploy(pipelineParametersDeploy *models.PipelineParametersDeploy) {
	o.PipelineParametersDeploy = pipelineParametersDeploy
}

// WithAppName adds the appName to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) WithAppName(appName string) *TriggerPipelineDeployParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the trigger pipeline deploy params
func (o *TriggerPipelineDeployParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerPipelineDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// header param Impersonate-Group
		if err := r.SetHeaderParam("Impersonate-Group", *o.ImpersonateGroup); err != nil {
			return err
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}
	if o.PipelineParametersDeploy != nil {
		if err := r.SetBodyParam(o.PipelineParametersDeploy); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
