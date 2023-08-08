package oauth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"k8s.io/client-go/rest"
)

func NewMsalAuthProviderPlugin() rest.Factory {
	return func(name string, config map[string]string, persister rest.AuthProviderConfigPersister) (rest.AuthProvider, error) {
		return NewMsalAuthProvider(name, config, persister)
	}
}

func NewMsalAuthProvider(name string, config map[string]string, persister rest.AuthProviderConfigPersister) (rest.AuthProvider, error) {
	return &malAuthProvider{
		name:      name,
		client:    http.DefaultClient,
		cfg:       radixconfig.ToConfig(config),
		persister: persister,
	}, nil
}

type malAuthProvider struct {
	client    *http.Client
	cfg       radixconfig.RadixConfig
	persister rest.AuthProviderConfigPersister
	name      string
}

func (p *malAuthProvider) WrapTransport(rt http.RoundTripper) http.RoundTripper {
	return &roundTripper{
		wrapped:  rt,
		provider: p,
	}
}

func (p *malAuthProvider) Login() error {
	return errors.New("not yet implemented")
}

func (p *malAuthProvider) GetToken() (string, error) {
	cacheAccessor := &TokenCache{file: "/Users/SSMOL/.radix/config2"}
	cache := public.WithCache(cacheAccessor)
	app, err := public.New(p.cfg.SessionConfig.ClientID, cache, public.WithAuthority(getAuthority(p.cfg.SessionConfig)))
	if err != nil {
		return "", err
	}

	// found a cached account, now see if an applicable token has been cached
	// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
	//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
	authResult, err := app.AcquireTokenSilent(context.Background(), getScopes())
	if err == nil {
		return authResult.AccessToken, nil
	}
	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	devCode, err := app.AcquireTokenByDeviceCode(ctx, getScopes())
	if err != nil {
		return "", err
	}
	fmt.Printf("Device Code is: %s\n", devCode.Result.Message)
	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		return "", fmt.Errorf("got error while waiting for user to input the device code: %s", err)
	}
	return result.AccessToken, nil
}

func getScopes() []string {
	return []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
}

func getAuthority(config *radixconfig.SessionConfig) string {
	return fmt.Sprintf("https://login.microsoftonline.com/%s", config.TenantID)
}
