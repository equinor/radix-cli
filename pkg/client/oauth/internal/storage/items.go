// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package storage

import (
	"reflect"
	"strings"
	"time"
)

const (
	// CacheKeySeparator is used in creating the keys of the cache.
	CacheKeySeparator = "-"
)

// Contract is the JSON structure that is written to any storage medium when serializing
// the internal cache. This design is shared between MSAL versions in many languages.
// This cannot be changed without design that includes other SDKs.
type Contract struct {
	AccessTokens  map[string]AccessToken  `json:"AccessToken,omitempty"`
	RefreshTokens map[string]RefreshToken `json:"RefreshToken,omitempty"`
	IDTokens      map[string]IDToken      `json:"IdToken,omitempty"`
	Accounts      map[string]Account      `json:"Account,omitempty"`
	AppMetaData   map[string]AppMetaData  `json:"AppMetadata,omitempty"`

	AdditionalFields map[string]interface{}
}

// NewContract is the constructor for Contract.
func NewContract() *Contract {
	return &Contract{
		AccessTokens:     map[string]AccessToken{},
		RefreshTokens:    map[string]RefreshToken{},
		IDTokens:         map[string]IDToken{},
		Accounts:         map[string]Account{},
		AppMetaData:      map[string]AppMetaData{},
		AdditionalFields: map[string]interface{}{},
	}
}

type Account struct {
	HomeAccountID     string `json:"home_account_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	Realm             string `json:"realm,omitempty"`
	LocalAccountID    string `json:"local_account_id,omitempty"`
	AuthorityType     string `json:"authority_type,omitempty"`
	PreferredUsername string `json:"username,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	FamilyName        string `json:"family_name,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	Name              string `json:"name,omitempty"`
	AlternativeID     string `json:"alternative_account_id,omitempty"`
	RawClientInfo     string `json:"client_info,omitempty"`
	UserAssertionHash string `json:"user_assertion_hash,omitempty"`

	AdditionalFields map[string]interface{}
}

// RefreshToken is the JSON representation of a MSAL refresh token for encoding to storage.
type RefreshToken struct {
	HomeAccountID     string `json:"home_account_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	CredentialType    string `json:"credential_type,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
	FamilyID          string `json:"family_id,omitempty"`
	Secret            string `json:"secret,omitempty"`
	Realm             string `json:"realm,omitempty"`
	Target            string `json:"target,omitempty"`
	UserAssertionHash string `json:"user_assertion_hash,omitempty"`

	AdditionalFields map[string]interface{}
}

// AccessToken is the JSON representation of a MSAL access token for encoding to storage.
type AccessToken struct {
	HomeAccountID     string `json:"home_account_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	Realm             string `json:"realm,omitempty"`
	CredentialType    string `json:"credential_type,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
	Secret            string `json:"secret,omitempty"`
	Scopes            string `json:"target,omitempty"`
	ExpiresOn         string `json:"expires_on,omitempty"`
	ExtendedExpiresOn string `json:"extended_expires_on,omitempty"`
	CachedAt          string `json:"cached_at,omitempty"`
	UserAssertionHash string `json:"user_assertion_hash,omitempty"`

	AdditionalFields map[string]interface{}
}

// NewAccessToken is the constructor for AccessToken.
func NewAccessToken(homeID, env, realm, clientID string, cachedAt, expiresOn, extendedExpiresOn time.Time, scopes, token string) AccessToken {
	return AccessToken{
		HomeAccountID:     homeID,
		Environment:       env,
		Realm:             realm,
		CredentialType:    "AccessToken",
		ClientID:          clientID,
		Secret:            token,
		Scopes:            scopes,
		CachedAt:          string(cachedAt.UnixMilli()),
		ExpiresOn:         string(expiresOn.UnixMilli()),
		ExtendedExpiresOn: string(extendedExpiresOn.UnixMilli()),
	}
}

// Key outputs the key that can be used to uniquely look up this entry in a map.
func (a AccessToken) Key() string {
	key := strings.Join(
		[]string{a.HomeAccountID, a.Environment, a.CredentialType, a.ClientID, a.Realm, a.Scopes},
		CacheKeySeparator,
	)
	return strings.ToLower(key)
}

// FakeValidate enables tests to fake access token validation
var FakeValidate func(AccessToken) error

// IDToken is the JSON representation of an MSAL id token for encoding to storage.
type IDToken struct {
	HomeAccountID     string `json:"home_account_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	Realm             string `json:"realm,omitempty"`
	CredentialType    string `json:"credential_type,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
	Secret            string `json:"secret,omitempty"`
	UserAssertionHash string `json:"user_assertion_hash,omitempty"`
	AdditionalFields  map[string]interface{}
}

// IsZero determines if IDToken is the zero value.
func (i IDToken) IsZero() bool {
	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			switch field.Kind() {
			case reflect.Map, reflect.Slice:
				if field.Len() == 0 {
					continue
				}
			}
			return false
		}
	}
	return true
}

// NewIDToken is the constructor for IDToken.
func NewIDToken(homeID, env, realm, clientID, idToken string) IDToken {
	return IDToken{
		HomeAccountID:  homeID,
		Environment:    env,
		Realm:          realm,
		CredentialType: "IDToken",
		ClientID:       clientID,
		Secret:         idToken,
	}
}

// Key outputs the key that can be used to uniquely look up this entry in a map.
func (id IDToken) Key() string {
	key := strings.Join(
		[]string{id.HomeAccountID, id.Environment, id.CredentialType, id.ClientID, id.Realm},
		CacheKeySeparator,
	)
	return strings.ToLower(key)
}

// AppMetaData is the JSON representation of application metadata for encoding to storage.
type AppMetaData struct {
	FamilyID    string `json:"family_id,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	Environment string `json:"environment,omitempty"`

	AdditionalFields map[string]interface{}
}

// NewAppMetaData is the constructor for AppMetaData.
func NewAppMetaData(familyID, clientID, environment string) AppMetaData {
	return AppMetaData{
		FamilyID:    familyID,
		ClientID:    clientID,
		Environment: environment,
	}
}

// Key outputs the key that can be used to uniquely look up this entry in a map.
func (a AppMetaData) Key() string {
	key := strings.Join(
		[]string{"AppMetaData", a.Environment, a.ClientID},
		CacheKeySeparator,
	)
	return strings.ToLower(key)
}
