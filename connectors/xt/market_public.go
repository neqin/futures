package xt

import (
	"context"
	"fmt"
	"net/http"

	// "net/url" // No longer needed here
	"strconv"
)

// --- Public Market Endpoints ---

// GetServerTime fetches the current server time from the API.
// Endpoint: GET /future/market/v1/public/time
func (c *Client) GetServerTime(ctx context.Context) (*ServerTimeResult, error) {
	path := "/future/market/v1/public/time"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M for this generic endpoint
	var result ServerTimeResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetServerTime failed: %w", err)
	}
	return &result, nil
}

// GetClientIP fetches the client's current IP address as seen by the server.
// Endpoint: GET /future/public/client
func (c *Client) GetClientIP(ctx context.Context) (*ClientIPResult, error) {
	path := "/future/public/client"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result ClientIPResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetClientIP failed: %w", err)
	}
	return &result, nil
}

// GetCoinsInfo gets the currency information of the transaction pairs (usually just ["usdt"]).
// Endpoint: GET /future/market/v1/public/symbol/coins
func (c *Client) GetCoinsInfo(ctx context.Context) (*CoinsInfoResult, error) {
	path := "/future/market/v1/public/symbol/coins"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result CoinsInfoResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetCoinsInfo failed: %w", err)
	}
	return &result, nil
}

// GetMarketConfig fetches configuration information for a single transaction pair.
// Endpoint: GET /future/market/v1/public/symbol/detail
func (c *Client) GetMarketConfig(ctx context.Context, symbol string) (*SingleContractResult, error) {
	path := "/future/market/v1/public/symbol/detail"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{      // Changed to map[string]string
		"symbol": symbol,
	}
	var result SingleContractResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetMarketConfig for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetAllMarketConfigV3 fetches configuration details for all available USDT-M futures contracts (using v3 endpoint).
// Endpoint: GET /future/market/v3/public/symbol/list
func (c *Client) GetAllMarketConfigV3(ctx context.Context) (*ContractsResult, error) {
	path := "/future/market/v3/public/symbol/list" // Using v3 endpoint from docs
	baseURL := c.getBaseURL("USDT-M")
	var result ContractsResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetAllMarketConfigV3 failed: %w", err)
	}
	return &result, nil
}

// GetLeverageDetail queries a single transaction pair for leverage stratification.
// Endpoint: GET /future/market/v1/public/leverage/bracket/detail
func (c *Client) GetLeverageDetail(ctx context.Context, symbol string) (*LeverageDetailResult, error) {
	path := "/future/market/v1/public/leverage/bracket/detail"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result LeverageDetailResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetLeverageDetail for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetLeverageDetailList queries all trading pairs for leverage stratification.
// Endpoint: GET /future/market/v1/public/leverage/bracket/list
func (c *Client) GetLeverageDetailList(ctx context.Context) (*LeverageDetailListResult, error) {
	path := "/future/market/v1/public/leverage/bracket/list"
	baseURL := c.getBaseURL("USDT-M")
	var result LeverageDetailListResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetLeverageDetailList failed: %w", err)
	}
	return &result, nil
}

// GetMarketTicker fetches market information for a specified trading pair.
// Endpoint: GET /future/market/v1/public/q/ticker
func (c *Client) GetMarketTicker(ctx context.Context, symbol string) (*SingleTickerResult, error) {
	path := "/future/market/v1/public/q/ticker"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result SingleTickerResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetMarketTicker for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetMarketTickers fetches market information for all trading pairs.
// Endpoint: GET /future/market/v1/public/q/tickers
func (c *Client) GetMarketTickers(ctx context.Context) (*TickersResult, error) {
	path := "/future/market/v1/public/q/tickers"
	baseURL := c.getBaseURL("USDT-M")
	var result TickersResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetMarketTickers failed: %w", err)
	}
	return &result, nil
}

// GetMarketDeal fetches recent public trades (deals) for a specific symbol.
// Endpoint: GET /future/market/v1/public/q/deal
func (c *Client) GetMarketDeal(ctx context.Context, symbol string, num int) (*TradesResult, error) {
	path := "/future/market/v1/public/q/deal"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
		"num":    strconv.Itoa(num),
	}
	var result TradesResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetMarketDeal for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetDepth fetches the order book depth for a specific symbol.
// Endpoint: GET /future/market/v1/public/q/depth
func (c *Client) GetDepth(ctx context.Context, symbol string, level int) (*DepthResult, error) {
	path := "/future/market/v1/public/q/depth"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
		"level":  strconv.Itoa(level),
	}
	var result DepthResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetDepth for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetIndexPrice gets the index price for a single trading pair.
// Endpoint: GET /future/market/v1/public/q/symbol-index-price
func (c *Client) GetIndexPrice(ctx context.Context, symbol string) (*IndexPriceResult, error) {
	path := "/future/market/v1/public/q/symbol-index-price"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result IndexPriceResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetIndexPrice for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetAllIndexPrice gets the index price for all trading pairs.
// Endpoint: GET /future/market/v1/public/q/index-price
func (c *Client) GetAllIndexPrice(ctx context.Context) (*AllIndexPriceResult, error) {
	path := "/future/market/v1/public/q/index-price"
	baseURL := c.getBaseURL("USDT-M")
	var result AllIndexPriceResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetAllIndexPrice failed: %w", err)
	}
	return &result, nil
}

// GetMarketPrice fetches the current mark price for a single symbol.
// Endpoint: GET /future/market/v1/public/q/symbol-mark-price
func (c *Client) GetMarketPrice(ctx context.Context, symbol string) (*SingleMarkPriceResult, error) {
	path := "/future/market/v1/public/q/symbol-mark-price"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result SingleMarkPriceResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetMarketPrice for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetAllMarketPrice fetches the current mark price for all symbols.
// Endpoint: GET /future/market/v1/public/q/mark-price
func (c *Client) GetAllMarketPrice(ctx context.Context) (*MarkPriceResult, error) {
	path := "/future/market/v1/public/q/mark-price"
	baseURL := c.getBaseURL("USDT-M")
	var result MarkPriceResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetAllMarketPrice failed: %w", err)
	}
	return &result, nil
}

// GetKlines fetches candlestick/k-line data for a specific symbol.
// Endpoint: GET /future/market/v1/public/q/kline
func (c *Client) GetKlines(ctx context.Context, symbol, interval string, startTime, endTime *int64, limit *int) (*KlinesResult, error) {
	path := "/future/market/v1/public/q/kline"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol":   symbol,
		"interval": interval,
	}
	if startTime != nil {
		params["startTime"] = strconv.FormatInt(*startTime, 10)
	}
	if endTime != nil {
		params["endTime"] = strconv.FormatInt(*endTime, 10)
	}
	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}
	var result KlinesResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetKlines for %s (%s) failed: %w", symbol, interval, err)
	}
	return &result, nil
}

// GetAggTicker gets aggregate market information for a specified trade pair.
// Endpoint: GET /future/market/v1/public/q/agg-ticker
func (c *Client) GetAggTicker(ctx context.Context, symbol string) (*AggTickerResult, error) {
	path := "/future/market/v1/public/q/agg-ticker"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result AggTickerResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetAggTicker for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetAllAggTicker gets aggregate market information for all trading pairs.
// Endpoint: GET /future/market/v1/public/q/agg-tickers
func (c *Client) GetAllAggTicker(ctx context.Context) (*AllAggTickerResult, error) {
	path := "/future/market/v1/public/q/agg-tickers"
	baseURL := c.getBaseURL("USDT-M")
	var result AllAggTickerResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetAllAggTicker failed: %w", err)
	}
	return &result, nil
}

// GetFundRate fetches the current funding rate for a symbol.
// Endpoint: GET /future/market/v1/public/q/funding-rate
func (c *Client) GetFundRate(ctx context.Context, symbol string) (*FundingRateResult, error) {
	path := "/future/market/v1/public/q/funding-rate"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result FundingRateResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetFundRate for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetBookTicker gets ask/bid market information for a specific trading pair.
// Endpoint: GET /future/market/v1/public/q/ticker/book
func (c *Client) GetBookTicker(ctx context.Context, symbol string) (*BookTickerResult, error) {
	path := "/future/market/v1/public/q/ticker/book"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result BookTickerResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetBookTicker for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetFundRateRecord gets funding rate records.
// Endpoint: GET /future/market/v1/public/q/funding-rate-record
func (c *Client) GetFundRateRecord(ctx context.Context, symbol string, direction *string, id *int64, limit *int) (*FundRateRecordResult, error) {
	path := "/future/market/v1/public/q/funding-rate-record"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol, // Required
	}
	if direction != nil {
		params["direction"] = *direction
	}
	if id != nil {
		params["id"] = strconv.FormatInt(*id, 10)
	}
	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}
	var result FundRateRecordResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetFundRateRecord for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetAllBookTickers gets ask/bid market information for all trading pairs.
// Endpoint: GET /future/market/v1/public/q/ticker/books
func (c *Client) GetAllBookTickers(ctx context.Context) (*AllBookTickerResult, error) {
	path := "/future/market/v1/public/q/ticker/books"
	baseURL := c.getBaseURL("USDT-M")
	var result AllBookTickerResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, nil, &result) // Pass nil for params
	if err != nil {
		return nil, fmt.Errorf("GetAllBookTickers failed: %w", err)
	}
	return &result, nil
}

// GetRiskBalance obtains trading pairs of venture fund balances.
// Endpoint: GET /future/market/v1/public/contract/risk-balance
func (c *Client) GetRiskBalance(ctx context.Context, symbol string, direction *string, id *int64, limit *int) (*RiskBalanceResult, error) {
	path := "/future/market/v1/public/contract/risk-balance"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol, // Required
	}
	if direction != nil {
		params["direction"] = *direction
	}
	if id != nil {
		params["id"] = strconv.FormatInt(*id, 10)
	}
	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}
	var result RiskBalanceResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetRiskBalance for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetOpenInterest fetches the open interest for a specific symbol.
// Endpoint: GET /future/market/v1/public/contract/open-interest
func (c *Client) GetOpenInterest(ctx context.Context, symbol string) (*OpenInterestResult, error) {
	path := "/future/market/v1/public/contract/open-interest"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{ // Changed to map[string]string
		"symbol": symbol,
	}
	var result OpenInterestResult
	err := c.SendPublicRequest(ctx, http.MethodGet, baseURL, path, params, &result) // Pass map
	if err != nil {
		return nil, fmt.Errorf("GetOpenInterest for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// TODO: Implement CoinGecko compatible endpoints if needed
// /future/market/v1/public/cg/contracts
// /future/market/v1/public/cg/orderbook
