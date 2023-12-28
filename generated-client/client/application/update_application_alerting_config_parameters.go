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

	"github.com/equinor/radix-cli/generated-client/models"
)

// NewUpdateApplicationAlertingConfigParams creates a new UpdateApplicationAlertingConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateApplicationAlertingConfigParams() *UpdateApplicationAlertingConfigParams {
	return &UpdateApplicationAlertingConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationAlertingConfigParamsWithTimeout creates a new UpdateApplicationAlertingConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateApplicationAlertingConfigParamsWithTimeout(timeout time.Duration) *UpdateApplicationAlertingConfigParams {
	return &UpdateApplicationAlertingConfigParams{
		timeout: timeout,
	}
}

// NewUpdateApplicationAlertingConfigParamsWithContext creates a new UpdateApplicationAlertingConfigParams object
// with the ability to set a context for a request.
func NewUpdateApplicationAlertingConfigParamsWithContext(ctx context.Context) *UpdateApplicationAlertingConfigParams {
	return &UpdateApplicationAlertingConfigParams{
		Context: ctx,
	}
}

// NewUpdateApplicationAlertingConfigParamsWithHTTPClient creates a new UpdateApplicationAlertingConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateApplicationAlertingConfigParamsWithHTTPClient(client *http.Client) *UpdateApplicationAlertingConfigParams {
	return &UpdateApplicationAlertingConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateApplicationAlertingConfigParams contains all the parameters to send to the API endpoint

	for the update application alerting config operation.

	Typically these are written to a http.Request.
*/
type UpdateApplicationAlertingConfigParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AlertsConfig.

	   Alerts configuration
	*/
	AlertsConfig *models.UpdateAlertingConfig

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update application alerting config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateApplicationAlertingConfigParams) WithDefaults() *UpdateApplicationAlertingConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update application alerting config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateApplicationAlertingConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithTimeout(timeout time.Duration) *UpdateApplicationAlertingConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithContext(ctx context.Context) *UpdateApplicationAlertingConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithHTTPClient(client *http.Client) *UpdateApplicationAlertingConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithImpersonateGroup(impersonateGroup *string) *UpdateApplicationAlertingConfigParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithImpersonateUser(impersonateUser *string) *UpdateApplicationAlertingConfigParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAlertsConfig adds the alertsConfig to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithAlertsConfig(alertsConfig *models.UpdateAlertingConfig) *UpdateApplicationAlertingConfigParams {
	o.SetAlertsConfig(alertsConfig)
	return o
}

// SetAlertsConfig adds the alertsConfig to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetAlertsConfig(alertsConfig *models.UpdateAlertingConfig) {
	o.AlertsConfig = alertsConfig
}

// WithAppName adds the appName to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) WithAppName(appName string) *UpdateApplicationAlertingConfigParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update application alerting config params
func (o *UpdateApplicationAlertingConfigParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationAlertingConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.AlertsConfig != nil {
		if err := r.SetBodyParam(o.AlertsConfig); err != nil {
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
