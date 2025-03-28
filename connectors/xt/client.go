package xt

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	defaultUsdtBaseURL = "https://fapi.xt.com"
	defaultCoinBaseURL = "https://dapi.xt.com"
	defaultRecvWindow  = "5000" // Default 5 seconds validity window
)

// Client is the main XT.com Futures API client.
type Client struct {
	apiKey      string
	secretKey   string
	usdtBaseURL string
	coinBaseURL string
	httpClient  *http.Client
	recvWindow  string // Receive window in milliseconds as a string
}

// NewClient creates a new XT.com Futures API client.
func NewClient(apiKey, secretKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second} // Default timeout
	}
	return &Client{
		usdtBaseURL: defaultUsdtBaseURL,
		coinBaseURL: defaultCoinBaseURL,
		apiKey:      apiKey,
		secretKey:   secretKey,
		httpClient:  httpClient,
		recvWindow:  defaultRecvWindow,
	}
}

// SetUsdtBaseURL allows overriding the default USDT-M base URL.
func (c *Client) SetUsdtBaseURL(baseURL string) {
	c.usdtBaseURL = strings.TrimSuffix(baseURL, "/")
	log.Printf("XT Client USDT Base URL set to: %s", c.usdtBaseURL)
}

// SetCoinBaseURL allows overriding the default COIN-M base URL.
func (c *Client) SetCoinBaseURL(baseURL string) {
	c.coinBaseURL = strings.TrimSuffix(baseURL, "/")
	log.Printf("XT Client COIN Base URL set to: %s", c.coinBaseURL)
}

// SetRecvWindow sets the request validity window in milliseconds. Default is 5000 (5 seconds).
func (c *Client) SetRecvWindow(ms int64) {
	c.recvWindow = strconv.FormatInt(ms, 10)
}

// getBaseURL returns the appropriate base URL based on underlying type (e.g., USDT-M or COIN-M).
// For now, defaulting to USDT-M as most examples use it. This might need refinement
// if methods need to dynamically choose based on symbol or explicit parameter.
func (c *Client) getBaseURL(underlyingType string) string {
	// TODO: Add logic to select URL based on underlyingType if needed
	// if underlyingType == "COIN-M" { return c.coinBaseURL }
	return c.usdtBaseURL
}

// generateSignature creates the HMAC SHA256 signature based on XT documentation (xt2.txt).
func (c *Client) generateSignature(timestamp, path, sortedQuery, bodyString string) string {
	// X = Sorted header parameters
	headerPart := fmt.Sprintf("validate-appkey=%s&validate-timestamp=%s", c.apiKey, timestamp)
	// Optional: Add recvWindow if needed
	// headerPart += "&validate-recvwindow=" + c.recvWindow

	// Y = #path#query#body (adjust if query or body is empty)
	dataPart := "#" + path
	if sortedQuery != "" {
		dataPart += "#" + sortedQuery
	}
	if bodyString != "" {
		dataPart += "#" + bodyString
	}

	// sign = XY
	signStr := headerPart + dataPart

	// signature = HMAC-SHA256(secretKey, sign)
	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(signStr))
	signature := hex.EncodeToString(mac.Sum(nil))

	// log.Printf("XT Sig Base X: %s", headerPart) // Debugging
	// log.Printf("XT Sig Base Y: %s", dataPart) // Debugging
	// log.Printf("XT Sig sign=XY: %s", signStr) // Debugging
	// log.Printf("XT Signature: %s", signature) // Debugging

	return signature
}

// sortAndEncodeParams sorts map keys alphabetically and returns URL-encoded string "key=value&key=value..."
func sortAndEncodeParams(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for i, k := range keys {
		if i > 0 {
			builder.WriteString("&")
		}
		builder.WriteString(url.QueryEscape(k))
		builder.WriteString("=")
		builder.WriteString(url.QueryEscape(params[k]))
	}
	return builder.String()
}

// sendRequest handles sending HTTP requests (both public and private).
func (c *Client) sendRequest(ctx context.Context, method, baseURL, path string, queryParams map[string]string, bodyParams interface{}, isPrivate bool, target interface{}) error {

	// --- Prepare URL and Query String ---
	fullURL := baseURL + path
	sortedQueryString := sortAndEncodeParams(queryParams) // Sort query params for potential signature use and request URL
	if sortedQueryString != "" {
		fullURL += "?" + sortedQueryString
	}

	// --- Prepare Body ---
	var bodyReader io.Reader
	var bodyBytes []byte
	var bodyStringForSig string // String representation of body for signature
	var contentType string = "" // Default empty, set based on body type
	var err error

	if bodyParams != nil {
		// Check if it's form data (map[string]string)
		if formData, ok := bodyParams.(map[string]string); ok {
			bodyStringForSig = sortAndEncodeParams(formData) // Sort and encode form data for signature
			bodyReader = strings.NewReader(bodyStringForSig) // Use encoded string as request body
			contentType = "application/x-www-form-urlencoded"
		} else {
			// Assume JSON otherwise
			bodyBytes, err = json.Marshal(bodyParams)
			if err != nil {
				return fmt.Errorf("failed to marshal request body to JSON: %w", err)
			}
			bodyStringForSig = string(bodyBytes)
			bodyReader = bytes.NewReader(bodyBytes)
			contentType = "application/json"
		}
	}

	// --- Create Request ---
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// --- Set Headers ---
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json") // Assume JSON response generally

	// --- Add Authentication Headers (if private) ---
	if isPrivate {
		if c.apiKey == "" || c.secretKey == "" {
			return fmt.Errorf("API key and secret key must be provided for private endpoints")
		}
		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

		// Determine query string part for signature (only for GET/DELETE)
		sigQueryPart := ""
		if method == http.MethodGet || method == http.MethodDelete {
			sigQueryPart = sortedQueryString
		}

		signature := c.generateSignature(timestamp, path, sigQueryPart, bodyStringForSig)

		req.Header.Set("validate-appkey", c.apiKey)
		req.Header.Set("validate-timestamp", timestamp)
		req.Header.Set("validate-signature", signature)
		// Optional: Add recvWindow if needed
		// req.Header.Set("validate-recvwindow", c.recvWindow)
	}

	// log.Printf("[XT:%s] %s", method, fullURL) // Debugging request URL
	// log.Printf("XT Request Headers: %v", req.Header) // Debugging headers
	// if bodyStringForSig != "" {
	// 	log.Printf("XT Request Body String for Sig: %s", bodyStringForSig) // Debugging body
	// }

	// --- Send Request ---
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// --- Read Response Body ---
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// log.Printf("XT Response Status: %s", resp.Status) // Debugging response status
	// if len(responseBody) < 2000 { // Avoid logging huge responses
	// 	log.Printf("XT Response Body: %s", string(responseBody))
	// } else {
	// 	log.Printf("XT Response Body: (omitted, length %d)", len(responseBody))
	// }

	// --- Handle Errors and Unmarshal ---
	// Try unmarshalling into CommonResponse first to check returnCode
	var commonResp CommonResponse
	if err := json.Unmarshal(responseBody, &commonResp); err != nil {
		// If basic unmarshal fails, return generic error
		return fmt.Errorf("failed to unmarshal basic response structure: %w (body: %s)", err, string(responseBody))
	}

	if commonResp.ReturnCode != 0 {
		// Return structured API error
		return fmt.Errorf("XT API error: code=%d, msg=%s, error=%s", commonResp.ReturnCode, commonResp.MsgInfo, string(commonResp.Error))
	}

	// Unmarshal into the specific target struct if provided
	if target != nil {
		err = json.Unmarshal(responseBody, target)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response body into target: %w (body: %s)", err, string(responseBody))
		}
	}

	return nil
}

// --- Public and Private Request Helpers ---

// SendPublicRequest sends a request to a public endpoint.
func (c *Client) SendPublicRequest(ctx context.Context, method, baseURL, path string, params map[string]string, target interface{}) error {
	return c.sendRequest(ctx, method, baseURL, path, params, nil, false, target)
}

// SendPrivateRequest sends an authenticated request.
// queryParams are used for GET/DELETE.
// bodyParams are used for POST/PUT (can be map[string]string for form-urlencoded or struct/map for JSON).
func (c *Client) SendPrivateRequest(ctx context.Context, method, baseURL, path string, queryParams map[string]string, bodyParams interface{}, target interface{}) error {
	return c.sendRequest(ctx, method, baseURL, path, queryParams, bodyParams, true, target)
}
