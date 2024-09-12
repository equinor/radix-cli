// Code generated by go-swagger; DO NOT EDIT.

package environment

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
)

// NewResetManuallyScaledComponentsInEnvironmentParams creates a new ResetManuallyScaledComponentsInEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetManuallyScaledComponentsInEnvironmentParams() *ResetManuallyScaledComponentsInEnvironmentParams {
	return &ResetManuallyScaledComponentsInEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetManuallyScaledComponentsInEnvironmentParamsWithTimeout creates a new ResetManuallyScaledComponentsInEnvironmentParams object
// with the ability to set a timeout on a request.
func NewResetManuallyScaledComponentsInEnvironmentParamsWithTimeout(timeout time.Duration) *ResetManuallyScaledComponentsInEnvironmentParams {
	return &ResetManuallyScaledComponentsInEnvironmentParams{
		timeout: timeout,
	}
}

// NewResetManuallyScaledComponentsInEnvironmentParamsWithContext creates a new ResetManuallyScaledComponentsInEnvironmentParams object
// with the ability to set a context for a request.
func NewResetManuallyScaledComponentsInEnvironmentParamsWithContext(ctx context.Context) *ResetManuallyScaledComponentsInEnvironmentParams {
	return &ResetManuallyScaledComponentsInEnvironmentParams{
		Context: ctx,
	}
}

// NewResetManuallyScaledComponentsInEnvironmentParamsWithHTTPClient creates a new ResetManuallyScaledComponentsInEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetManuallyScaledComponentsInEnvironmentParamsWithHTTPClient(client *http.Client) *ResetManuallyScaledComponentsInEnvironmentParams {
	return &ResetManuallyScaledComponentsInEnvironmentParams{
		HTTPClient: client,
	}
}

/*
ResetManuallyScaledComponentsInEnvironmentParams contains all the parameters to send to the API endpoint

	for the reset manually scaled components in environment operation.

	Typically these are written to a http.Request.
*/
type ResetManuallyScaledComponentsInEnvironmentParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset manually scaled components in environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithDefaults() *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset manually scaled components in environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithTimeout(timeout time.Duration) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithContext(ctx context.Context) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithHTTPClient(client *http.Client) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithImpersonateGroup(impersonateGroup *string) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithImpersonateUser(impersonateUser *string) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithAppName(appName string) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WithEnvName(envName string) *ResetManuallyScaledComponentsInEnvironmentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the reset manually scaled components in environment params
func (o *ResetManuallyScaledComponentsInEnvironmentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *ResetManuallyScaledComponentsInEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
