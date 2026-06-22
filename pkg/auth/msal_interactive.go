package auth

import (
	"context"
	"fmt"
	"time"

	msalcache "github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

type MsalInteractive struct {
	cache     msalcache.ExportReplace
	authority string
}

var _ GetAccessTokener = &MsalInteractive{}

func NewMsalInteractive(cache msalcache.ExportReplace, authority string) *MsalInteractive {
	return &MsalInteractive{
		cache:     cache,
		authority: authority,
	}
}

func (p *MsalInteractive) Authenticate(ctx context.Context, scopes []string) (AccessToken, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	client, err := p.getClient()
	if err != nil {
		return AccessToken{}, err
	}

	fmt.Printf("A web browser has been opened at %s/oauth2/v2.0/authorize. Please continue the login in the web browser.\n", p.authority)
	authResult, err := client.AcquireTokenInteractive(ctx, scopes)
	if err != nil {
		return AccessToken{}, err
	}

	return AccessToken{Token: authResult.AccessToken, ExpiresOn: authResult.ExpiresOn}, nil
}

func (p *MsalInteractive) GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error) {
	client, err := p.getClient()
	if err != nil {
		return AccessToken{}, err
	}

	accounts, err := client.Accounts(ctx)
	if err != nil {
		return AccessToken{}, err
	}

	if len(accounts) > 0 {
		// found a cached account, now see if an applicable token has been cached
		// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
		//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
		authResult, err := client.AcquireTokenSilent(ctx, scopes, public.WithSilentAccount(accounts[0]))
		if err == nil {
			return AccessToken{Token: authResult.AccessToken, ExpiresOn: authResult.ExpiresOn}, nil
		}
	}

	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	return p.Authenticate(ctx, scopes)
}

func (p *MsalInteractive) getClient() (public.Client, error) {
	return public.New(radixCliClientID, public.WithCache(p.cache), public.WithAuthority(p.authority))
}
