package gateio

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://api.gateio.ws"
	apiPrefix      = "/api/v4"
)

// Client is the main Gate.io API client.
type Client struct {
	apiKey     string
	secretKey  string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Gate.io API client.
// Provide apiKey and secretKey for accessing private endpoints.
// If apiKey and secretKey are empty, only public endpoints can be accessed.
func NewClient(apiKey, secretKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second} // Default timeout
	}
	return &Client{
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		secretKey:  secretKey,
		httpClient: httpClient,
	}
}

// SetBaseURL allows overriding the default base URL (e.g., for testing environments).
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = strings.TrimSuffix(baseURL, "/")
	log.Printf("Gate Client Base URL set to: %s", c.baseURL)
}

// generateSignature creates the HMAC SHA512 signature for Gate.io API v4 private requests.
func (c *Client) generateSignature(method, path, query, body string, timestamp string) string {
	// Hash the body using SHA512
	bodyHash := sha512.New()
	bodyHash.Write([]byte(body))
	hashedPayload := hex.EncodeToString(bodyHash.Sum(nil))

	// Create the signature string
	// METHOD\nURL_PATH\nQUERY_STRING\nHASHED_REQUEST_PAYLOAD\nTIMESTAMP
	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, path, query, hashedPayload, timestamp)

	// Sign using HMAC-SHA512
	mac := hmac.New(sha512.New, []byte(c.secretKey))
	mac.Write([]byte(signStr))
	signature := hex.EncodeToString(mac.Sum(nil))

	// log.Printf("Gate Sign String: %s", signStr) // Debugging
	// log.Printf("Gate Hashed Payload: %s", hashedPayload) // Debugging
	// log.Printf("Gate Signature: %s", signature) // Debugging
	return signature
}

// sendRequest creates, signs (if private), and sends an HTTP request.
func (c *Client) sendRequest(ctx context.Context, method, endpointPath string, queryParams url.Values, bodyPayload interface{}, target interface{}) error {
	isPrivate := c.apiKey != "" && c.secretKey != ""

	// Prepare URL
	fullURL := c.baseURL + apiPrefix + endpointPath
	queryString := ""
	if queryParams != nil {
		queryString = queryParams.Encode()
	}
	if queryString != "" {
		fullURL += "?" + queryString
	}

	// Prepare Body
	var bodyReader io.Reader
	var bodyBytes []byte
	var err error

	if bodyPayload != nil {
		bodyBytes, err = json.Marshal(bodyPayload)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	} else {
		// Ensure bodyBytes is an empty slice, not nil, for hashing
		bodyBytes = []byte{}
	}

	// Create Request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set Headers
	req.Header.Set("Accept", "application/json")
	if method == "POST" || method == "PUT" || method == "DELETE" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add Authentication Headers if private
	if isPrivate {
		if c.apiKey == "" || c.secretKey == "" {
			return fmt.Errorf("API key and secret key must be provided for private endpoints")
		}
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		signature := c.generateSignature(method, apiPrefix+endpointPath, queryString, string(bodyBytes), timestamp)

		req.Header.Set("KEY", c.apiKey)
		req.Header.Set("Timestamp", timestamp)
		req.Header.Set("SIGN", signature)
	}

	// log.Printf("[GATE.IO:%s] %s", method, fullURL) // Debugging request URL

	// Send Request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read Response Body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// log.Printf("Gate Response Status: %s", resp.Status) // Debugging response status
	// if len(responseBody) < 1000 { // Avoid logging huge responses
	// 	log.Printf("Gate Response Body: %s", string(responseBody))
	// } else {
	// 	log.Printf("Gate Response Body: (omitted, length %d)", len(responseBody))
	// }

	// Handle Errors
	if resp.StatusCode >= 400 {
		var apiErr APIError
		err = json.Unmarshal(responseBody, &apiErr)
		if err == nil && apiErr.Label != "" {
			// Return the structured API error
			return apiErr
		}
		// Return a generic error if parsing fails or it's not the expected format
		return fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(responseBody))
	}

	// Unmarshal Success Response
	if target != nil {
		err = json.Unmarshal(responseBody, target)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response body into target: %w (body: %s)", err, string(responseBody))
		}
	}

	return nil
}

// --- Helper methods for different request types ---

func (c *Client) get(ctx context.Context, endpointPath string, params url.Values, target interface{}) error {
	return c.sendRequest(ctx, http.MethodGet, endpointPath, params, nil, target)
}

func (c *Client) post(ctx context.Context, endpointPath string, params url.Values, payload interface{}, target interface{}) error {
	return c.sendRequest(ctx, http.MethodPost, endpointPath, params, payload, target)
}

func (c *Client) delete(ctx context.Context, endpointPath string, params url.Values, payload interface{}, target interface{}) error {
	return c.sendRequest(ctx, http.MethodDelete, endpointPath, params, payload, target)
}

func (c *Client) put(ctx context.Context, endpointPath string, params url.Values, payload interface{}, target interface{}) error {
	return c.sendRequest(ctx, http.MethodPut, endpointPath, params, payload, target)
}
