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
	Login(ctx context.Context, useDeviceCode bool) error
	Logout(ctx context.Context) error
	runtime.ClientAuthInfoWriter
}

// NewMSALAuthProvider creates a new MSALAuthProvider
func NewMSALAuthProvider(radixConfig *radixconfig.RadixConfig, clientID, tenantID string) (MSALAuthProvider, error) {
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", tenantID)
	client, err := newPublicClient(radixConfig, clientID, authority)
	if err != nil {
		return nil, err
	}
	return &msalAuthProvider{
		client:    client,
		authority: authority,
	}, nil
}

type msalAuthProvider struct {
	authority string
	client    *public.Client
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
func (provider *msalAuthProvider) Login(ctx context.Context, useDeviceCode bool) error {
	var loginCmd func(context.Context) (string, error) = provider.loginInteractive
	if useDeviceCode {
		loginCmd = provider.loginDeviceCode
	}
	_, err := loginCmd(ctx)
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
	return provider.loginInteractive(ctx)
}

func (provider *msalAuthProvider) loginInteractive(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	fmt.Printf("A web browser has been opened at %s/oauth2/v2.0/authorize. Please continue the login in the web browser.\n", provider.authority)
	result, err := provider.client.AcquireTokenInteractive(ctx, getScopes())
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (provider *msalAuthProvider) loginDeviceCode(ctx context.Context) (string, error) {
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
