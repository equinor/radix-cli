package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	cache2 "github.com/equinor/radix-cli/pkg/cache"
)

type AzureGithub struct {
	authority     string
	azureClientId string
	cache         cache2.Cache
}

func NewAzureGithub(cache cache2.Cache, authority string) *AzureGithub {

	azureClientId, _ := cache.GetItem(azureClientIdCacheKey)

	return &AzureGithub{
		authority:     authority,
		cache:         cache,
		azureClientId: azureClientId,
	}
}

func (p *AzureGithub) Authenticate(ctx context.Context, azureClientId string) (string, error) {
	// Mostly copied from kubelogin\pkg\internal\token\githubactionscredential.go:newGithubActionsCredential()
	if token, ok := p.cache.GetItem(AccessTokenCacheKey); ok {
		return token, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	cred := confidential.NewCredFromAssertionCallback(func(ctx context.Context, _ confidential.AssertionRequestOptions) (string, error) {
		return getGithubFedCred(ctx)
	})

	client, err := confidential.New(p.authority, azureClientId, cred)
	if err != nil {
		return "", fmt.Errorf("failed to create github actions credential: %w", err)
	}

	authResult, err := client.AcquireTokenByCredential(ctx, getScopes())
	if err != nil {
		return "", err
	}

	p.cache.SetItem(azureClientIdCacheKey, azureClientId, authResult.ExpiresOn.Sub(time.Now()))
	p.cache.SetItem(AccessTokenCacheKey, authResult.AccessToken, authResult.ExpiresOn.Sub(time.Now()))
	return authResult.AccessToken, nil
}

func (p *AzureGithub) GetAccessToken(ctx context.Context) (string, error) {
	if token, ok := p.cache.GetItem(AccessTokenCacheKey); ok {
		return token, nil
	}

	azureClientId, _ := p.cache.GetItem(azureClientIdCacheKey)
	return p.Authenticate(ctx, azureClientId)
}

func getGithubFedCred(ctx context.Context) (string, error) {
	// All code copied from kubelogins kubelogin\pkg\internal\token\githubactionscredential.go:getGithubToken()
	reqToken := os.Getenv(actionsIDTokenRequestToken)
	reqURL := os.Getenv(actionsIDTokenRequestURL)

	if reqToken == "" || reqURL == "" {
		return "", errors.New("ACTIONS_ID_TOKEN_REQUEST_TOKEN or ACTIONS_ID_TOKEN_REQUEST_URL is not set")
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse ACTIONS_ID_TOKEN_REQUEST_URL: %w", err)
	}
	q := u.Query()
	q.Set("audience", azureADAudience)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	// reference:
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", reqToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; api-version=2.0")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body string
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			body = err.Error()
		} else {
			body = string(b)
		}

		return "", fmt.Errorf("github actions ID token request failed with status code: %d, response body: %s", resp.StatusCode, body)
	}

	var tokenResp githubTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.Value == "" {
		return "", errors.New("github actions ID token is empty")
	}

	return tokenResp.Value, nil
}
