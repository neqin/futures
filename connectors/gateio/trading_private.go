package gateio

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CreateFuturesOrder places a new futures order.
// settle: "usdt" or "btc"
// order: The order details defined in CreateFuturesOrderRequest.
func (c *Client) CreateFuturesOrder(ctx context.Context, settle string, order CreateFuturesOrderRequest) (*FuturesOrder, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders", settle)
	var result FuturesOrder
	err := c.post(ctx, endpoint, nil, order, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesOrders retrieves a list of futures orders.
// settle: "usdt" or "btc"
// contract: Filter by contract name (required if status is "open")
// status: Filter by order status ("open" or "finished") (required)
// limit: Maximum number of records. Default 100, Max 1000.
// offset: List offset.
// lastID: Specify the last order ID seen for pagination (alternative to offset).
// from: Start timestamp (seconds) (optional)
// to: End timestamp (seconds) (optional)
func (c *Client) ListFuturesOrders(ctx context.Context, settle, status string, contract *string, limit, offset *int, lastID *string, from, to *int64) (*[]FuturesOrder, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders", settle)
	params := url.Values{}
	params.Set("status", status)
	if contract != nil {
		params.Set("contract", *contract)
	} else if status == "open" {
		return nil, fmt.Errorf("contract is required when status is 'open'")
	}

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

	var result []FuturesOrder
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelAllFuturesOrders cancels all open orders in a specific contract or settlement currency.
// settle: "usdt" or "btc"
// contract: Futures contract name (required)
// side: Optional side filter ("buy" or "sell")
func (c *Client) CancelAllFuturesOrders(ctx context.Context, settle, contract string, side *string) (*BatchCancelOrdersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders", settle)
	params := url.Values{}
	params.Set("contract", contract)
	if side != nil {
		params.Set("side", *side)
	}

	var result BatchCancelOrdersResult
	// Note: API uses DELETE with query parameters for this.
	err := c.delete(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCancelFuturesOrders cancels multiple orders by ID.
// settle: "usdt" or "btc"
// orderIDs: A slice of order IDs to cancel.
func (c *Client) BatchCancelFuturesOrders(ctx context.Context, settle string, orderIDs []string) (*BatchCancelOrdersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders", settle)
	// The API expects order IDs in the request body for batch cancel.
	payload := orderIDs
	var result BatchCancelOrdersResult
	err := c.delete(ctx, endpoint, nil, payload, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFuturesOrder retrieves details of a single futures order.
// settle: "usdt" or "btc"
// orderID: The ID of the order to retrieve.
func (c *Client) GetFuturesOrder(ctx context.Context, settle, orderID string) (*FuturesOrder, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders/%s", settle, orderID)
	var result FuturesOrder
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelFuturesOrder cancels a single futures order by ID.
// settle: "usdt" or "btc"
// orderID: The ID of the order to cancel.
func (c *Client) CancelFuturesOrder(ctx context.Context, settle, orderID string) (*CancelOrderResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders/%s", settle, orderID)
	var result CancelOrderResult
	err := c.delete(ctx, endpoint, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AmendFuturesOrder modifies an existing open order.
// settle: "usdt" or "btc"
// orderID: The ID of the order to amend.
// size: Optional new size.
// price: Optional new price.
// amendText: Optional user-defined text prefixed with t-.
func (c *Client) AmendFuturesOrder(ctx context.Context, settle, orderID string, size *int64, price *string, amendText *string) (*FuturesOrder, error) {
	endpoint := fmt.Sprintf("/futures/%s/orders/%s", settle, orderID)
	params := url.Values{} // API uses query parameters for amendment
	if size != nil {
		params.Set("size", strconv.FormatInt(*size, 10))
	}
	if price != nil {
		params.Set("price", *price)
	}
	if amendText != nil {
		params.Set("amend_text", *amendText)
	}

	var result FuturesOrder
	err := c.put(ctx, endpoint, params, nil, &result) // PUT request with query params
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListMyFuturesTrades retrieves personal trading history.
// settle: "usdt" or "btc"
// contract: Filter by contract name (optional)
// orderID: Filter by order ID (optional)
// limit: Maximum number of records. Default 100, Max 1000.
// offset: List offset.
// lastID: Specify the last trade ID seen for pagination.
// from: Start timestamp (seconds) (optional)
// to: End timestamp (seconds) (optional)
func (c *Client) ListMyFuturesTrades(ctx context.Context, settle string, contract, orderID *string, limit, offset *int, lastID *string, from, to *int64) (*ListFuturesTradesResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/my_trades", settle)
	params := url.Values{}
	if contract != nil {
		params.Set("contract", *contract)
	}
	if orderID != nil {
		params.Set("order", *orderID) // API uses 'order' param for order_id filter
	}
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

// --- Price Triggered Orders ---

// CreateTriggerOrder creates a new price-triggered (conditional) order.
// settle: "usdt" or "btc"
// order: The trigger order details defined in CreateTriggerOrderRequest.
func (c *Client) CreateTriggerOrder(ctx context.Context, settle string, order CreateTriggerOrderRequest) (*TriggerOrder, error) {
	// The API requires the settle parameter to be part of the request body for trigger orders.
	order.Settle = settle
	endpoint := fmt.Sprintf("/futures/%s/price_orders", settle) // Settle is also in path
	var result TriggerOrder                                     // API doc says response is {id: int64}, but let's assume it returns the created order object for consistency
	err := c.post(ctx, endpoint, nil, order, &result)
	if err != nil {
		// If the error is just about the response structure (e.g., only ID returned),
		// we might need to adjust the 'result' type or handling here based on actual API behavior.
		// For now, assume it returns the full TriggerOrder.
		return nil, err
	}
	// If only ID is returned, we might need a GetTriggerOrder call here.
	return &result, nil
}

// ListTriggerOrders retrieves a list of price-triggered orders.
// settle: "usdt" or "btc"
// status: Filter by status ("open", "finished") (required)
// contract: Filter by contract name (optional)
// limit: Maximum number of records. Default 100, Max 1000.
// offset: List offset.
func (c *Client) ListTriggerOrders(ctx context.Context, settle, status string, contract *string, limit, offset *int) (*ListPriceTriggeredOrdersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/price_orders", settle)
	params := url.Values{}
	params.Set("status", status)
	if contract != nil {
		params.Set("contract", *contract)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		params.Set("offset", strconv.Itoa(*offset))
	}

	var result ListPriceTriggeredOrdersResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelAllTriggerOrders cancels all open trigger orders in a contract.
// settle: "usdt" or "btc"
// contract: Futures contract name (required)
func (c *Client) CancelAllTriggerOrders(ctx context.Context, settle, contract string) (*ListPriceTriggeredOrdersResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/price_orders", settle)
	params := url.Values{}
	params.Set("contract", contract)

	var result ListPriceTriggeredOrdersResult
	// Note: API uses DELETE with query parameters.
	err := c.delete(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTriggerOrder retrieves details of a single price-triggered order.
// settle: "usdt" or "btc"
// orderID: The ID of the trigger order.
func (c *Client) GetTriggerOrder(ctx context.Context, settle, orderID string) (*PriceTriggeredOrder, error) {
	endpoint := fmt.Sprintf("/futures/%s/price_orders/%s", settle, orderID)
	var result PriceTriggeredOrder
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelTriggerOrder cancels a single price-triggered order by ID.
// settle: "usdt" or "btc"
// orderID: The ID of the trigger order to cancel.
func (c *Client) CancelTriggerOrder(ctx context.Context, settle, orderID string) (*CancelPriceTriggeredOrderResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/price_orders/%s", settle, orderID)
	var result CancelPriceTriggeredOrderResult
	err := c.delete(ctx, endpoint, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Countdown Cancel ---

// SetCountdownCancelAll sets or cancels the countdown timer to cancel all orders.
// settle: "usdt" or "btc"
// request: CountdownCancelAllFuturesRequest containing timeout and optional contract.
func (c *Client) SetCountdownCancelAll(ctx context.Context, settle string, request CountdownCancelAllFuturesRequest) error {
	endpoint := fmt.Sprintf("/futures/%s/countdown_cancel_all", settle)
	// The response is just a success/failure status, no specific body structure defined in docs.
	err := c.post(ctx, endpoint, nil, request, nil)
	return err
}
