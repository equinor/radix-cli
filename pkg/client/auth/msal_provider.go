package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"k8s.io/client-go/rest"
)

type MSALAuthProvider interface {
	Login() error
	Logout() error
	WrapTransport(rt http.RoundTripper) http.RoundTripper
}

func NewMSALAuthProvider(radixConfig *radixconfig.RadixConfig, persister rest.AuthProviderConfigPersister) (MSALAuthProvider, error) {
	return &msalAuthProvider{
		name:        "msal",
		client:      http.DefaultClient,
		radixConfig: radixConfig,
		persister:   persister,
	}, nil
}

type msalAuthProvider struct {
	client      *http.Client
	radixConfig *radixconfig.RadixConfig
	persister   rest.AuthProviderConfigPersister
	name        string
}

func (p *msalAuthProvider) WrapTransport(rt http.RoundTripper) http.RoundTripper {
	return &roundTripper{
		wrapped:  rt,
		provider: p,
	}
}

func (p *msalAuthProvider) Login() error {
	// login not supported for this AuthProvider
	return nil
}

func (p *msalAuthProvider) Logout() error {
	client, err := NewClient(p.radixConfig)
	if err != nil {
		return err
	}
	ctx := context.Background()
	for {
		account, err := getExistingAccount(ctx, client)
		if err != nil {
			return err
		}
		if account.IsZero() {
			break
		}
		err = client.RemoveAccount(ctx, account)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *msalAuthProvider) GetToken(ctx context.Context) (string, error) {
	client, err := NewClient(p.radixConfig)
	if err != nil {
		return "", err
	}
	account, err := getExistingAccount(ctx, client)
	if err != nil {
		return "", err
	}
	// found a cached account, now see if an applicable token has been cached
	// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
	//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
	authResult, err := client.AcquireTokenSilent(ctx, getScopes(), public.WithSilentAccount(account))
	if err == nil {
		return authResult.AccessToken, nil
	}
	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	result, err := p.loginWithDeviceCode(ctx, client)
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (p *msalAuthProvider) loginWithDeviceCode(ctx context.Context, client *public.Client) (*public.AuthResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	devCode, err := client.AcquireTokenByDeviceCode(ctx, getScopes())
	if err != nil {
		return nil, fmt.Errorf("got error while waiting for user to input the device code: %s", err)
	}
	fmt.Printf("Device Code is: %s\n", devCode.Result.Message)
	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func getScopes() []string {
	return []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
}
