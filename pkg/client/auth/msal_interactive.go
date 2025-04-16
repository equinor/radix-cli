package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

type MsalInteractive struct {
	client    *public.Client
	authority string
}

var _ Provider = &MsalInteractive{}

func NewMsalInteractive(cache cache.ExportReplace, authority string) (*MsalInteractive, error) {

	client, err := public.New(RadixCliClientID, public.WithCache(cache), public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}

	return &MsalInteractive{
		client:    &client,
		authority: authority,
	}, nil
}

func (p *MsalInteractive) Authenticate(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	fmt.Printf("A web browser has been opened at %s/oauth2/v2.0/authorize. Please continue the login in the web browser.\n", p.authority)
	result, err := p.client.AcquireTokenInteractive(ctx, getScopes())
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (p *MsalInteractive) GetAccessToken(ctx context.Context) (string, error) {
	accounts, err := p.client.Accounts(ctx)

	if err != nil {
		return "", err
	}
	if len(accounts) > 0 {
		// found a cached account, now see if an applicable token has been cached
		// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
		//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
		authResult, err := p.client.AcquireTokenSilent(ctx, getScopes(), public.WithSilentAccount(accounts[0]))
		if err == nil {
			return authResult.AccessToken, nil
		}
	}

	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	return p.Authenticate(ctx)
}
