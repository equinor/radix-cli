package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// MSALAuthProvider is an AuthProvider that uses MSAL
type MSALAuthProvider interface {
	Login(ctx context.Context) error
	Logout(ctx context.Context) error
	// WrapTransport(rt http.RoundTripper) http.RoundTripper
	runtime.ClientAuthInfoWriter
}

// NewMSALAuthProvider creates a new MSALAuthProvider
func NewMSALAuthProvider(radixConfig *radixconfig.RadixConfig, clientID, tenantID string) (MSALAuthProvider, error) {
	client, err := newPublicClient(radixConfig, clientID, tenantID)
	if err != nil {
		return nil, err
	}
	return &msalAuthProvider{
		client: client,
	}, nil
}

type msalAuthProvider struct {
	client *public.Client
}

func (provider *msalAuthProvider) AuthenticateRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	token, err := provider.GetToken(context.Background())
	if err != nil {
		return err
	}
	return r.SetHeaderParam(runtime.HeaderAuthorization, "Bearer "+token)
}

// Login allows the plugin to initialize its configuration. It must not
// require direct user interaction.
func (provider *msalAuthProvider) Login(ctx context.Context) error {
	_, err := provider.loginWithDeviceCode(ctx)
	return err
}

// Logout removes all cached accounts with tokens
func (provider *msalAuthProvider) Logout(ctx context.Context) error {
	accounts, err := provider.client.Accounts(ctx)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		if err := provider.client.RemoveAccount(ctx, account); err != nil {
			return err
		}
	}
	return nil
}

// GetToken returns a valid token for the given scopes
func (provider *msalAuthProvider) GetToken(ctx context.Context) (string, error) {
	accounts, err := provider.client.Accounts(ctx)
	if err != nil {
		return "", err
	}
	if len(accounts) > 0 {
		// found a cached account, now see if an applicable token has been cached
		// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
		//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
		authResult, err := provider.client.AcquireTokenSilent(ctx, getScopes(), public.WithSilentAccount(accounts[0]))
		if err == nil {
			return authResult.AccessToken, nil
		}
	}

	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	return provider.loginWithDeviceCode(ctx)
}

func (provider *msalAuthProvider) loginWithDeviceCode(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	devCode, err := provider.client.AcquireTokenByDeviceCode(ctx, getScopes())
	if err != nil {
		return "", fmt.Errorf("got error while waiting for user to input the device code: %s", err)
	}

	fmt.Println(devCode.Result.Message) // show authentication link with device code

	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func getScopes() []string {
	return []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
}
