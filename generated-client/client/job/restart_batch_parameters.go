// Code generated by go-swagger; DO NOT EDIT.

package job

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
	"github.com/go-openapi/swag"
)

// NewRestartBatchParams creates a new RestartBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartBatchParams() *RestartBatchParams {
	return &RestartBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartBatchParamsWithTimeout creates a new RestartBatchParams object
// with the ability to set a timeout on a request.
func NewRestartBatchParamsWithTimeout(timeout time.Duration) *RestartBatchParams {
	return &RestartBatchParams{
		timeout: timeout,
	}
}

// NewRestartBatchParamsWithContext creates a new RestartBatchParams object
// with the ability to set a context for a request.
func NewRestartBatchParamsWithContext(ctx context.Context) *RestartBatchParams {
	return &RestartBatchParams{
		Context: ctx,
	}
}

// NewRestartBatchParamsWithHTTPClient creates a new RestartBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartBatchParamsWithHTTPClient(client *http.Client) *RestartBatchParams {
	return &RestartBatchParams{
		HTTPClient: client,
	}
}

/*
RestartBatchParams contains all the parameters to send to the API endpoint

	for the restart batch operation.

	Typically these are written to a http.Request.
*/
type RestartBatchParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* BatchName.

	   Name of batch
	*/
	BatchName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartBatchParams) WithDefaults() *RestartBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart batch params
func (o *RestartBatchParams) WithTimeout(timeout time.Duration) *RestartBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart batch params
func (o *RestartBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart batch params
func (o *RestartBatchParams) WithContext(ctx context.Context) *RestartBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart batch params
func (o *RestartBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart batch params
func (o *RestartBatchParams) WithHTTPClient(client *http.Client) *RestartBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart batch params
func (o *RestartBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the restart batch params
func (o *RestartBatchParams) WithImpersonateGroup(impersonateGroup []string) *RestartBatchParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the restart batch params
func (o *RestartBatchParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the restart batch params
func (o *RestartBatchParams) WithImpersonateUser(impersonateUser *string) *RestartBatchParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the restart batch params
func (o *RestartBatchParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the restart batch params
func (o *RestartBatchParams) WithAppName(appName string) *RestartBatchParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the restart batch params
func (o *RestartBatchParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBatchName adds the batchName to the restart batch params
func (o *RestartBatchParams) WithBatchName(batchName string) *RestartBatchParams {
	o.SetBatchName(batchName)
	return o
}

// SetBatchName adds the batchName to the restart batch params
func (o *RestartBatchParams) SetBatchName(batchName string) {
	o.BatchName = batchName
}

// WithEnvName adds the envName to the restart batch params
func (o *RestartBatchParams) WithEnvName(envName string) *RestartBatchParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the restart batch params
func (o *RestartBatchParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the restart batch params
func (o *RestartBatchParams) WithJobComponentName(jobComponentName string) *RestartBatchParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the restart batch params
func (o *RestartBatchParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WriteToRequest writes these params to a swagger request
func (o *RestartBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// binding items for Impersonate-Group
		joinedImpersonateGroup := o.bindParamImpersonateGroup(reg)

		// header array param Impersonate-Group
		if len(joinedImpersonateGroup) > 0 {
			if err := r.SetHeaderParam("Impersonate-Group", joinedImpersonateGroup[0]); err != nil {
				return err
			}
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param batchName
	if err := r.SetPathParam("batchName", o.BatchName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRestartBatch binds the parameter Impersonate-Group
func (o *RestartBatchParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
	impersonateGroupIR := o.ImpersonateGroup

	var impersonateGroupIC []string
	for _, impersonateGroupIIR := range impersonateGroupIR { // explode []string

		impersonateGroupIIV := impersonateGroupIIR // string as string
		impersonateGroupIC = append(impersonateGroupIC, impersonateGroupIIV)
	}

	// items.CollectionFormat: ""
	impersonateGroupIS := swag.JoinByFormat(impersonateGroupIC, "")

	return impersonateGroupIS
}
