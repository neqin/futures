package gateio

import "fmt" // Added for APIError

// TickerResult defines the result for listing contracts or dual contracts.
type TickerResult []Ticker

// Ticker defines the structure for contract details.
type Ticker struct {
	FundingRateIndicative string  `json:"funding_rate_indicative"`
	MarkPriceRound        string  `json:"mark_price_round"`
	FundingOffset         int     `json:"funding_offset"`
	InDelisting           bool    `json:"in_delisting"`
	RiskLimitBase         string  `json:"risk_limit_base"`
	InterestRate          string  `json:"interest_rate"`
	IndexPrice            string  `json:"index_price"`
	OrderPriceRound       string  `json:"order_price_round"`
	OrderSizeMin          int     `json:"order_size_min"`
	RefRebateRate         string  `json:"ref_rebate_rate"`
	Name                  string  `json:"name"` // Contract name
	RefDiscountRate       string  `json:"ref_discount_rate"`
	OrderPriceDeviate     string  `json:"order_price_deviate"`
	MaintenanceRate       string  `json:"maintenance_rate"`
	MarkType              string  `json:"mark_type"`
	FundingInterval       int     `json:"funding_interval"`
	Type                  string  `json:"type"`
	RiskLimitStep         string  `json:"risk_limit_step"`
	EnableBonus           bool    `json:"enable_bonus"`
	EnableCredit          bool    `json:"enable_credit"`
	LeverageMin           string  `json:"leverage_min"`
	FundingRate           string  `json:"funding_rate"`
	LastPrice             float64 `json:"last_price,string"` // Use float64 and string tag for potential flexibility
	MarkPrice             string  `json:"mark_price"`
	OrderSizeMax          int     `json:"order_size_max"`
	FundingNextApply      int     `json:"funding_next_apply"`
	ShortUsers            int     `json:"short_users"`
	ConfigChangeTime      int     `json:"config_change_time"`
	CreateTime            int     `json:"create_time"`
	TradeSize             int     `json:"trade_size"`
	PositionSize          int     `json:"position_size"`
	LongUsers             int     `json:"long_users"`
	QuantoMultiplier      string  `json:"quanto_multiplier"`
	FundingImpactValue    string  `json:"funding_impact_value"`
	LeverageMax           string  `json:"leverage_max"`
	CrossLeverageDefault  string  `json:"cross_leverage_default"`
	RiskLimitMax          string  `json:"risk_limit_max"`
	MakerFeeRate          string  `json:"maker_fee_rate"`
	TakerFeeRate          string  `json:"taker_fee_rate"`
	OrdersLimit           int     `json:"orders_limit"`
	TradeID               int     `json:"trade_id"`
	OrderbookID           int     `json:"orderbook_id"`
	FundingCapRatio       string  `json:"funding_cap_ratio"`
	VoucherLeverage       string  `json:"voucher_leverage"`
	IsPreMarket           bool    `json:"is_pre_market"`
}

// ContractStats defines the statistics of a futures contract.
type ContractStats struct {
	Time                  int64   `json:"time"`                    // Timestamp of the start of the candlestick, in milliseconds
	Loi                   int64   `json:"loi"`                     // Long/Short open interest ratio
	LsrAccount            float64 `json:"lsr_account"`             // Long/Short account ratio (Corrected type)
	LsrTaker              float64 `json:"lsr_taker"`               // Long/Short taker ratio (Corrected type)
	OpenInterest          int64   `json:"open_interest"`           // Open interest size
	MarkPrice             float64 `json:"mark_price"`              // Mark price (Corrected type)
	TopLsrSize            float64 `json:"top_lsr_size"`            // Top L/S size ratio (Corrected type)
	FundingRate           float64 `json:"funding_rate"`            // Funding rate (Corrected type)
	TopLsrAccount         float64 `json:"top_lsr_account"`         // Top L/S account ratio (Corrected type)
	IndexPrice            float64 `json:"index_price"`             // Index price (Corrected type)
	OpenInterestUsd       float64 `json:"open_interest_usd"`       // Open interest in USDT (Corrected type)
	FundingRateIndicative float64 `json:"funding_rate_indicative"` // Indicative Funding rate (Corrected type)
	Contract              string  `json:"contract"`                // Contract name
	Volume                int64   `json:"volume"`                  // Trade size accumulated in 1 day
	VolumeUsd             float64 `json:"volume_usd"`              // Trade volume in USDT accumulated in 1 day (Corrected type)
	// Added fields based on error output
	LongLiqSize     int64   `json:"long_liq_size"`
	ShortLiqSize    int64   `json:"short_liq_size"`
	ShortLiqUsd     float64 `json:"short_liq_usd"`
	TopLongSize     int64   `json:"top_long_size"`
	TopShortSize    int64   `json:"top_short_size"`
	ShortLiqAmount  float64 `json:"short_liq_amount"`
	LongLiqAmount   float64 `json:"long_liq_amount"`
	TopLongAccount  int64   `json:"top_long_account"`
	TopShortAccount int64   `json:"top_short_account"`
	LongLiqUsd      float64 `json:"long_liq_usd"`
	LongTakerSize   int64   `json:"long_taker_size"`
	ShortTakerSize  int64   `json:"short_taker_size"`
	LongUsers       int64   `json:"long_users"`
	ShortUsers      int64   `json:"short_users"`
}

// ListContractStatsResult defines the result for listing contract stats.
type ListContractStatsResult []ContractStats

// FutureOrderBookEntry defines a single entry in the order book.
type FutureOrderBookEntry struct {
	Price string `json:"p"` // Price
	Size  int64  `json:"s"` // Size
}

// FutureOrderBook defines the structure for the futures order book response.
type FutureOrderBook struct {
	ID       int64                  `json:"id"`      // Order Book ID
	Current  float64                `json:"current"` // Current timestamp (seconds with microseconds)
	Update   float64                `json:"update"`  // Update timestamp (seconds with microseconds)
	Asks     []FutureOrderBookEntry `json:"asks"`
	Bids     []FutureOrderBookEntry `json:"bids"`
	Contract string                 `json:"contract"` // Added based on documentation example
}

// FuturesTrade defines the structure for a single futures trade.
type FuturesTrade struct {
	ID         int64   `json:"id"`          // Trade ID
	CreateTime float64 `json:"create_time"` // Trading time (seconds with microseconds)
	Contract   string  `json:"contract"`    // Futures contract name
	Size       int64   `json:"size"`        // Trading size, >0 means buy, <0 means sell
	Price      string  `json:"price"`       // Trading price
}

// ListFuturesTradesResult defines the result for listing futures trades.
type ListFuturesTradesResult []FuturesTrade

// CandlestickData represents the structure of a single candlestick object from the API.
type CandlestickData struct {
	Timestamp int64   `json:"t"`          // Timestamp (seconds) - Corrected type
	Volume    int64   `json:"v"`          // Volume (Using float64 with ,string for flexibility)
	Close     float64 `json:"c,string"`   // Close price (Using float64 with ,string)
	High      float64 `json:"h,string"`   // High price (Using float64 with ,string)
	Low       float64 `json:"l,string"`   // Low price (Using float64 with ,string)
	Open      float64 `json:"o,string"`   // Open price (Using float64 with ,string)
	Sum       float64 `json:"sum,string"` // Total traded value (Using float64 with ,string)
}

// FuturesCandlestick represents a single candlestick entry (now an object).
type FuturesCandlestick CandlestickData // Alias for clarity, uses the object structure

// ListFuturesCandlesticksResult defines the result for listing futures candlesticks.
type ListFuturesCandlesticksResult []FuturesCandlestick // Reverted: Array of objects

// PremiumIndexData represents the structure of a single premium index object from the API.
// Corrected based on API documentation: [timestamp, mark_price, index_price]
type PremiumIndexData struct {
	Timestamp  int64   `json:"t"`        // Timestamp (seconds) - Corrected type
	MarkPrice  float64 `json:"m,string"` // Mark price (Assuming 'm', using float64 with ,string)
	IndexPrice float64 `json:"i,string"` // Index price (Assuming 'i', using float64 with ,string)
}

// FuturesPremiumIndex defines the structure for premium index K-line data (now an object).
type FuturesPremiumIndex PremiumIndexData // Alias for clarity, uses the object structure

// ListFuturesPremiumIndexResult defines the result for listing premium index k-lines.
type ListFuturesPremiumIndexResult []FuturesPremiumIndex // Reverted: Array of objects

// FuturesTicker defines the structure for a futures ticker.
type FuturesTicker struct {
	Contract              string  `json:"contract"`                // Futures contract name
	Last                  string  `json:"last"`                    // Last traded price
	ChangePercentage      string  `json:"change_percentage"`       // Change percentage.
	TotalSize             string  `json:"total_size"`              // Total size traded in the last 24 hours
	Low24H                string  `json:"low_24h"`                 // Lowest price in 24h
	High24H               string  `json:"high_24h"`                // Highest price in 24h
	Volume24H             string  `json:"volume_24h"`              // Trade size in the last 24 hours
	Volume24HBtc          string  `json:"volume_24h_btc"`          // Trade volumes in BTC in the last 24 hours
	Volume24HUsd          string  `json:"volume_24h_usd"`          // Trade volumes in USD in the last 24 hours
	Volume24HQuote        string  `json:"volume_24h_quote"`        // Trade volumes in quote currency in the last 24 hours
	MarkPrice             string  `json:"mark_price"`              // Mark price
	FundingRate           string  `json:"funding_rate"`            // Funding rate
	FundingRateIndicative string  `json:"funding_rate_indicative"` // Indicative Funding rate
	IndexPrice            string  `json:"index_price"`             // Index price
	QuantoBaseRate        *string `json:"quanto_base_rate"`        // Quanto base rate (nullable)
	HighestBid            *string `json:"highest_bid"`             // Highest bid price (nullable)
	LowestAsk             *string `json:"lowest_ask"`              // Lowest ask price (nullable)
}

// ListFuturesTickersResult defines the result for listing futures tickers.
type ListFuturesTickersResult []FuturesTicker

// FundingRate defines the structure for a funding rate history entry.
type FundingRate struct {
	Timestamp int64  `json:"t"` // Timestamp (seconds)
	Rate      string `json:"r"` // Funding rate
}

// ListFuturesFundingRateHistoryResult defines the result for listing funding rate history.
type ListFuturesFundingRateHistoryResult []FundingRate

// InsuranceRecord defines the structure for an insurance ledger entry.
type InsuranceRecord struct {
	Timestamp int64  `json:"t"` // Timestamp (seconds)
	Change    string `json:"d"` // Change amount
}

// ListFuturesInsuranceLedgerResult defines the result for listing insurance ledger records.
type ListFuturesInsuranceLedgerResult []InsuranceRecord

// IndexConstituent defines the structure for index constituents.
type IndexConstituent struct {
	Index    string   `json:"index"`    // Index name
	Exchange string   `json:"exchange"` // Exchange name
	Symbols  []string `json:"symbols"`  // Symbol list
}

// ListDualCompIndexConstituentsResult defines the result for listing index constituents.
// The actual API response is a map where keys are exchange names.
type ListDualCompIndexConstituentsResult map[string][]string

// LiquidationOrder defines the structure for a liquidation order record.
type LiquidationOrder struct {
	Time       int64  `json:"time"`        // Liquidation time (seconds)
	Contract   string `json:"contract"`    // Futures contract
	Size       int64  `json:"size"`        // Position size liquidated
	Leverage   string `json:"leverage"`    // Position leverage
	Margin     string `json:"margin"`      // Position margin
	EntryPrice string `json:"entry_price"` // Average entry price
	LiqPrice   string `json:"liq_price"`   // Liquidation price
	MarkPrice  string `json:"mark_price"`  // Mark price at liquidation time
	OrderID    int64  `json:"order_id"`    // Order ID of the liquidation order
	OrderPrice string `json:"order_price"` // Order price of the liquidation order
	FillPrice  string `json:"fill_price"`  // Fill price of the liquidation order
	Left       int64  `json:"left"`        // Size left after liquidation
}

// GetLiquidationHistoryResult defines the result for listing liquidation history.
type GetLiquidationHistoryResult []LiquidationOrder

// RiskLimitTier defines the structure for a risk limit tier.
type RiskLimitTier struct {
	Tier            int    `json:"tier"`             // Risk limit tier
	RiskLimit       string `json:"risk_limit"`       // Risk limit
	InitialRate     string `json:"initial_rate"`     // Initial margin rate
	MaintenanceRate string `json:"maintenance_rate"` // Maintenance margin rate
	LeverageMax     string `json:"leverage_max"`     // Maximum leverage
}

// GetRiskLimitTiersResult defines the result for listing risk limit tiers.
type GetRiskLimitTiersResult []RiskLimitTier

// FuturesAccount defines the structure for futures account details.
type FuturesAccount struct {
	User                  int      `json:"user"`                    // User ID
	Total                 string   `json:"total"`                   // Total account balance in USDT
	UnrealisedPnl         string   `json:"unrealised_pnl"`          // Unrealized PNL
	PositionMargin        string   `json:"position_margin"`         // Position margin
	OrderMargin           string   `json:"order_margin"`            // Order margin
	Available             string   `json:"available"`               // Available balance
	Point                 string   `json:"point"`                   // POINT amount
	Currency              string   `json:"currency"`                // Settle currency
	InDualMode            bool     `json:"in_dual_mode"`            // Whether dual mode is enabled
	EnableCredit          bool     `json:"enable_credit"`           // Whether portfolio margin account mode is enabled
	PositionInitialMargin string   `json:"position_initial_margin"` // Initial margin position
	MaintenanceMargin     string   `json:"maintenance_margin"`      // Maintenance margin position
	Bonus                 string   `json:"bonus"`                   // Perpetual Contract Bonus
	History               struct { // History stats
		Pnl         string `json:"pnl"`          // PNL
		Fee         string `json:"fee"`          // Fee
		Refr        string `json:"refr"`         // Referral fee
		Fund        string `json:"fund"`         // Funding fee
		PointPnl    string `json:"point_pnl"`    // POINT PNL
		PointFee    string `json:"point_fee"`    // POINT fee
		PointRefr   string `json:"point_refr"`   // POINT referral fee
		BonusPnl    string `json:"bonus_pnl"`    // Bonus PNL
		BonusOffset string `json:"bonus_offset"` // Bonus deduction
	} `json:"history"`
}

// Position defines the structure for a futures position.
type Position struct {
	User               int                 `json:"user"`                 // User ID
	Contract           string              `json:"contract"`             // Futures contract
	Size               int64               `json:"size"`                 // Position size
	Leverage           string              `json:"leverage"`             // Position leverage
	RiskLimit          string              `json:"risk_limit"`           // Position risk limit
	LeverageMax        string              `json:"leverage_max"`         // Maximum leverage under current risk limit
	MaintenanceRate    string              `json:"maintenance_rate"`     // Maintenance rate under current risk limit
	Value              string              `json:"value"`                // Position value
	Margin             string              `json:"margin"`               // Position margin
	EntryPrice         string              `json:"entry_price"`          // Entry price
	LiqPrice           string              `json:"liq_price"`            // Liquidation price
	MarkPrice          string              `json:"mark_price"`           // Mark price
	InitialMargin      string              `json:"initial_margin"`       // Initial margin
	MaintenanceMargin  string              `json:"maintenance_margin"`   // Maintenance margin
	UnrealisedPnl      string              `json:"unrealised_pnl"`       // Unrealized PNL
	RealisedPnl        string              `json:"realised_pnl"`         // Realized PNL
	HistoryPnl         string              `json:"history_pnl"`          // History PNL
	LastClosePnl       string              `json:"last_close_pnl"`       // PNL from last close
	RealisedPoint      string              `json:"realised_point"`       // Realized POINT PNL
	HistoryPoint       string              `json:"history_point"`        // History POINT PNL
	AdlRanking         int                 `json:"adl_ranking"`          // ADL ranking, range from 1 to 5
	PendingOrders      int                 `json:"pending_orders"`       // Current open orders count for the position
	CloseOrder         *PositionCloseOrder `json:"close_order"`          // Position close order (nullable)
	Mode               string              `json:"mode"`                 // Position mode, single or dual.
	CrossLeverageLimit string              `json:"cross_leverage_limit"` // Cross margin leverage(valid only when cross margin is used)
}

// PositionCloseOrder defines the structure for a position's close order.
type PositionCloseOrder struct {
	ID    int64  `json:"id"`     // Close order ID
	Price string `json:"price"`  // Close order price
	IsLiq bool   `json:"is_liq"` // Is the close order a liquidation order
}

// FuturesOrder defines the structure for a futures order.
type FuturesOrder struct {
	ID           int64   `json:"id"`             // Futures order ID
	User         int     `json:"user"`           // User ID
	CreateTime   float64 `json:"create_time"`    // Creation time
	FinishTime   float64 `json:"finish_time"`    // Finish time
	FinishAs     string  `json:"finish_as"`      // How the order was finished. Enum: "filled", "cancelled", "liquidated", "ioc", "auto_deleveraged", "reduce_only", "position_closed", "reduce_out"
	Status       string  `json:"status"`         // Order status. Enum: "open", "finished"
	Contract     string  `json:"contract"`       // Futures contract
	Size         int64   `json:"size"`           // Order size. Positive means buy, negative means sell. Set to 0 to close the position
	Iceberg      int64   `json:"iceberg"`        // Display size for iceberg order. 0 for non-iceberg. Note that you will have to pay the taker fee for the hidden size
	Price        string  `json:"price"`          // Order price. 0 for market order with tif set as ioc
	Close        bool    `json:"close"`          // Set as true to close the position, with size set to 0
	IsClose      bool    `json:"is_close"`       // Is the order to close position
	ReduceOnly   bool    `json:"reduce_only"`    // Set as true to be reduce-only order
	IsReduceOnly bool    `json:"is_reduce_only"` // Is the order reduce-only
	IsLiq        bool    `json:"is_liq"`         // Is the order for liquidation
	Tif          string  `json:"tif"`            // Time in force. Enum: "gtc", "ioc", "poc", "fok"
	Left         int64   `json:"left"`           // Size left to be traded
	FillPrice    string  `json:"fill_price"`     // Fill price of the order
	Text         string  `json:"text"`           // User defined information. If not empty, must follow the rules below:  1. prefixed with t- 2. no longer than 28 bytes without prefix, consisting of letters, numbers, underscores, hyphen -, periods .
	Tkfr         string  `json:"tkfr"`           // Taker fee
	Mkfr         string  `json:"mkfr"`           // Maker fee
	Refu         int     `json:"refu"`           // Reference user ID
	AutoSize     string  `json:"auto_size"`      // Set side to close dual-mode position. Required if close is true. ("long", "short")
	StpAct       string  `json:"stp_act"`        // Self-Trading Prevention Action. Enum: "cn", "co", "cb", ""
	StpID        int     `json:"stp_id"`         // Self-Trading Prevention ID. Orders with the same stp_id will be prevented from matching. Valid range: [1, 9223372036854775807]
}

// CreateFuturesOrderRequest defines the structure for creating a futures order.
type CreateFuturesOrderRequest struct {
	Contract   string  `json:"contract"`              // Futures contract
	Size       int64   `json:"size"`                  // Order size. Positive means buy, negative means sell. Set to 0 to close the position
	Iceberg    *int64  `json:"iceberg,omitempty"`     // Display size for iceberg order. 0 for non-iceberg.
	Price      *string `json:"price,omitempty"`       // Order price. Set to 0 to use market price
	Close      bool    `json:"close,omitempty"`       // Set as true to close the position, with size set to 0
	ReduceOnly bool    `json:"reduce_only,omitempty"` // Set as true to be reduce-only order
	Tif        string  `json:"tif,omitempty"`         // Time in force. gtc, ioc, poc, fok. Defaults to gtc
	Text       string  `json:"text,omitempty"`        // User defined information. prefixed with t-
	AutoSize   string  `json:"auto_size,omitempty"`   // Set side to close dual-mode position. Required if close is true. ("long", "short")
	StpAct     string  `json:"stp_act,omitempty"`     // Self-Trading Prevention Action. cn, co, cb, ""
	StpID      *int    `json:"stp_id,omitempty"`      // Self-Trading Prevention ID.
}

// FuturesAccountBookEntry defines the structure for an account book entry.
type FuturesAccountBookEntry struct {
	Time     float64 `json:"time"`     // Change time
	Change   string  `json:"change"`   // Change amount
	Balance  string  `json:"balance"`  // Balance after change
	Type     string  `json:"type"`     // Changing Type: - dnw: Deposit & Withdraw - pnl: PNL - fee: Trading fee - refr: Referrer rebate - fund: Funding fee - point_dnw: POINT Deposit & Withdraw - point_fee: POINT Trading fee - point_refr: POINT Referrer rebate - bonus_offset: bonus offset
	Text     string  `json:"text"`     // Comment
	Contract string  `json:"contract"` // Futures contract, Required for pnl, fee, fund type
	TradeID  string  `json:"trade_id"` // Trade ID, Required for fee, pnl type
}

// ListFuturesAccountBookResult defines the result for listing account book entries.
type ListFuturesAccountBookResult []FuturesAccountBookEntry

// PositionClose defines the structure for closing a position.
type PositionClose struct {
	Time     float64 `json:"time"`     // Position close time
	Contract string  `json:"contract"` // Futures contract
	Side     string  `json:"side"`     // Position side, long or short
	Pnl      string  `json:"pnl"`      // PNL
	Text     string  `json:"text"`     // Text of close order
}

// ListPositionCloseResult defines the result for listing position close history.
type ListPositionCloseResult []PositionClose

// TriggerOrder defines the structure for a price trigger order.
type TriggerOrder struct {
	Initial    FuturesOrder `json:"initial"`     // Order details upon creation
	Trigger    Trigger      `json:"trigger"`     // Trigger condition
	Trail      *Trail       `json:"trail"`       // Trailing parameters (nullable)
	Status     string       `json:"status"`      // Status: open, cancelled, finished, failed, expired
	FinishTime int64        `json:"finish_time"` // Finish timestamp
	TradeID    int64        `json:"trade_id"`    // Corresponding trade ID
	FinishAs   string       `json:"finish_as"`   // How the order is finished
	Reason     string       `json:"reason"`      // Additional information for failure or cancellation
	OrderType  string       `json:"order_type"`  // Order type, "positional" or "contractual" or "conditional"
	MeOrderID  string       `json:"me_order_id"` // Corresponding order ID generated by matching engine
}

// Trigger defines the trigger condition for a price trigger order.
type Trigger struct {
	Price      string `json:"price"`      // Trigger price
	Rule       int    `json:"rule"`       // Trigger rule. 1: >=, 2: <=
	Expiration int    `json:"expiration"` // Trigger expiration time in seconds
	PriceType  string `json:"price_type"` // Price type, 0 - latest price, 1 - mark price, 2 - index price
}

// Trail defines the trailing parameters for a price trigger order.
type Trail struct {
	Amount string `json:"amount"` // Trailing amount
	Offset string `json:"offset"` // Trailing offset
}

// CreateTriggerOrderRequest defines the structure for creating a price trigger order.
type CreateTriggerOrderRequest struct {
	Initial   FuturesOrder `json:"initial"`              // Order details
	Trigger   Trigger      `json:"trigger"`              // Trigger condition
	Trail     *Trail       `json:"trail,omitempty"`      // Trailing parameters
	Settle    string       `json:"settle"`               // Settle currency (usdt or btc)
	OrderType string       `json:"order_type,omitempty"` // Order type, "positional" or "contractual" or "conditional"
}

// CancelOrderResult defines the result of cancelling a single order.
type CancelOrderResult FuturesOrder // Reuses FuturesOrder structure

// BatchCancelOrdersRequest defines the structure for batch cancelling orders.
type BatchCancelOrdersRequest struct {
	Contract string `json:"contract"`       // Futures contract
	Side     string `json:"side,omitempty"` // Optional: "buy" or "sell"
}

// BatchCancelOrdersResult defines the result of batch cancelling orders.
type BatchCancelOrdersResult []FuturesOrder // List of cancelled orders

// CountdownCancelAllFuturesRequest defines the structure for countdown cancel all orders.
type CountdownCancelAllFuturesRequest struct {
	Timeout  int    `json:"timeout"`            // Countdown time in seconds. 0 to cancel the countdown
	Contract string `json:"contract,omitempty"` // Optional: Futures contract to cancel orders for
}

// PriceTriggeredOrder defines the structure for a price triggered order (used in list response).
type PriceTriggeredOrder struct {
	ID         int64        `json:"id"`          // Auto order ID
	User       int          `json:"user"`        // User ID
	Contract   string       `json:"contract"`    // Futures contract
	CreateTime int64        `json:"create_time"` // Creation time
	Trigger    Trigger      `json:"trigger"`     // Trigger settings
	Initial    FuturesOrder `json:"initial"`     // Initial order details
	Status     string       `json:"status"`      // Order status
	Reason     string       `json:"reason"`      // Additional reason for modification or cancellation
	OrderType  string       `json:"order_type"`  // Order type
}

// ListPriceTriggeredOrdersResult defines the result for listing price triggered orders.
type ListPriceTriggeredOrdersResult []PriceTriggeredOrder

// CancelPriceTriggeredOrderResult defines the result of cancelling a price triggered order.
type CancelPriceTriggeredOrderResult PriceTriggeredOrder // Reuses PriceTriggeredOrder structure

// APIError defines the standard error response structure from Gate.io API v4.
type APIError struct {
	Label   string `json:"label"`   // Error label
	Message string `json:"message"` // Error message
}

// Error returns the error message string.
func (e APIError) Error() string {
	return fmt.Sprintf("Gate.io API Error: %s - %s", e.Label, e.Message)
}
