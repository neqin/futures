package xt

import (
	"context"
	"encoding/json" // Needed for batch order list marshaling
	"fmt"
	"net/http"
	"strconv"
)

// --- Private Trading Endpoints ---

// PlaceOrderRequest defines parameters for placing a new order.
// Note: Field names adjusted to match API docs (e.g., orderSide, orderType, origQty).
type PlaceOrderRequest struct {
	ClientOrderID      *string `json:"clientOrderId,omitempty"`      // Optional
	Symbol             string  `json:"symbol"`                       // Required
	OrderSide          string  `json:"orderSide"`                    // Required: BUY, SELL
	OrderType          string  `json:"orderType"`                    // Required: LIMIT, MARKET
	OrigQty            string  `json:"origQty"`                      // Required: Quantity (Cont)
	Price              *string `json:"price,omitempty"`              // Required for LIMIT orders
	TimeInForce        *string `json:"timeInForce,omitempty"`        // Optional: GTC, IOC, FOK, GTX
	TriggerProfitPrice *string `json:"triggerProfitPrice,omitempty"` // Optional: TP trigger price
	TriggerStopPrice   *string `json:"triggerStopPrice,omitempty"`   // Optional: SL trigger price
	PositionSide       string  `json:"positionSide"`                 // Required: LONG, SHORT
}

// PlaceOrder creates a new futures order.
// Endpoint: POST /future/trade/v1/order/create
func (c *Client) PlaceOrder(ctx context.Context, orderReq PlaceOrderRequest) (*PlaceOrderResult, error) {
	path := "/future/trade/v1/order/create"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M

	// Basic validation
	if orderReq.Symbol == "" || orderReq.OrderSide == "" || orderReq.OrderType == "" || orderReq.OrigQty == "" || orderReq.PositionSide == "" {
		return nil, fmt.Errorf("missing required fields in PlaceOrderRequest (symbol, orderSide, orderType, origQty, positionSide)")
	}
	if orderReq.OrderType == "LIMIT" && (orderReq.Price == nil || *orderReq.Price == "") {
		return nil, fmt.Errorf("price is required for LIMIT orders")
	}

	var result PlaceOrderResult
	// API accepts application/json or application/x-www-form-urlencoded
	// Let's use JSON as it's generally easier with structs.
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, orderReq, &result)
	if err != nil {
		return nil, fmt.Errorf("PlaceOrder for %s failed: %w", orderReq.Symbol, err)
	}
	return &result, nil
}

// PlaceBatchOrderRequest defines parameters for placing multiple orders.
// The 'List' field should contain a slice of PlaceOrderRequest structs.
type PlaceBatchOrderRequest struct {
	List []PlaceOrderRequest `json:"-"` // Use struct slice, will be marshaled to JSON string for the 'list' parameter
}

// PlaceBatchOrder places multiple orders at once.
// Endpoint: POST /future/trade/v2/order/create-batch
// Note: API expects the list of orders as a JSON *string* within the 'list' form parameter.
func (c *Client) PlaceBatchOrder(ctx context.Context, batchReq PlaceBatchOrderRequest) (*PlaceBatchOrderResult, error) {
	path := "/future/trade/v2/order/create-batch" // Using v2 endpoint from docs
	baseURL := c.getBaseURL("USDT-M")

	if len(batchReq.List) == 0 {
		return nil, fmt.Errorf("order list cannot be empty for batch order")
	}

	// Marshal the list of orders into a JSON string
	listJSON, err := json.Marshal(batchReq.List)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal order list for batch request: %w", err)
	}

	// Send as application/x-www-form-urlencoded with the 'list' parameter containing the JSON string
	bodyParams := map[string]string{
		"list": string(listJSON),
	}

	var result PlaceBatchOrderResult
	err = c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("PlaceBatchOrder failed: %w", err)
	}
	return &result, nil
}

// CancelOrder cancels a single futures order by ID.
// Endpoint: POST /future/trade/v1/order/cancel
func (c *Client) CancelOrder(ctx context.Context, orderID int64) (*CancelOrderResult, error) {
	path := "/future/trade/v1/order/cancel"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{ // Docs indicate x-www-form-urlencoded or JSON
		"orderId": strconv.FormatInt(orderID, 10),
	}
	var result CancelOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelOrder for ID %d failed: %w", orderID, err)
	}
	return &result, nil
}

// CancelBatchOrder cancels all orders, optionally filtered by symbol.
// Endpoint: POST /future/trade/v1/order/cancel-all
func (c *Client) CancelBatchOrder(ctx context.Context, symbol *string) (*CancelBatchOrderResult, error) {
	path := "/future/trade/v1/order/cancel-all"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{} // Use map for optional param
	if symbol != nil {
		bodyParams["symbol"] = *symbol // API expects empty string to cancel all
	} else {
		bodyParams["symbol"] = "" // Explicitly send empty string if nil
	}

	var result CancelBatchOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		symbolStr := "all symbols"
		if symbol != nil {
			symbolStr = *symbol
		}
		return nil, fmt.Errorf("CancelBatchOrder for %s failed: %w", symbolStr, err)
	}
	return &result, nil
}

// GetOrder queries the details of a specific order by ID.
// Endpoint: GET /future/trade/v1/order/detail
func (c *Client) GetOrder(ctx context.Context, orderID int64) (*GetOrderResult, error) {
	path := "/future/trade/v1/order/detail"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"orderId": strconv.FormatInt(orderID, 10),
	}
	var result GetOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrder for ID %d failed: %w", orderID, err)
	}
	return &result, nil
}

// GetOrderListRequest defines parameters for querying orders.
type GetOrderListRequest struct {
	State         *string `url:"state,omitempty"`         // Optional: NEW, PARTIALLY_FILLED, FILLED, CANCELED, etc. (Use HISTORY for history endpoint)
	Symbol        *string `url:"symbol,omitempty"`        // Optional: filter by symbol
	ClientOrderID *string `url:"clientOrderId,omitempty"` // Optional: filter by client order ID
	Page          *int    `url:"page,omitempty"`          // Optional: pagination (default 1)
	Size          *int    `url:"size,omitempty"`          // Optional: pagination (default 10)
	StartTime     *int64  `url:"startTime,omitempty"`     // Optional: filter by time (ms)
	EndTime       *int64  `url:"endTime,omitempty"`       // Optional: filter by time (ms)
}

// GetOrderList queries orders based on state and other filters.
// Endpoint: GET /future/trade/v1/order/list
func (c *Client) GetOrderList(ctx context.Context, queryReq GetOrderListRequest) (*GetOrderListResult, error) {
	path := "/future/trade/v1/order/list"
	baseURL := c.getBaseURL("USDT-M")
	params := make(map[string]string)
	if queryReq.State != nil {
		params["state"] = *queryReq.State
	}
	if queryReq.Symbol != nil {
		params["symbol"] = *queryReq.Symbol
	}
	if queryReq.ClientOrderID != nil {
		params["clientOrderId"] = *queryReq.ClientOrderID
	}
	if queryReq.Page != nil {
		params["page"] = strconv.Itoa(*queryReq.Page)
	}
	if queryReq.Size != nil {
		params["size"] = strconv.Itoa(*queryReq.Size)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetOrderListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetOrderList failed: %w", err)
	}
	return &result, nil
}

// GetHistoryListRequest defines parameters for querying order history.
type GetHistoryListRequest struct {
	Symbol    string  `url:"symbol"`              // Required
	Direction *string `url:"direction,omitempty"` // Optional: NEXT, PREV (default NEXT)
	ID        *int64  `url:"id,omitempty"`        // Optional: ID for pagination anchor
	Limit     *int    `url:"limit,omitempty"`     // Optional: default 10
	StartTime *int64  `url:"startTime,omitempty"` // Optional: filter by time (ms)
	EndTime   *int64  `url:"endTime,omitempty"`   // Optional: filter by time (ms)
}

// GetHistoryList queries order history.
// Endpoint: GET /future/trade/v1/order/list-history
func (c *Client) GetHistoryList(ctx context.Context, queryReq GetHistoryListRequest) (*GetHistoryListResult, error) {
	path := "/future/trade/v1/order/list-history"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"symbol": queryReq.Symbol, // Required
	}
	if queryReq.Direction != nil {
		params["direction"] = *queryReq.Direction
	}
	if queryReq.ID != nil {
		params["id"] = strconv.FormatInt(*queryReq.ID, 10)
	}
	if queryReq.Limit != nil {
		params["limit"] = strconv.Itoa(*queryReq.Limit)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetHistoryListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetHistoryList for %s failed: %w", queryReq.Symbol, err)
	}
	return &result, nil
}

// GetTradeListRequest defines parameters for querying trade details.
type GetTradeListRequest struct {
	OrderID   *int64  `url:"orderId,omitempty"`   // Optional: Filter by order ID
	Symbol    *string `url:"symbol,omitempty"`    // Optional: Filter by symbol
	Page      *int    `url:"page,omitempty"`      // Optional: default 1
	Size      *int    `url:"size,omitempty"`      // Optional: default 10
	StartTime *int64  `url:"startTime,omitempty"` // Optional: filter by time (ms)
	EndTime   *int64  `url:"endTime,omitempty"`   // Optional: filter by time (ms)
}

// GetTradeList queries transaction details.
// Endpoint: GET /future/trade/v1/order/trade-list
func (c *Client) GetTradeList(ctx context.Context, queryReq GetTradeListRequest) (*GetTradeListResult, error) {
	path := "/future/trade/v1/order/trade-list"
	baseURL := c.getBaseURL("USDT-M")
	params := make(map[string]string)
	if queryReq.OrderID != nil {
		params["orderId"] = strconv.FormatInt(*queryReq.OrderID, 10)
	}
	if queryReq.Symbol != nil {
		params["symbol"] = *queryReq.Symbol
	}
	if queryReq.Page != nil {
		params["page"] = strconv.Itoa(*queryReq.Page)
	}
	if queryReq.Size != nil {
		params["size"] = strconv.Itoa(*queryReq.Size)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetTradeListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTradeList failed: %w", err)
	}
	return &result, nil
}

// UpdateOrderRequest defines parameters for updating an order.
type UpdateOrderRequest struct {
	OrderID                   int64   `json:"orderId"`                             // Required
	Price                     *string `json:"price,omitempty"`                     // Optional: Target price
	OrigQty                   *string `json:"origQty,omitempty"`                   // Optional: Target quantity (cont)
	TriggerProfitPrice        *string `json:"triggerProfitPrice,omitempty"`        // Optional: Profit target price
	TriggerStopPrice          *string `json:"triggerStopPrice,omitempty"`          // Optional: Stop-Loss price
	TriggerPriceType          *string `json:"triggerPriceType,omitempty"`          // Optional: INDEX_PRICE, MARK_PRICE, LATEST_PRICE
	ProfitDelegateOrderType   *string `json:"profitDelegateOrderType,omitempty"`   // Optional: LIMIT, MARKET
	ProfitDelegateTimeInForce *string `json:"profitDelegateTimeInForce,omitempty"` // Optional: GTC, IOC, FOK, GTX
	ProfitDelegatePrice       *string `json:"profitDelegatePrice,omitempty"`       // Optional: Take-Profit order price
	StopDelegateOrderType     *string `json:"stopDelegateOrderType,omitempty"`     // Optional: LIMIT, MARKET
	StopDelegateTimeInForce   *string `json:"stopDelegateTimeInForce,omitempty"`   // Optional: GTC, IOC, FOK, GTX
	StopDelegatePrice         *string `json:"stopDelegatePrice,omitempty"`         // Optional: Stop-Loss order price
	FollowUpOrder             *bool   `json:"followUpOrder,omitempty"`             // Optional: If true, indicates chase order
}

// UpdateOrder modifies an existing open order.
// Endpoint: POST /future/trade/v1/order/update
func (c *Client) UpdateOrder(ctx context.Context, updateReq UpdateOrderRequest) (*UpdateOrderResult, error) {
	path := "/future/trade/v1/order/update"
	baseURL := c.getBaseURL("USDT-M")

	// API accepts application/json or application/x-www-form-urlencoded
	// Using JSON for simplicity with optional fields.
	var result UpdateOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, updateReq, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder for ID %d failed: %w", updateReq.OrderID, err)
	}
	return &result, nil
}

// --- Trigger Orders (Plan) ---

// CreatePlanOrderRequest defines parameters for creating trigger orders.
type CreatePlanOrderRequest struct {
	ClientOrderID    *string `json:"clientOrderId,omitempty"` // Optional
	Symbol           string  `json:"symbol"`                  // Required
	OrderSide        string  `json:"orderSide"`               // Required: BUY, SELL
	EntrustType      string  `json:"entrustType"`             // Required: TAKE_PROFIT, STOP, TAKE_PROFIT_MARKET, STOP_MARKET
	OrigQty          string  `json:"origQty"`                 // Required: Quantity (Cont)
	Price            *string `json:"price,omitempty"`         // Required for TAKE_PROFIT, STOP (limit types)
	StopPrice        string  `json:"stopPrice"`               // Required: Trigger price
	TimeInForce      string  `json:"timeInForce"`             // Required: GTC, IOC, FOK, GTX (Market orders only support IOC)
	TriggerPriceType string  `json:"triggerPriceType"`        // Required: INDEX_PRICE, MARK_PRICE, LATEST_PRICE
	PositionSide     string  `json:"positionSide"`            // Required: LONG, SHORT
}

// CreatePlanOrder creates a new trigger order.
// Endpoint: POST /future/trade/v1/entrust/create-plan
func (c *Client) CreatePlanOrder(ctx context.Context, orderReq CreatePlanOrderRequest) (*CreatePlanOrderResult, error) {
	path := "/future/trade/v1/entrust/create-plan"
	baseURL := c.getBaseURL("USDT-M")

	// Basic Validation
	if orderReq.EntrustType == "TAKE_PROFIT" || orderReq.EntrustType == "STOP" {
		if orderReq.Price == nil || *orderReq.Price == "" {
			return nil, fmt.Errorf("price is required for LIMIT trigger orders (TAKE_PROFIT, STOP)")
		}
	}
	if orderReq.EntrustType == "TAKE_PROFIT_MARKET" || orderReq.EntrustType == "STOP_MARKET" {
		if orderReq.TimeInForce != "IOC" {
			// return nil, fmt.Errorf("timeInForce must be IOC for MARKET trigger orders") // Relaxing this based on potential API flexibility
		}
	}

	var result CreatePlanOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, orderReq, &result)
	if err != nil {
		return nil, fmt.Errorf("CreatePlanOrder for %s failed: %w", orderReq.Symbol, err)
	}
	return &result, nil
}

// CancelPlanOrder cancels a single trigger order by ID.
// Endpoint: POST /future/trade/v1/entrust/cancel-plan
func (c *Client) CancelPlanOrder(ctx context.Context, entrustID int64) (*CancelPlanOrderResult, error) {
	path := "/future/trade/v1/entrust/cancel-plan"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{
		"entrustId": strconv.FormatInt(entrustID, 10),
	}
	var result CancelPlanOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelPlanOrder for ID %d failed: %w", entrustID, err)
	}
	return &result, nil
}

// CancelAllPlanOrder cancels all trigger orders for a symbol.
// Endpoint: POST /future/trade/v1/entrust/cancel-all-plan
func (c *Client) CancelAllPlanOrder(ctx context.Context, symbol string) (*CancelAllPlanOrderResult, error) {
	path := "/future/trade/v1/entrust/cancel-all-plan"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{
		"symbol": symbol, // Required
	}
	var result CancelAllPlanOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelAllPlanOrder for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetPlanOrderListRequest defines parameters for querying trigger orders.
type GetPlanOrderListRequest struct {
	Symbol    string `url:"symbol"`              // Required
	Page      *int   `url:"page,omitempty"`      // Optional: default 1
	Size      *int   `url:"size,omitempty"`      // Optional: default 10
	StartTime *int64 `url:"startTime,omitempty"` // Optional: filter by time (ms)
	EndTime   *int64 `url:"endTime,omitempty"`   // Optional: filter by time (ms)
	State     string `url:"state"`               // Required: NOT_TRIGGERED,TRIGGERING,TRIGGERED,USER_REVOCATION,PLATFORM_REVOCATION,EXPIRED,UNFINISHED,HISTORY
}

// GetPlanOrderList queries trigger orders.
// Endpoint: GET /future/trade/v1/entrust/plan-list
func (c *Client) GetPlanOrderList(ctx context.Context, queryReq GetPlanOrderListRequest) (*GetPlanOrderListResult, error) {
	path := "/future/trade/v1/entrust/plan-list"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"symbol": queryReq.Symbol,
		"state":  queryReq.State,
	}
	if queryReq.Page != nil {
		params["page"] = strconv.Itoa(*queryReq.Page)
	}
	if queryReq.Size != nil {
		params["size"] = strconv.Itoa(*queryReq.Size)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetPlanOrderListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPlanOrderList for %s failed: %w", queryReq.Symbol, err)
	}
	return &result, nil
}

// GetPlanOrderDetail gets details of a single trigger order.
// Endpoint: GET /future/trade/v1/entrust/plan-detail
func (c *Client) GetPlanOrderDetail(ctx context.Context, entrustID int64) (*GetPlanOrderDetailResult, error) {
	path := "/future/trade/v1/entrust/plan-detail"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"entrustId": strconv.FormatInt(entrustID, 10),
	}
	var result GetPlanOrderDetailResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPlanOrderDetail for ID %d failed: %w", entrustID, err)
	}
	return &result, nil
}

// GetPlanHistoryListRequest defines parameters for querying trigger order history.
type GetPlanHistoryListRequest struct {
	Symbol    string  `url:"symbol"`              // Required
	Direction *string `url:"direction,omitempty"` // Optional: NEXT, PREV
	ID        *int64  `url:"id,omitempty"`        // Optional: ID for pagination anchor
	Limit     *int    `url:"limit,omitempty"`     // Optional: default 10
	StartTime *int64  `url:"startTime,omitempty"` // Optional: filter by time (ms)
	EndTime   *int64  `url:"endTime,omitempty"`   // Optional: filter by time (ms)
}

// GetPlanHistoryList queries trigger order history.
// Endpoint: GET /future/trade/v1/entrust/plan-list-history
func (c *Client) GetPlanHistoryList(ctx context.Context, queryReq GetPlanHistoryListRequest) (*GetPlanHistoryListResult, error) {
	path := "/future/trade/v1/entrust/plan-list-history"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"symbol": queryReq.Symbol,
	}
	if queryReq.Direction != nil {
		params["direction"] = *queryReq.Direction
	}
	if queryReq.ID != nil {
		params["id"] = strconv.FormatInt(*queryReq.ID, 10)
	}
	if queryReq.Limit != nil {
		params["limit"] = strconv.Itoa(*queryReq.Limit)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetPlanHistoryListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetPlanHistoryList for %s failed: %w", queryReq.Symbol, err)
	}
	return &result, nil
}

// --- Stop Limit Orders ---

// CreateProfitStopRequest defines parameters for creating stop limit orders.
type CreateProfitStopRequest struct {
	Symbol             string `json:"symbol"`               // Required
	OrigQty            string `json:"origQty"`              // Required: Quantity (Cont)
	TriggerProfitPrice string `json:"triggerProfitPrice"`   // Required: TP trigger price
	TriggerStopPrice   string `json:"triggerStopPrice"`     // Required: SL trigger price
	ExpireTime         *int64 `json:"expireTime,omitempty"` // Optional: Expiration time (ms)
	PositionSide       string `json:"positionSide"`         // Required: LONG, SHORT
}

// CreateProfitStop creates a new stop limit order for a position.
// Endpoint: POST /future/trade/v1/entrust/create-profit
func (c *Client) CreateProfitStop(ctx context.Context, orderReq CreateProfitStopRequest) (*CreateProfitStopResult, error) {
	path := "/future/trade/v1/entrust/create-profit"
	baseURL := c.getBaseURL("USDT-M")
	var result CreateProfitStopResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, orderReq, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateProfitStop for %s (%s) failed: %w", orderReq.Symbol, orderReq.PositionSide, err)
	}
	return &result, nil
}

// CancelProfitStop cancels a single stop limit order by ID.
// Endpoint: POST /future/trade/v1/entrust/cancel-profit-stop
func (c *Client) CancelProfitStop(ctx context.Context, profitID int64) (*CancelProfitStopResult, error) {
	path := "/future/trade/v1/entrust/cancel-profit-stop"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{
		"profitId": strconv.FormatInt(profitID, 10),
	}
	var result CancelProfitStopResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelProfitStop for ID %d failed: %w", profitID, err)
	}
	return &result, nil
}

// CancelAllProfitStop cancels all stop limit orders for a symbol.
// Endpoint: POST /future/trade/v1/entrust/cancel-all-profit-stop
func (c *Client) CancelAllProfitStop(ctx context.Context, symbol string) (*CancelAllProfitStopResult, error) {
	path := "/future/trade/v1/entrust/cancel-all-profit-stop"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{
		"symbol": symbol, // Required
	}
	var result CancelAllProfitStopResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelAllProfitStop for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetProfitStopListRequest defines parameters for querying stop limit orders.
type GetProfitStopListRequest struct {
	Symbol    string `url:"symbol"`              // Required
	Page      *int   `url:"page,omitempty"`      // Optional: default 1
	Size      *int   `url:"size,omitempty"`      // Optional: default 10
	StartTime *int64 `url:"startTime,omitempty"` // Optional: filter by time (ms)
	EndTime   *int64 `url:"endTime,omitempty"`   // Optional: filter by time (ms)
	State     string `url:"state"`               // Required: NOT_TRIGGERED,TRIGGERING,TRIGGERED,USER_REVOCATION,PLATFORM_REVOCATION,EXPIRED,UNFINISHED,HISTORY
}

// GetProfitStopList queries stop limit orders.
// Endpoint: GET /future/trade/v1/entrust/profit-list
func (c *Client) GetProfitStopList(ctx context.Context, queryReq GetProfitStopListRequest) (*GetProfitStopListResult, error) {
	path := "/future/trade/v1/entrust/profit-list"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"symbol": queryReq.Symbol,
		"state":  queryReq.State,
	}
	if queryReq.Page != nil {
		params["page"] = strconv.Itoa(*queryReq.Page)
	}
	if queryReq.Size != nil {
		params["size"] = strconv.Itoa(*queryReq.Size)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}

	var result GetProfitStopListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetProfitStopList for %s failed: %w", queryReq.Symbol, err)
	}
	return &result, nil
}

// GetProfitStopDetail gets details of a single stop limit order.
// Endpoint: GET /future/trade/v1/entrust/profit-detail
func (c *Client) GetProfitStopDetail(ctx context.Context, profitID int64) (*GetProfitStopDetailResult, error) {
	path := "/future/trade/v1/entrust/profit-detail"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"profitId": strconv.FormatInt(profitID, 10),
	}
	var result GetProfitStopDetailResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetProfitStopDetail for ID %d failed: %w", profitID, err)
	}
	return &result, nil
}

// UpdateProfitStopRequest defines parameters for altering stop limit orders.
type UpdateProfitStopRequest struct {
	ProfitID           int64   `json:"profitId"`                     // Required: Stop limit ID
	TriggerProfitPrice *string `json:"triggerProfitPrice,omitempty"` // Optional: TP trigger price
	TriggerStopPrice   *string `json:"triggerStopPrice,omitempty"`   // Optional: SL trigger price
}

// UpdateProfitStop alters an existing stop limit order.
// Endpoint: POST /future/trade/v1/entrust/update-profit-stop
func (c *Client) UpdateProfitStop(ctx context.Context, updateReq UpdateProfitStopRequest) (*UpdateProfitStopResult, error) {
	path := "/future/trade/v1/entrust/update-profit-stop"
	baseURL := c.getBaseURL("USDT-M")
	var result UpdateProfitStopResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, updateReq, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdateProfitStop for ID %d failed: %w", updateReq.ProfitID, err)
	}
	return &result, nil
}

// --- Track Orders ---

// CreateTrackOrderRequest defines parameters for creating track orders.
type CreateTrackOrderRequest struct {
	Callback           string  `json:"callback"`                     // Required: FIXED, PROPORTION
	CallbackVal        string  `json:"callbackVal"`                  // Required: Callback value (> 0)
	OrderSide          string  `json:"orderSide"`                    // Required: BUY, SELL
	OrigQty            string  `json:"origQty"`                      // Required: Original quantity(count)
	PositionSide       string  `json:"positionSide"`                 // Required: BOTH, LONG, SHORT
	PositionType       string  `json:"positionType"`                 // Required: CROSSED, ISOLATED
	Symbol             string  `json:"symbol"`                       // Required: Trading pair
	TriggerPriceType   string  `json:"triggerPriceType"`             // Required: INDEX_PRICE, MARK_PRICE, LATEST_PRICE
	ActivationPrice    *string `json:"activationPrice,omitempty"`    // Optional: Activation price
	ClientMedia        *string `json:"clientMedia,omitempty"`        // Optional
	ClientMediaChannel *string `json:"clientMediaChannel,omitempty"` // Optional
	ClientOrderID      *string `json:"clientOrderId,omitempty"`      // Optional
	ExpireTime         *int64  `json:"expireTime,omitempty"`         // Optional: expire time (ms)
}

// CreateTrackOrder creates a new track order.
// Endpoint: POST /future/trade/v1/entrust/create-track
func (c *Client) CreateTrackOrder(ctx context.Context, orderReq CreateTrackOrderRequest) (*CreateTrackOrderResult, error) {
	path := "/future/trade/v1/entrust/create-track"
	baseURL := c.getBaseURL("USDT-M")
	// API expects application/x-www-form-urlencoded
	bodyParams := map[string]string{
		"callback":         orderReq.Callback,
		"callbackVal":      orderReq.CallbackVal,
		"orderSide":        orderReq.OrderSide,
		"origQty":          orderReq.OrigQty,
		"positionSide":     orderReq.PositionSide,
		"positionType":     orderReq.PositionType,
		"symbol":           orderReq.Symbol,
		"triggerPriceType": orderReq.TriggerPriceType,
	}
	if orderReq.ActivationPrice != nil {
		bodyParams["activationPrice"] = *orderReq.ActivationPrice
	}
	if orderReq.ClientMedia != nil {
		bodyParams["clientMedia"] = *orderReq.ClientMedia
	}
	if orderReq.ClientMediaChannel != nil {
		bodyParams["clientMediaChannel"] = *orderReq.ClientMediaChannel
	}
	if orderReq.ClientOrderID != nil {
		bodyParams["clientOrderId"] = *orderReq.ClientOrderID
	}
	if orderReq.ExpireTime != nil {
		bodyParams["expireTime"] = strconv.FormatInt(*orderReq.ExpireTime, 10)
	}

	var result CreateTrackOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CreateTrackOrder for %s failed: %w", orderReq.Symbol, err)
	}
	return &result, nil
}

// CancelTrackOrder cancels a single track order.
// Endpoint: POST /future/trade/v1/entrust/cancel-track
func (c *Client) CancelTrackOrder(ctx context.Context, trackID int64) (*CancelTrackOrderResult, error) {
	path := "/future/trade/v1/entrust/cancel-track"
	baseURL := c.getBaseURL("USDT-M")
	bodyParams := map[string]string{
		"trackId": strconv.FormatInt(trackID, 10),
	}
	var result CancelTrackOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelTrackOrder for ID %d failed: %w", trackID, err)
	}
	return &result, nil
}

// GetTrackOrderDetail gets details of a single track order.
// Endpoint: GET /future/trade/v1/entrust/track-detail
func (c *Client) GetTrackOrderDetail(ctx context.Context, trackID int64) (*GetTrackOrderDetailResult, error) {
	path := "/future/trade/v1/entrust/track-detail"
	baseURL := c.getBaseURL("USDT-M")
	params := map[string]string{
		"trackId": strconv.FormatInt(trackID, 10),
	}
	var result GetTrackOrderDetailResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTrackOrderDetail for ID %d failed: %w", trackID, err)
	}
	return &result, nil
}

// GetTrackOrderListRequest defines parameters for querying active track orders.
type GetTrackOrderListRequest struct {
	Page      *int    `url:"page,omitempty"`      // Optional: default 1
	Size      *int    `url:"size,omitempty"`      // Optional: default 10
	EndTime   *int64  `url:"endTime,omitempty"`   // Optional: filter by time (ms)
	StartTime *int64  `url:"startTime,omitempty"` // Optional: filter by time (ms)
	Symbol    *string `url:"symbol,omitempty"`    // Optional
}

// GetTrackOrderList gets the list of active track orders.
// Endpoint: GET /future/trade/v1/entrust/track-list
func (c *Client) GetTrackOrderList(ctx context.Context, queryReq GetTrackOrderListRequest) (*GetTrackOrderListResult, error) {
	path := "/future/trade/v1/entrust/track-list"
	baseURL := c.getBaseURL("USDT-M")
	params := make(map[string]string)
	if queryReq.Page != nil {
		params["page"] = strconv.Itoa(*queryReq.Page)
	}
	if queryReq.Size != nil {
		params["size"] = strconv.Itoa(*queryReq.Size)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.Symbol != nil {
		params["symbol"] = *queryReq.Symbol
	}

	var result GetTrackOrderListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTrackOrderList failed: %w", err)
	}
	return &result, nil
}

// CancelAllTrackOrder cancels all active track orders.
// Endpoint: POST /future/trade/v1/entrust/cancel-all-track
func (c *Client) CancelAllTrackOrder(ctx context.Context) (*CancelAllTrackOrderResult, error) {
	path := "/future/trade/v1/entrust/cancel-all-track"
	baseURL := c.getBaseURL("USDT-M")
	// No body parameters needed
	var result CancelAllTrackOrderResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("CancelAllTrackOrder failed: %w", err)
	}
	return &result, nil
}

// GetTrackHistoryListRequest defines parameters for querying inactive track orders.
type GetTrackHistoryListRequest struct {
	Direction *string `url:"direction,omitempty"` // Optional: NEXT, PREV
	Limit     *int    `url:"limit,omitempty"`     // Optional: default 10
	ID        *int64  `url:"id,omitempty"`        // Optional: ID for pagination anchor
	EndTime   *int64  `url:"endTime,omitempty"`   // Optional: filter by time (ms)
	StartTime *int64  `url:"startTime,omitempty"` // Optional: filter by time (ms)
	Symbol    *string `url:"symbol,omitempty"`    // Optional
}

// GetTrackHistoryList gets the list of inactive (history) track orders.
// Endpoint: GET /future/trade/v1/entrust/track-list-history
func (c *Client) GetTrackHistoryList(ctx context.Context, queryReq GetTrackHistoryListRequest) (*GetTrackHistoryListResult, error) {
	path := "/future/trade/v1/entrust/track-list-history"
	baseURL := c.getBaseURL("USDT-M")
	params := make(map[string]string)
	if queryReq.Direction != nil {
		params["direction"] = *queryReq.Direction
	}
	if queryReq.Limit != nil {
		params["limit"] = strconv.Itoa(*queryReq.Limit)
	}
	if queryReq.ID != nil {
		params["id"] = strconv.FormatInt(*queryReq.ID, 10)
	}
	if queryReq.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*queryReq.EndTime, 10)
	}
	if queryReq.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*queryReq.StartTime, 10)
	}
	if queryReq.Symbol != nil {
		params["symbol"] = *queryReq.Symbol
	}

	var result GetTrackHistoryListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetTrackHistoryList failed: %w", err)
	}
	return &result, nil
}
