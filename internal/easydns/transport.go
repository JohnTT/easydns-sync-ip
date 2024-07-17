package easydns

import (
	"encoding/base64"
	"net/http"
)

// BasicAuthTransport is a custom RoundTripper that adds basic auth credentials to each request
type BasicAuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

// RoundTrip executes a single HTTP transaction and adds the basic auth header
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	auth := t.Username + ":" + t.Password
	authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", authHeaderValue)

	// Delegate the actual request to the default transport
	return t.Transport.RoundTrip(req)
}
