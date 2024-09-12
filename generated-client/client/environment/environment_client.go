// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new environment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new environment API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new environment API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for environment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeComponentSecret(params *ChangeComponentSecretParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeComponentSecretOK, error)

	CreateEnvironment(params *CreateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEnvironmentOK, error)

	DeleteEnvironment(params *DeleteEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEnvironmentOK, error)

	DisableEnvironmentAlerting(params *DisableEnvironmentAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableEnvironmentAlertingOK, error)

	EnableEnvironmentAlerting(params *EnableEnvironmentAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableEnvironmentAlertingOK, error)

	GetApplicationEnvironmentDeployments(params *GetApplicationEnvironmentDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationEnvironmentDeploymentsOK, error)

	GetAzureKeyVaultSecretVersions(params *GetAzureKeyVaultSecretVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAzureKeyVaultSecretVersionsOK, error)

	GetEnvironment(params *GetEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentOK, error)

	GetEnvironmentAlertingConfig(params *GetEnvironmentAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentAlertingConfigOK, error)

	GetEnvironmentEvents(params *GetEnvironmentEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentEventsOK, error)

	GetEnvironmentSummary(params *GetEnvironmentSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentSummaryOK, error)

	ResetManuallyScaledComponentsInEnvironment(params *ResetManuallyScaledComponentsInEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetManuallyScaledComponentsInEnvironmentOK, error)

	RestartEnvironment(params *RestartEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartEnvironmentOK, error)

	StartEnvironment(params *StartEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartEnvironmentOK, error)

	StopEnvironment(params *StopEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopEnvironmentOK, error)

	UpdateEnvironmentAlertingConfig(params *UpdateEnvironmentAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEnvironmentAlertingConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ChangeComponentSecret updates an application environment component secret
*/
func (a *Client) ChangeComponentSecret(params *ChangeComponentSecretParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeComponentSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeComponentSecretParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeComponentSecret",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeComponentSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeComponentSecretOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeComponentSecret: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateEnvironment creates application environment
*/
func (a *Client) CreateEnvironment(params *CreateEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createEnvironment",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteEnvironment deletes application environment
*/
func (a *Client) DeleteEnvironment(params *DeleteEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteEnvironment",
		Method:             "DELETE",
		PathPattern:        "/applications/{appName}/environments/{envName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisableEnvironmentAlerting disables alerting for an environment
*/
func (a *Client) DisableEnvironmentAlerting(params *DisableEnvironmentAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableEnvironmentAlertingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableEnvironmentAlertingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disableEnvironmentAlerting",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/alerting/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableEnvironmentAlertingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableEnvironmentAlertingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableEnvironmentAlerting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnableEnvironmentAlerting enables alerting for an environment
*/
func (a *Client) EnableEnvironmentAlerting(params *EnableEnvironmentAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableEnvironmentAlertingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableEnvironmentAlertingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableEnvironmentAlerting",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/alerting/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableEnvironmentAlertingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableEnvironmentAlertingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableEnvironmentAlerting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetApplicationEnvironmentDeployments lists the application environment deployments
*/
func (a *Client) GetApplicationEnvironmentDeployments(params *GetApplicationEnvironmentDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationEnvironmentDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationEnvironmentDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationEnvironmentDeployments",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationEnvironmentDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationEnvironmentDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationEnvironmentDeployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureKeyVaultSecretVersions gets azure key vault secret versions for a component
*/
func (a *Client) GetAzureKeyVaultSecretVersions(params *GetAzureKeyVaultSecretVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAzureKeyVaultSecretVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureKeyVaultSecretVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAzureKeyVaultSecretVersions",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/components/{componentName}/secrets/azure/keyvault/{azureKeyVaultName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureKeyVaultSecretVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureKeyVaultSecretVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureKeyVaultSecretVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironment gets details for an application environment
*/
func (a *Client) GetEnvironment(params *GetEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironment",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironmentAlertingConfig gets alerts configuration for an environment
*/
func (a *Client) GetEnvironmentAlertingConfig(params *GetEnvironmentAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentAlertingConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentAlertingConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironmentAlertingConfig",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/alerting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentAlertingConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentAlertingConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentAlertingConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironmentEvents lists events for an application environment
*/
func (a *Client) GetEnvironmentEvents(params *GetEnvironmentEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironmentEvents",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironmentSummary lists the environments for an application
*/
func (a *Client) GetEnvironmentSummary(params *GetEnvironmentSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironmentSummary",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetManuallyScaledComponentsInEnvironment resets all manually scaled component and resumes normal operation in environment
*/
func (a *Client) ResetManuallyScaledComponentsInEnvironment(params *ResetManuallyScaledComponentsInEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetManuallyScaledComponentsInEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetManuallyScaledComponentsInEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetManuallyScaledComponentsInEnvironment",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/reset-scale",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResetManuallyScaledComponentsInEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetManuallyScaledComponentsInEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetManuallyScaledComponentsInEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestartEnvironment restarts all components in the environment stops all running components in the environment pulls new images from image hub in radix configuration starts all components in the environment again using up to date image
*/
func (a *Client) RestartEnvironment(params *RestartEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartEnvironment",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restartEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartEnvironment deprecateds use reset scale instead that does the same thing but with better naming this method will be removed after 1 september 2025
*/
func (a *Client) StartEnvironment(params *StartEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startEnvironment",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopEnvironment stops all components in the environment
*/
func (a *Client) StopEnvironment(params *StopEnvironmentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopEnvironmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopEnvironment",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopEnvironmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopEnvironment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateEnvironmentAlertingConfig updates alerts configuration for an environment
*/
func (a *Client) UpdateEnvironmentAlertingConfig(params *UpdateEnvironmentAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEnvironmentAlertingConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentAlertingConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEnvironmentAlertingConfig",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/environments/{envName}/alerting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateEnvironmentAlertingConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateEnvironmentAlertingConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateEnvironmentAlertingConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
