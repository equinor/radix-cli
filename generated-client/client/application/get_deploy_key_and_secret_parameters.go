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
	"github.com/go-openapi/swag"
)

// NewGetDeployKeyAndSecretParams creates a new GetDeployKeyAndSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeployKeyAndSecretParams() *GetDeployKeyAndSecretParams {
	return &GetDeployKeyAndSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeployKeyAndSecretParamsWithTimeout creates a new GetDeployKeyAndSecretParams object
// with the ability to set a timeout on a request.
func NewGetDeployKeyAndSecretParamsWithTimeout(timeout time.Duration) *GetDeployKeyAndSecretParams {
	return &GetDeployKeyAndSecretParams{
		timeout: timeout,
	}
}

// NewGetDeployKeyAndSecretParamsWithContext creates a new GetDeployKeyAndSecretParams object
// with the ability to set a context for a request.
func NewGetDeployKeyAndSecretParamsWithContext(ctx context.Context) *GetDeployKeyAndSecretParams {
	return &GetDeployKeyAndSecretParams{
		Context: ctx,
	}
}

// NewGetDeployKeyAndSecretParamsWithHTTPClient creates a new GetDeployKeyAndSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeployKeyAndSecretParamsWithHTTPClient(client *http.Client) *GetDeployKeyAndSecretParams {
	return &GetDeployKeyAndSecretParams{
		HTTPClient: client,
	}
}

/*
GetDeployKeyAndSecretParams contains all the parameters to send to the API endpoint

	for the get deploy key and secret operation.

	Typically these are written to a http.Request.
*/
type GetDeployKeyAndSecretParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deploy key and secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeployKeyAndSecretParams) WithDefaults() *GetDeployKeyAndSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deploy key and secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeployKeyAndSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithTimeout(timeout time.Duration) *GetDeployKeyAndSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithContext(ctx context.Context) *GetDeployKeyAndSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithHTTPClient(client *http.Client) *GetDeployKeyAndSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithImpersonateGroup(impersonateGroup []string) *GetDeployKeyAndSecretParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithImpersonateUser(impersonateUser *string) *GetDeployKeyAndSecretParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) WithAppName(appName string) *GetDeployKeyAndSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get deploy key and secret params
func (o *GetDeployKeyAndSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeployKeyAndSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDeployKeyAndSecret binds the parameter Impersonate-Group
func (o *GetDeployKeyAndSecretParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
