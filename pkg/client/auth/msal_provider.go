package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

// MSALAuthProvider is an AuthProvider that uses MSAL
type MSALAuthProvider interface {
	Login() error
	Logout() error
	WrapTransport(rt http.RoundTripper) http.RoundTripper
}

// NewMSALAuthProvider creates a new MSALAuthProvider
func NewMSALAuthProvider(radixConfig *radixconfig.RadixConfig) (MSALAuthProvider, error) {
	return &msalAuthProvider{
		client:      http.DefaultClient,
		radixConfig: radixConfig,
	}, nil
}

type msalAuthProvider struct {
	client      *http.Client
	radixConfig *radixconfig.RadixConfig
}

// WrapTransport allows the plugin to create a modified RoundTripper that
// attaches authorization headers (or other info) to requests.
func (provider *msalAuthProvider) WrapTransport(rt http.RoundTripper) http.RoundTripper {
	return &roundTripper{
		wrapped:  rt,
		provider: provider,
	}
}

// Login allows the plugin to initialize its configuration. It must not
// require direct user interaction.
func (provider *msalAuthProvider) Login() error {
	// login not supported for this AuthProvider
	return nil
}

// Logout removes all cached accounts with tokens
func (provider *msalAuthProvider) Logout() error {
	client, err := New(provider.radixConfig)
	if err != nil {
		return err
	}
	ctx := context.Background()
	accounts, err := client.Accounts(ctx)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		if err := client.RemoveAccount(ctx, account); err != nil {
			return err
		}
	}
	return nil
}

// GetToken returns a valid token for the given scopes
func (provider *msalAuthProvider) GetToken(ctx context.Context) (string, error) {
	client, err := New(provider.radixConfig)
	if err != nil {
		return "", err
	}

	accounts, err := client.Accounts(ctx)
	if err != nil {
		return "", err
	}
	if len(accounts) > 0 {
		// found a cached account, now see if an applicable token has been cached
		// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
		//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
		authResult, err := client.AcquireTokenSilent(ctx, getScopes(), public.WithSilentAccount(accounts[0]))
		if err == nil {
			return authResult.AccessToken, nil
		}
	}

	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	return provider.loginWithDeviceCode(ctx, client)
}

func (provider *msalAuthProvider) loginWithDeviceCode(ctx context.Context, client *public.Client) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	devCode, err := client.AcquireTokenByDeviceCode(ctx, getScopes())
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
