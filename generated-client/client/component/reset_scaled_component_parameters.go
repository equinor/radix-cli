// Code generated by go-swagger; DO NOT EDIT.

package component

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

// NewResetScaledComponentParams creates a new ResetScaledComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetScaledComponentParams() *ResetScaledComponentParams {
	return &ResetScaledComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetScaledComponentParamsWithTimeout creates a new ResetScaledComponentParams object
// with the ability to set a timeout on a request.
func NewResetScaledComponentParamsWithTimeout(timeout time.Duration) *ResetScaledComponentParams {
	return &ResetScaledComponentParams{
		timeout: timeout,
	}
}

// NewResetScaledComponentParamsWithContext creates a new ResetScaledComponentParams object
// with the ability to set a context for a request.
func NewResetScaledComponentParamsWithContext(ctx context.Context) *ResetScaledComponentParams {
	return &ResetScaledComponentParams{
		Context: ctx,
	}
}

// NewResetScaledComponentParamsWithHTTPClient creates a new ResetScaledComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetScaledComponentParamsWithHTTPClient(client *http.Client) *ResetScaledComponentParams {
	return &ResetScaledComponentParams{
		HTTPClient: client,
	}
}

/*
ResetScaledComponentParams contains all the parameters to send to the API endpoint

	for the reset scaled component operation.

	Typically these are written to a http.Request.
*/
type ResetScaledComponentParams struct {

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

	/* ComponentName.

	   Name of component
	*/
	ComponentName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset scaled component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetScaledComponentParams) WithDefaults() *ResetScaledComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset scaled component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetScaledComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset scaled component params
func (o *ResetScaledComponentParams) WithTimeout(timeout time.Duration) *ResetScaledComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset scaled component params
func (o *ResetScaledComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset scaled component params
func (o *ResetScaledComponentParams) WithContext(ctx context.Context) *ResetScaledComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset scaled component params
func (o *ResetScaledComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset scaled component params
func (o *ResetScaledComponentParams) WithHTTPClient(client *http.Client) *ResetScaledComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset scaled component params
func (o *ResetScaledComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the reset scaled component params
func (o *ResetScaledComponentParams) WithImpersonateGroup(impersonateGroup *string) *ResetScaledComponentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the reset scaled component params
func (o *ResetScaledComponentParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the reset scaled component params
func (o *ResetScaledComponentParams) WithImpersonateUser(impersonateUser *string) *ResetScaledComponentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the reset scaled component params
func (o *ResetScaledComponentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the reset scaled component params
func (o *ResetScaledComponentParams) WithAppName(appName string) *ResetScaledComponentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the reset scaled component params
func (o *ResetScaledComponentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the reset scaled component params
func (o *ResetScaledComponentParams) WithComponentName(componentName string) *ResetScaledComponentParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the reset scaled component params
func (o *ResetScaledComponentParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the reset scaled component params
func (o *ResetScaledComponentParams) WithEnvName(envName string) *ResetScaledComponentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the reset scaled component params
func (o *ResetScaledComponentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *ResetScaledComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param componentName
	if err := r.SetPathParam("componentName", o.ComponentName); err != nil {
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
