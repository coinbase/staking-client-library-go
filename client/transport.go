package client

import (
	"fmt"
	"net/http"

	"github.com/coinbase/staking-client-library-go/auth"
)

// Transport implements the http.RoundTripper interface for use by Staking HTTP clients.
type Transport struct {
	roundTripper  http.RoundTripper
	authenticator *auth.Authenticator
	serviceName   string
}

// NewTransport returns new transport based on the given inputs.
func NewTransport(
	roundTripper http.RoundTripper,
	serviceName string,
	apiKey *auth.APIKey,
) (*Transport, error) {
	return &Transport{
		roundTripper:  roundTripper,
		authenticator: auth.NewAuthenticator(apiKey),
		serviceName:   serviceName,
	}, nil
}

// RoundTrip implements the http.RoundTripper interface and wraps
// the base round tripper with logic to inject the API key auth-based HTTP headers
// into the request. Reference: https://pkg.go.dev/net/http#RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	jwt, err := t.authenticator.BuildJWT(
		t.serviceName,
		fmt.Sprintf("%s %s%s", req.Method, req.URL.Host, req.URL.Path),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))

	return t.roundTripper.RoundTrip(req)
}
