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

// NewGetReplicaEventsParams creates a new GetReplicaEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReplicaEventsParams() *GetReplicaEventsParams {
	return &GetReplicaEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicaEventsParamsWithTimeout creates a new GetReplicaEventsParams object
// with the ability to set a timeout on a request.
func NewGetReplicaEventsParamsWithTimeout(timeout time.Duration) *GetReplicaEventsParams {
	return &GetReplicaEventsParams{
		timeout: timeout,
	}
}

// NewGetReplicaEventsParamsWithContext creates a new GetReplicaEventsParams object
// with the ability to set a context for a request.
func NewGetReplicaEventsParamsWithContext(ctx context.Context) *GetReplicaEventsParams {
	return &GetReplicaEventsParams{
		Context: ctx,
	}
}

// NewGetReplicaEventsParamsWithHTTPClient creates a new GetReplicaEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReplicaEventsParamsWithHTTPClient(client *http.Client) *GetReplicaEventsParams {
	return &GetReplicaEventsParams{
		HTTPClient: client,
	}
}

/*
GetReplicaEventsParams contains all the parameters to send to the API endpoint

	for the get replica events operation.

	Typically these are written to a http.Request.
*/
type GetReplicaEventsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	/* ComponentName.

	   Name of component
	*/
	ComponentName string

	/* EnvName.

	   name of environment
	*/
	EnvName string

	/* PodName.

	   Name of pod
	*/
	PodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get replica events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicaEventsParams) WithDefaults() *GetReplicaEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get replica events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicaEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get replica events params
func (o *GetReplicaEventsParams) WithTimeout(timeout time.Duration) *GetReplicaEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replica events params
func (o *GetReplicaEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replica events params
func (o *GetReplicaEventsParams) WithContext(ctx context.Context) *GetReplicaEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replica events params
func (o *GetReplicaEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replica events params
func (o *GetReplicaEventsParams) WithHTTPClient(client *http.Client) *GetReplicaEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replica events params
func (o *GetReplicaEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get replica events params
func (o *GetReplicaEventsParams) WithImpersonateGroup(impersonateGroup *string) *GetReplicaEventsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get replica events params
func (o *GetReplicaEventsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get replica events params
func (o *GetReplicaEventsParams) WithImpersonateUser(impersonateUser *string) *GetReplicaEventsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get replica events params
func (o *GetReplicaEventsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get replica events params
func (o *GetReplicaEventsParams) WithAppName(appName string) *GetReplicaEventsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get replica events params
func (o *GetReplicaEventsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the get replica events params
func (o *GetReplicaEventsParams) WithComponentName(componentName string) *GetReplicaEventsParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the get replica events params
func (o *GetReplicaEventsParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the get replica events params
func (o *GetReplicaEventsParams) WithEnvName(envName string) *GetReplicaEventsParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get replica events params
func (o *GetReplicaEventsParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithPodName adds the podName to the get replica events params
func (o *GetReplicaEventsParams) WithPodName(podName string) *GetReplicaEventsParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the get replica events params
func (o *GetReplicaEventsParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicaEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}