package auth

// import (
// 	"fmt"
// 	"net/http"
// )

// type roundTripper struct {
// 	provider *msalAuthProvider
// 	wrapped  http.RoundTripper
// }

// // RoundTrip executes a single HTTP transaction, returning
// // a Response for the provided Request.
// func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
// 	if len(req.Header.Get("Authorization")) != 0 {
// 		return r.wrapped.RoundTrip(req)
// 	}
// 	token, err := r.provider.GetToken(req.Context())
// 	if err != nil {
// 		return nil, err
// 	}

// 	// shallow copy of the struct
// 	copyReq := new(http.Request)
// 	*copyReq = *req
// 	// deep copy of the Header, so we don't modify the original
// 	// request's Header (as per RoundTripper contract).
// 	copyReq.Header = make(http.Header)
// 	for key, val := range req.Header {
// 		copyReq.Header[key] = val
// 	}
// 	copyReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

// 	return r.wrapped.RoundTrip(copyReq)
// }

// // WrappedRoundTripper returns the underlying RoundTripper
// func (r *roundTripper) WrappedRoundTripper() http.RoundTripper { return r.wrapped }
