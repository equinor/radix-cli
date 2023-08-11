package auth

import (
	"fmt"
	"net/http"
)

type roundTripper struct {
	provider *msalAuthProvider
	wrapped  http.RoundTripper
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return r.wrapped.RoundTrip(req)
	}
	token, err := r.provider.GetToken(req.Context())
	if err != nil {
		return nil, err
	}

	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *req
	// deep copy of the Header, so we don't modify the original
	// request's Header (as per RoundTripper contract).
	r2.Header = make(http.Header)
	for k, s := range req.Header {
		r2.Header[k] = s
	}
	r2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return r.wrapped.RoundTrip(r2)
}

func (r *roundTripper) WrappedRoundTripper() http.RoundTripper { return r.wrapped }
