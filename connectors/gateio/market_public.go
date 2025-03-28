package gateio

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ListFuturesContracts retrieves list of futures contracts.
// settle: "usdt" or "btc"
func (c *Client) ListFuturesContracts(ctx context.Context, settle string) (*TickerResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/contracts", settle)
	var result TickerResult
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListContractStats retrieves stats of a futures contract.
// settle: "usdt" or "btc"
// contract: Futures contract name
// interval: Candlestick interval. Default is 5m. Allowed: 5m, 15m, 30m, 1h, 4h, 1d
// limit: Maximum number of records to be returned. Default is 30, max 100
// startTime: Start timestamp of the query (seconds)
// endTime: End timestamp of the query (seconds)
func (c *Client) ListContractStats(ctx context.Context, settle, contract string, interval *string, limit *int, startTime, endTime *int64) (*ListContractStatsResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/contract_stats", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if interval != nil {
		params.Set("interval", *interval)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if startTime != nil {
		params.Set("from", strconv.FormatInt(*startTime, 10))
	}
	if endTime != nil {
		params.Set("to", strconv.FormatInt(*endTime, 10))
	}

	var result ListContractStatsResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesOrderBook retrieves futures order book.
// settle: "usdt" or "btc"
// contract: Futures contract name
// interval: Order book depth aggregation interval. '0' means no aggregation. Allowed values: 0, 0.1, 0.01, 0.001
// limit: Maximum number of order depth data in asks or bids. Default is 10, max 100
// withID: Whether the order book ID will be returned. Default is false
func (c *Client) ListFuturesOrderBook(ctx context.Context, settle, contract string, interval *string, limit *int, withID *bool) (*FutureOrderBook, error) {
	endpoint := fmt.Sprintf("/futures/%s/order_book", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if interval != nil {
		params.Set("interval", *interval)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if withID != nil {
		params.Set("with_id", strconv.FormatBool(*withID))
	}

	var orderBook FutureOrderBook
	err := c.get(ctx, endpoint, params, &orderBook)
	if err != nil {
		return nil, err
	}
	orderBook.Contract = contract // Add contract name to the result for context
	return &orderBook, nil
}

// ListFuturesTrades retrieves futures trading history.
// settle: "usdt" or "btc"
// contract: Futures contract name
// limit: Maximum number of records to be returned. Default is 100, max 1000
// offset: List offset, starting from 0
// lastID: Specify the starting point for this list based on the last retrieved ID
// from: Start timestamp of the query (seconds)
// to: End timestamp of the query (seconds)
func (c *Client) ListFuturesTrades(ctx context.Context, settle, contract string, limit, offset *int, lastID *string, from, to *int64) (*ListFuturesTradesResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/trades", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		params.Set("offset", strconv.Itoa(*offset))
	}
	if lastID != nil {
		params.Set("last_id", *lastID)
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}

	var result ListFuturesTradesResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesCandlesticks retrieves futures candlestick data.
// settle: "usdt" or "btc"
// contract: Futures contract name
// limit: Maximum number of records to be returned. Default is 100, max 1000
// interval: Interval time between candlesticks. Allowed values: 10s, 30s, 1m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 7d, 30d
// from: Start timestamp of the query (seconds)
// to: End timestamp of the query (seconds)
func (c *Client) ListFuturesCandlesticks(ctx context.Context, settle, contract string, limit *int, interval *string, from, to *int64) (*ListFuturesCandlesticksResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/candlesticks", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if interval != nil {
		params.Set("interval", *interval)
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}

	var result ListFuturesCandlesticksResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesPremiumIndex retrieves premium index K-line data.
// settle: "usdt" or "btc"
// contract: Futures contract name
// limit: Maximum number of records to be returned. Default is 100, max 1000
// interval: Interval time between candlesticks. Allowed values: 1m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 7d, 30d
// from: Start timestamp of the query (seconds)
// to: End timestamp of the query (seconds)
func (c *Client) ListFuturesPremiumIndex(ctx context.Context, settle, contract string, limit *int, interval *string, from, to *int64) (*ListFuturesPremiumIndexResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/premium_index", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if interval != nil {
		params.Set("interval", *interval)
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}

	var result ListFuturesPremiumIndexResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesTickers retrieves list of futures tickers.
// settle: "usdt" or "btc"
// contract: Futures contract name (optional)
func (c *Client) ListFuturesTickers(ctx context.Context, settle string, contract *string) (*ListFuturesTickersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/tickers", settle)
	params := url.Values{}
	if contract != nil {
		params.Set("contract", *contract)
	}

	var result ListFuturesTickersResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesFundingRateHistory retrieves funding rate history.
// settle: "usdt" or "btc"
// contract: Futures contract name
// limit: Maximum number of records to be returned. Default is 100, max 1000
func (c *Client) ListFuturesFundingRateHistory(ctx context.Context, settle, contract string, limit *int) (*ListFuturesFundingRateHistoryResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/funding_rate", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}

	var result ListFuturesFundingRateHistoryResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesInsuranceLedger retrieves insurance balance history.
// settle: "usdt" or "btc"
// limit: Maximum number of records to be returned. Default is 100, max 1000
func (c *Client) ListFuturesInsuranceLedger(ctx context.Context, settle string, limit *int) (*ListFuturesInsuranceLedgerResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/insurance", settle)
	params := url.Values{}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}

	var result ListFuturesInsuranceLedgerResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLiquidationHistory retrieves liquidation history. (Public Endpoint)
// settle: "usdt" or "btc"
// contract: Futures contract name (optional)
// limit: Maximum number of records to be returned. Default is 100, max 1000
// at: Specify the starting point for this list based on the liquidation time (seconds, timestamp)
// from: Start timestamp of the query (seconds)
// to: End timestamp of the query (seconds)
func (c *Client) GetLiquidationHistory(ctx context.Context, settle string, contract *string, limit *int, at, from, to *int64) (*GetLiquidationHistoryResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/liq_orders", settle)
	params := url.Values{}
	if contract != nil {
		params.Set("contract", *contract)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if at != nil {
		// Gate API uses seconds for 'at' timestamp
		params.Set("at", strconv.FormatInt(*at, 10))
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}

	var result GetLiquidationHistoryResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRiskLimitTiers retrieves risk limit tiers. (Public Endpoint)
// settle: "usdt" or "btc"
// contract: Futures contract name
func (c *Client) GetRiskLimitTiers(ctx context.Context, settle, contract string) (*GetRiskLimitTiersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/risk_limit_tiers", settle)
	params := url.Values{}
	params.Set("contract", contract)

	var result GetRiskLimitTiersResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
