package xt

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// --- Private Account/User Endpoints ---

// GetAccountInfo fetches the user's futures account information.
// Endpoint: GET /future/user/v1/account/info
func (c *Client) GetAccountInfo(ctx context.Context) (*AccountInfoResult, error) {
	path := "/future/user/v1/account/info"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result AccountInfoResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo failed: %w", err)
	}
	return &result, nil
}

// GetListenKey gets a listen key for user data stream. Valid for 8 hours.
// Endpoint: GET /future/user/v1/user/listen-key
func (c *Client) GetListenKey(ctx context.Context) (*ListenKeyResult, error) {
	path := "/future/user/v1/user/listen-key"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result ListenKeyResult
	// Docs say GET, but xt.txt example uses POST? Let's try GET first based on docs.
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetListenKey failed: %w", err)
	}
	return &result, nil
}

// AccountOpen opens the futures account for the user.
// Endpoint: POST /future/user/v1/account/open
func (c *Client) AccountOpen(ctx context.Context) (*AccountOpenResult, error) {
	path := "/future/user/v1/account/open"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result AccountOpenResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, nil, &result) // POST with empty body
	if err != nil {
		return nil, fmt.Errorf("AccountOpen failed: %w", err)
	}
	return &result, nil
}

// GetBalance gets the user's single-currency funds.
// Endpoint: GET /future/user/v1/balance/detail
func (c *Client) GetBalance(ctx context.Context, coin string) (*GetBalanceResult, error) {
	path := "/future/user/v1/balance/detail"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{
		"coin": coin,
	}
	var result GetBalanceResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetBalance for %s failed: %w", coin, err)
	}
	return &result, nil
}

// GetBalanceList gets the user's funds information for all currencies.
// Endpoint: GET /future/user/v1/balance/list
func (c *Client) GetBalanceList(ctx context.Context) (*BalanceListResult, error) {
	path := "/future/user/v1/balance/list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result BalanceListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetBalanceList failed: %w", err)
	}
	return &result, nil
}

// GetCompatBalanceList gets contract account assets (alternative endpoint).
// Endpoint: GET /future/user/v1/compat/balance/list
func (c *Client) GetCompatBalanceList(ctx context.Context, queryAccountID *string) (*CompatBalanceListResult, error) {
	path := "/future/user/v1/compat/balance/list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{}
	if queryAccountID != nil {
		params["queryAccountId"] = *queryAccountID
	}
	var result CompatBalanceListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCompatBalanceList failed: %w", err)
	}
	return &result, nil
}

// GetBalanceBills gets user account flow (ledger).
// Endpoint: GET /future/user/v1/balance/bills
func (c *Client) GetBalanceBills(ctx context.Context, symbol string, direction *string, id *int64, limit *int, startTime, endTime *int64) (*GetBalanceBillsResult, error) {
	path := "/future/user/v1/balance/bills"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{
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
	if startTime != nil {
		params["startTime"] = strconv.FormatInt(*startTime, 10)
	}
	if endTime != nil {
		params["endTime"] = strconv.FormatInt(*endTime, 10)
	}
	// Add type filter if needed based on API capabilities

	var result GetBalanceBillsResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetBalanceBills for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetFundingRateList gets user funding rate fees.
// Endpoint: GET /future/user/v1/balance/funding-rate-list
func (c *Client) GetFundingRateList(ctx context.Context, symbol string, direction *string, id *int64, limit *int, startTime, endTime *int64) (*GetUserFundingRateListResult, error) {
	path := "/future/user/v1/balance/funding-rate-list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{
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
	if startTime != nil {
		params["startTime"] = strconv.FormatInt(*startTime, 10)
	}
	if endTime != nil {
		params["endTime"] = strconv.FormatInt(*endTime, 10)
	}

	var result GetUserFundingRateListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetFundingRateList for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// GetPositions fetches the user's open positions. Uses /v1/position/list endpoint.
// Endpoint: GET /future/user/v1/position/list
func (c *Client) GetPositions(ctx context.Context, symbol *string) (*GetPositionsResult, error) {
	path := "/future/user/v1/position/list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{}
	if symbol != nil {
		params["symbol"] = *symbol
	}
	var result GetPositionsResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		symbolStr := "all symbols"
		if symbol != nil {
			symbolStr = *symbol
		}
		return nil, fmt.Errorf("GetPositions for %s failed: %w", symbolStr, err)
	}
	return &result, nil
}

// GetActivePositions fetches active position information (alternative endpoint).
// Endpoint: GET /future/user/v1/position
func (c *Client) GetActivePositions(ctx context.Context, symbol *string) (*GetPositionsResult, error) {
	path := "/future/user/v1/position" // Different endpoint path
	baseURL := c.getBaseURL("USDT-M")  // Assuming USDT-M
	params := map[string]string{}
	if symbol != nil {
		params["symbol"] = *symbol
	}
	var result GetPositionsResult // Assuming the result structure is the same as /list
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		symbolStr := "all symbols"
		if symbol != nil {
			symbolStr = *symbol
		}
		return nil, fmt.Errorf("GetActivePositions for %s failed: %w", symbolStr, err)
	}
	return &result, nil
}

// GetUserStepRate gets the user's current fee tier rate.
// Endpoint: GET /future/user/v1/user/step-rate
func (c *Client) GetUserStepRate(ctx context.Context) (*StepRateResult, error) {
	path := "/future/user/v1/user/step-rate"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result StepRateResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("GetUserStepRate failed: %w", err)
	}
	return &result, nil
}

// AdjustLeverage adjusts the leverage ratio for a position.
// Endpoint: POST /future/user/v1/position/adjust-leverage
func (c *Client) AdjustLeverage(ctx context.Context, symbol, positionSide string, leverage int) (*AdjustLeverageResult, error) {
	path := "/future/user/v1/position/adjust-leverage"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	bodyParams := map[string]string{  // Docs indicate x-www-form-urlencoded or JSON, let's try map for form
		"symbol":       symbol,
		"positionSide": positionSide,
		"leverage":     strconv.Itoa(leverage),
	}
	var result AdjustLeverageResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("AdjustLeverage for %s (%s) failed: %w", symbol, positionSide, err)
	}
	return &result, nil
}

// UpdatePositionMargin modifies the margin for an isolated position.
// Endpoint: POST /future/user/v1/position/margin
func (c *Client) UpdatePositionMargin(ctx context.Context, symbol, margin, marginType string, positionSide *string) (*UpdatePositionMarginResult, error) {
	path := "/future/user/v1/position/margin"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M

	if marginType != "ADD" && marginType != "SUB" {
		return nil, fmt.Errorf("invalid marginType: must be ADD or SUB")
	}

	bodyParams := map[string]string{ // Docs indicate x-www-form-urlencoded or JSON, let's try map for form
		"symbol": symbol,
		"margin": margin,
		"type":   marginType,
	}
	if positionSide != nil {
		bodyParams["positionSide"] = *positionSide
	}

	var result UpdatePositionMarginResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("UpdatePositionMargin for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// AllPositionClose closes all open positions.
// Endpoint: POST /future/user/v1/position/close-all
func (c *Client) AllPositionClose(ctx context.Context) (*AllPositionCloseResult, error) {
	path := "/future/user/v1/position/close-all"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result AllPositionCloseResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, nil, &result) // POST with empty body
	if err != nil {
		return nil, fmt.Errorf("AllPositionClose failed: %w", err)
	}
	return &result, nil
}

// PositionADL gets ADL (Auto-Deleveraging) information.
// Endpoint: GET /future/user/v1/position/adl
func (c *Client) PositionADL(ctx context.Context) (*PositionADLResult, error) {
	path := "/future/user/v1/position/adl"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result PositionADLResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("PositionADL failed: %w", err)
	}
	return &result, nil
}

// CollectionAdd adds a trading pair to the collection (favorites).
// Endpoint: POST /future/user/v1/user/collection/add
func (c *Client) CollectionAdd(ctx context.Context, symbol string) (*CollectionAddResult, error) {
	path := "/future/user/v1/user/collection/add"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	bodyParams := map[string]string{  // Docs indicate x-www-form-urlencoded or JSON
		"symbol": symbol,
	}
	var result CollectionAddResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CollectionAdd for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// CollectionCancel removes a trading pair from the collection.
// Endpoint: POST /future/user/v1/user/collection/cancel
func (c *Client) CollectionCancel(ctx context.Context, symbol string) (*CollectionCancelResult, error) {
	path := "/future/user/v1/user/collection/cancel"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	bodyParams := map[string]string{  // Docs indicate x-www-form-urlencoded or JSON
		"symbol": symbol,
	}
	var result CollectionCancelResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("CollectionCancel for %s failed: %w", symbol, err)
	}
	return &result, nil
}

// CollectionList lists all collected trading pairs.
// Endpoint: GET /future/user/v1/user/collection/list
func (c *Client) CollectionList(ctx context.Context) (*CollectionListResult, error) {
	path := "/future/user/v1/user/collection/list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	var result CollectionListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("CollectionList failed: %w", err)
	}
	return &result, nil
}

// ChangePositionType changes position type (ISOLATED/CROSSED).
// Endpoint: POST /future/user/v1/position/change-type
func (c *Client) ChangePositionType(ctx context.Context, symbol, positionSide, positionType string) (*ChangePositionTypeResult, error) {
	path := "/future/user/v1/position/change-type"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	bodyParams := map[string]string{  // Docs indicate x-www-form-urlencoded or JSON
		"symbol":       symbol,
		"positionSide": positionSide,
		"positionType": positionType,
	}
	var result ChangePositionTypeResult
	err := c.SendPrivateRequest(ctx, http.MethodPost, baseURL, path, nil, bodyParams, &result)
	if err != nil {
		return nil, fmt.Errorf("ChangePositionType for %s (%s) failed: %w", symbol, positionSide, err)
	}
	return &result, nil
}

// GetBreakList gets margin call information.
// Endpoint: GET /future/user/v1/position/break-list
func (c *Client) GetBreakList(ctx context.Context, symbol *string) (*BreakListResult, error) {
	path := "/future/user/v1/position/break-list"
	baseURL := c.getBaseURL("USDT-M") // Assuming USDT-M
	params := map[string]string{}
	if symbol != nil {
		params["symbol"] = *symbol
	}
	var result BreakListResult
	err := c.SendPrivateRequest(ctx, http.MethodGet, baseURL, path, params, nil, &result)
	if err != nil {
		symbolStr := "all symbols"
		if symbol != nil {
			symbolStr = *symbol
		}
		return nil, fmt.Errorf("GetBreakList for %s failed: %w", symbolStr, err)
	}
	return &result, nil
}

// Note: UpdatePositionAutoMargin is missing from xt.txt/xt2.txt, skipping implementation.
