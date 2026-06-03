package auth

import (
	"context"
	"fmt"
	"time"

	msalcache "github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

type MsalDeviceCode struct {
	cache     msalcache.ExportReplace
	authority string
}

var _ GetAccessTokener = &MsalDeviceCode{}

func NewMsalDeviceCode(cache msalcache.ExportReplace, authority string) *MsalDeviceCode {
	return &MsalDeviceCode{
		cache:     cache,
		authority: authority,
	}
}

func (p *MsalDeviceCode) Authenticate(ctx context.Context, scopes []string) (AccessToken, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	client, err := p.getClient()
	if err != nil {
		return AccessToken{}, err
	}

	devCode, err := client.AcquireTokenByDeviceCode(ctx, scopes)
	if err != nil {
		return AccessToken{}, fmt.Errorf("got error while waiting for user to input the device code: %s", err)
	}

	fmt.Println(devCode.Result.Message) // show authentication link with device code

	authResult, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		return AccessToken{}, err
	}
	return AccessToken{Token: authResult.AccessToken, ExpiresOn: authResult.ExpiresOn}, nil
}

func (p *MsalDeviceCode) GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error) {
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

func (p *MsalDeviceCode) getClient() (public.Client, error) {
	return public.New(radixCliClientID, public.WithCache(p.cache), public.WithAuthority(p.authority))
}
