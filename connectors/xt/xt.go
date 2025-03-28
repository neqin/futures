package xt

import (
	"net/http"
	"time"
)

// DefaultHTTPClient is a default HTTP client with a timeout.
var DefaultHTTPClient = &http.Client{Timeout: 10 * time.Second}

// New creates a new XT.com Futures API client instance.
// Provide apiKey and secretKey for accessing private endpoints.
// If httpClient is nil, a default client with a 10-second timeout will be used.
func New(apiKey, secretKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = DefaultHTTPClient
	}
	return NewClient(apiKey, secretKey, httpClient)
}

// NewPublicOnly creates a new XT.com Futures API client instance for accessing only public endpoints.
// If httpClient is nil, a default client with a 10-second timeout will be used.
func NewPublicOnly(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = DefaultHTTPClient
	}
	return NewClient("", "", httpClient)
}
