package cmd

import "net/http"

type AuthHeaderTransport struct {
	Transport http.RoundTripper
	Value     string
}

func (t *AuthHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.Value)
	return t.Transport.RoundTrip(req)
}

func NewHTTPClientWithAuthHeader(value string) *http.Client {
	client := &http.Client{}
	client.Transport = &AuthHeaderTransport{
		Transport: http.DefaultTransport,
		Value:     value,
	}
	return client
}
