package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

type MsalDeviceCode struct {
	client *public.Client
}

var _ Provider = &MsalDeviceCode{}

func NewMsalDeviceCode(cache cache.ExportReplace, authority string) (*MsalDeviceCode, error) {

	client, err := public.New(RadixCliClientID, public.WithCache(cache), public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}

	return &MsalDeviceCode{
		client: &client,
	}, nil
}

func (p *MsalDeviceCode) Authenticate(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	devCode, err := p.client.AcquireTokenByDeviceCode(ctx, getScopes())
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

func (p *MsalDeviceCode) GetAccessToken(ctx context.Context) (string, error) {
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
