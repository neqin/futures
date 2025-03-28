package gateio

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// GetFuturesAccount retrieves the futures account details for a specific settlement currency.
// settle: "usdt" or "btc"
func (c *Client) GetFuturesAccount(ctx context.Context, settle string) (*FuturesAccount, error) {
	endpoint := fmt.Sprintf("/futures/%s/accounts", settle)
	var result FuturesAccount
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListPositions retrieves all open positions for the user.
// settle: "usdt" or "btc"
// holdingID: Specify a position ID to filter results (optional)
func (c *Client) ListPositions(ctx context.Context, settle string, holdingID *string) (*[]Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/positions", settle)
	params := url.Values{}
	if holdingID != nil {
		params.Set("holding_id", *holdingID) // Note: API doc uses holding_id, might be internal? Check if needed. Usually positions are listed without this.
	}
	var result []Position
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPosition retrieves details for a single position.
// settle: "usdt" or "btc"
// contract: Futures contract name
func (c *Client) GetPosition(ctx context.Context, settle, contract string) (*Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/positions/%s", settle, contract)
	var result Position
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePositionMargin updates the margin for a specific position.
// settle: "usdt" or "btc"
// contract: Futures contract name
// change: The amount to change the margin by (positive to add, negative to reduce).
func (c *Client) UpdatePositionMargin(ctx context.Context, settle, contract, change string) (*Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/positions/%s/margin", settle, contract)
	params := url.Values{}
	params.Set("change", change)
	var result Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePositionLeverage updates the leverage for a specific position.
// settle: "usdt" or "btc"
// contract: Futures contract name
// leverage: The new leverage value (e.g., "10"). "0" means cross margin.
// crossLeverageLimit: Cross margin leverage limit (required for cross margin, e.g., "10").
func (c *Client) UpdatePositionLeverage(ctx context.Context, settle, contract, leverage string, crossLeverageLimit *string) (*Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/positions/%s/leverage", settle, contract)
	params := url.Values{}
	params.Set("leverage", leverage)
	if crossLeverageLimit != nil {
		params.Set("cross_leverage_limit", *crossLeverageLimit)
	}
	var result Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePositionRiskLimit updates the risk limit for a specific position.
// settle: "usdt" or "btc"
// contract: Futures contract name
// riskLimit: The new risk limit value.
func (c *Client) UpdatePositionRiskLimit(ctx context.Context, settle, contract, riskLimit string) (*Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/positions/%s/risk_limit", settle, contract)
	params := url.Values{}
	params.Set("risk_limit", riskLimit)
	var result Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SetDualMode enables or disables dual mode for futures trading.
// settle: "usdt" or "btc"
// dualMode: true to enable dual mode, false to disable.
func (c *Client) SetDualMode(ctx context.Context, settle string, dualMode bool) (*FuturesAccount, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_mode", settle)
	params := url.Values{}
	params.Set("dual_mode", strconv.FormatBool(dualMode))
	var result FuturesAccount
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDualModePosition retrieves position details in dual mode for a specific contract.
// settle: "usdt" or "btc"
// contract: Futures contract name
func (c *Client) GetDualModePosition(ctx context.Context, settle, contract string) (*[]Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/positions/%s", settle, contract)
	var result []Position // Dual mode returns an array of two positions (long and short)
	err := c.get(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDualModePositionMargin updates margin for a position in dual mode.
// settle: "usdt" or "btc"
// contract: Futures contract name
// change: Margin change amount.
// dualSide: Position side, "dual_long" or "dual_short".
func (c *Client) UpdateDualModePositionMargin(ctx context.Context, settle, contract, change, dualSide string) (*[]Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/positions/%s/margin", settle, contract)
	params := url.Values{}
	params.Set("change", change)
	params.Set("dual_side", dualSide)
	var result []Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDualModePositionLeverage updates leverage for a position in dual mode.
// settle: "usdt" or "btc"
// contract: Futures contract name
// leverage: New leverage value. "0" for cross margin.
// crossLeverageLimit: Cross margin leverage limit (required for cross margin).
func (c *Client) UpdateDualModePositionLeverage(ctx context.Context, settle, contract, leverage string, crossLeverageLimit *string) (*[]Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/positions/%s/leverage", settle, contract)
	params := url.Values{}
	params.Set("leverage", leverage)
	if crossLeverageLimit != nil {
		params.Set("cross_leverage_limit", *crossLeverageLimit)
	}
	var result []Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDualModePositionRiskLimit updates risk limit for a position in dual mode.
// settle: "usdt" or "btc"
// contract: Futures contract name
// riskLimit: New risk limit value.
func (c *Client) UpdateDualModePositionRiskLimit(ctx context.Context, settle, contract, riskLimit string) (*[]Position, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/positions/%s/risk_limit", settle, contract)
	params := url.Values{}
	params.Set("risk_limit", riskLimit)
	var result []Position
	// Note: API uses POST for this, but query parameters.
	err := c.post(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFuturesAccountBook queries the account book (ledger) entries.
// settle: "usdt" or "btc"
// contract: Filter by contract name (optional)
// limit: Maximum number of records. Default 100, Max 1000.
// from: Start timestamp (seconds) (optional)
// to: End timestamp (seconds) (optional)
// typeFilter: Filter by entry type (dnw, pnl, fee, refr, fund, point_dnw, point_fee, point_refr, bonus_offset) (optional)
func (c *Client) ListFuturesAccountBook(ctx context.Context, settle string, contract *string, limit *int, from, to *int64, typeFilter *string) (*ListFuturesAccountBookResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/account_book", settle)
	params := url.Values{}
	if contract != nil {
		params.Set("contract", *contract)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}
	if typeFilter != nil {
		params.Set("type", *typeFilter)
	}

	var result ListFuturesAccountBookResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListPositionCloseHistory lists the history of closed positions.
// settle: "usdt" or "btc"
// contract: Filter by contract name (optional)
// limit: Maximum number of records. Default 100, Max 1000.
// offset: List offset (optional)
// from: Start timestamp (seconds) (optional)
// to: End timestamp (seconds) (optional)
// side: Filter by position side ("long" or "short") (optional)
// pnl: Filter by PNL (optional)
func (c *Client) ListPositionCloseHistory(ctx context.Context, settle string, contract *string, limit, offset *int, from, to *int64, side, pnl *string) (*ListPositionCloseResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/position_close", settle)
	params := url.Values{}
	if contract != nil {
		params.Set("contract", *contract)
	}
	if limit != nil {
		params.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		params.Set("offset", strconv.Itoa(*offset))
	}
	if from != nil {
		params.Set("from", strconv.FormatInt(*from, 10))
	}
	if to != nil {
		params.Set("to", strconv.FormatInt(*to, 10))
	}
	if side != nil {
		params.Set("side", *side)
	}
	if pnl != nil {
		params.Set("pnl", *pnl)
	}

	var result ListPositionCloseResult
	err := c.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Moved from market_public.go as they require authentication ---

// ListDualCompContracts retrieves list of dual swap contracts. (Requires Auth)
// settle: "usdt" or "btc"
func (c *Client) ListDualCompContracts(ctx context.Context, settle string) (*TickerResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/contracts", settle)
	var result TickerResult                   // Uses the same Ticker/Contract struct
	err := c.get(ctx, endpoint, nil, &result) // Uses c.get which handles auth
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListDualCompIndexConstituents retrieves constituent list of a dual swap index. (Requires Auth)
// settle: "usdt" or "btc"
// index: Index name (e.g., BTC_USDT_DUAL)
func (c *Client) ListDualCompIndexConstituents(ctx context.Context, settle, index string) (*ListDualCompIndexConstituentsResult, error) {
	endpoint := fmt.Sprintf("/futures/%s/dual_comp/index_constituents/%s", settle, index)
	var result ListDualCompIndexConstituentsResult
	err := c.get(ctx, endpoint, nil, &result) // Uses c.get which handles auth
	if err != nil {
		return nil, err
	}
	return &result, nil
}
