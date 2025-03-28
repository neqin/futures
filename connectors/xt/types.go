package xt

import "encoding/json"

// CommonResponse structure for basic API responses
type CommonResponse struct {
	ReturnCode int             `json:"returnCode"` // 0 for success
	MsgInfo    string          `json:"msgInfo"`
	Error      json.RawMessage `json:"error"` // Use RawMessage to handle null or object
}

// --- Public Market Data Structs ---

// ServerTimeResult defines the structure for the server time response
type ServerTimeResult struct {
	CommonResponse
	Result int64 `json:"result"` // Timestamp (milliseconds)
}

// ClientIPResult defines the structure for the client IP response
type ClientIPResult struct {
	CommonResponse
	Result struct {
		IP string `json:"ip"`
	} `json:"result"`
}

// CoinsInfoResult defines the structure for the list of coins response
type CoinsInfoResult struct {
	CommonResponse
	Result []string `json:"result"` // List of coin strings (e.g., ["usdt"])
}

// Contract defines the structure for a single contract's details
type Contract struct {
	ID                        int64    `json:"id"`
	Symbol                    string   `json:"symbol"`
	SymbolGroupId             int      `json:"symbolGroupId"`
	Pair                      string   `json:"pair"`
	ContractType              string   `json:"contractType"` // perpetual, delivery
	ProductType               string   `json:"productType"`  // perpetual, futures
	PredictEventType          *string  `json:"predictEventType"`
	PredictEventParam         *string  `json:"predictEventParam"`
	PredictEventSort          *int     `json:"predictEventSort"`
	UnderlyingType            string   `json:"underlyingType"` // Coin-M, USDT-M
	ContractSize              string   `json:"contractSize"`
	TradeSwitch               bool     `json:"tradeSwitch"`
	OpenSwitch                bool     `json:"openSwitch"`
	IsDisplay                 bool     `json:"isDisplay"`
	IsOpenApi                 bool     `json:"isOpenApi"`
	State                     int      `json:"state"`
	InitLeverage              int      `json:"initLeverage"`
	InitPositionType          string   `json:"initPositionType"`
	BaseCoin                  string   `json:"baseCoin"`
	SpotCoin                  string   `json:"spotCoin"` // Not in v3 list response?
	QuoteCoin                 string   `json:"quoteCoin"`
	SettleCoin                string   `json:"settleCoin"` // Not in v3 list response?
	BaseCoinPrecision         int      `json:"baseCoinPrecision"`
	BaseCoinDisplayPrecision  int      `json:"baseCoinDisplayPrecision"`
	QuoteCoinPrecision        int      `json:"quoteCoinPrecision"`
	QuoteCoinDisplayPrecision int      `json:"quoteCoinDisplayPrecision"`
	QuantityPrecision         int      `json:"quantityPrecision"` // Deprecated in v3 list?
	PricePrecision            int      `json:"pricePrecision"`
	SupportOrderType          string   `json:"supportOrderType"`
	SupportTimeInForce        string   `json:"supportTimeInForce"`
	SupportEntrustType        string   `json:"supportEntrustType"`
	SupportPositionType       string   `json:"supportPositionType"`
	MinQty                    string   `json:"minQty"`
	MinNotional               string   `json:"minNotional"`
	MaxNotional               string   `json:"maxNotional"`
	MultiplierDown            string   `json:"multiplierDown"`
	MultiplierUp              string   `json:"multiplierUp"`
	MaxOpenOrders             int      `json:"maxOpenOrders"`
	MaxEntrusts               int      `json:"maxEntrusts"`
	MakerFee                  string   `json:"makerFee"`
	TakerFee                  string   `json:"takerFee"`
	LiquidationFee            string   `json:"liquidationFee"`
	MarketTakeBound           string   `json:"marketTakeBound"`
	DepthPrecisionMerge       int      `json:"depthPrecisionMerge"`
	Labels                    []string `json:"labels"`
	OnboardDate               int64    `json:"onboardDate"`
	EnName                    string   `json:"enName"`
	CnName                    string   `json:"cnName"`
	MinStepPrice              string   `json:"minStepPrice"`
	MinPrice                  *string  `json:"minPrice"`
	MaxPrice                  *string  `json:"maxPrice"`
	DeliveryDate              *int64   `json:"deliveryDate"`
	DeliveryPrice             *string  `json:"deliveryPrice"`
	DeliveryCompletion        bool     `json:"deliveryCompletion"`
	CnDesc                    *string  `json:"cnDesc"`
	EnDesc                    *string  `json:"enDesc"`
	CnRemark                  *string  `json:"cnRemark"`
	EnRemark                  *string  `json:"enRemark"`
	Plates                    []int    `json:"plates"`
	FastTrackCallbackRate1    *string  `json:"fastTrackCallbackRate1"` // Nullable based on xt.txt
	FastTrackCallbackRate2    *string  `json:"fastTrackCallbackRate2"` // Nullable based on xt.txt
	MinTrackCallbackRate      *string  `json:"minTrackCallbackRate"`   // Nullable based on xt.txt
	MaxTrackCallbackRate      *string  `json:"maxTrackCallbackRate"`   // Nullable based on xt.txt
	LatestPriceDeviation      *float64 `json:"latestPriceDeviation"`   // Nullable based on xt.txt
}

// ContractsResult defines the structure for the list of contracts response (v3 endpoint)
type ContractsResult struct {
	CommonResponse
	Result struct {
		Time    int64      `json:"time"`
		Version string     `json:"version"`
		Symbols []Contract `json:"symbols"` // The list is nested under "symbols"
	} `json:"result"`
}

// SingleContractResult defines the structure for a single contract response
type SingleContractResult struct {
	CommonResponse
	Result Contract `json:"result"`
}

// LeverageBracket defines the structure for leverage stratification details
type LeverageBracket struct {
	Bracket            int    `json:"bracket"`
	MaintMarginRate    string `json:"maintMarginRate"` // Use string for precision
	MaxLeverage        string `json:"maxLeverage"`     // Use string
	MaxNominalValue    string `json:"maxNominalValue"` // Use string for precision
	MaxStartMarginRate string `json:"maxStartMarginRate"`
	MinLeverage        string `json:"minLeverage"`     // Use string
	StartMarginRate    string `json:"startMarginRate"` // Use string for precision
	Symbol             string `json:"symbol"`
}

// LeverageDetail defines the structure for leverage details for a symbol
type LeverageDetail struct {
	LeverageBrackets []LeverageBracket `json:"leverageBrackets"`
	Symbol           string            `json:"symbol"`
}

// LeverageDetailResult defines the structure for the leverage detail response (single symbol)
type LeverageDetailResult struct {
	CommonResponse
	Result LeverageDetail `json:"result"`
}

// LeverageDetailListResult defines the structure for the leverage detail list response (all symbols)
type LeverageDetailListResult struct {
	CommonResponse
	Result []LeverageDetail `json:"result"`
}

// TickerDetail defines the structure for a single ticker.
type TickerDetail struct {
	Amount      string `json:"a"` // 24h volume (Quote currency?)
	Close       string `json:"c"` // Latest price
	High        string `json:"h"` // 24h High
	Low         string `json:"l"` // 24h Low
	Open        string `json:"o"` // 24h Open
	ChangeRatio string `json:"r"` // 24h Change Ratio
	Symbol      string `json:"s"` // Trading pair
	Timestamp   int64  `json:"t"` // Timestamp (ms)
	Volume      string `json:"v"` // 24h Turnover (Base currency?)
}

// TickersResult defines the structure for the all tickers response.
type TickersResult struct {
	CommonResponse
	Result []TickerDetail `json:"result"`
}

// SingleTickerResult defines the structure for the single ticker response.
type SingleTickerResult struct {
	CommonResponse
	Result TickerDetail `json:"result"`
}

// Trade defines the structure for a single public trade (deal)
type Trade struct {
	Amount string `json:"a"` // Volume
	Maker  string `json:"m"` // Order side (BUY/SELL?)
	Price  string `json:"p"` // Price
	Symbol string `json:"s"` // Trading pair
	Time   int64  `json:"t"` // Time (ms)
}

// TradesResult defines the structure for the recent trades response
type TradesResult struct {
	CommonResponse
	Result []Trade `json:"result"`
}

// DepthEntry represents a single price level in the order book [price, quantity]
type DepthEntry [2]string

// DepthResult defines the structure for the order book depth response
type DepthResult struct {
	CommonResponse
	Result struct {
		Asks     []DepthEntry `json:"a"` // Ask levels [price, quantity]
		Bids     []DepthEntry `json:"b"` // Bid levels [price, quantity]
		Symbol   string       `json:"s"` // Symbol
		Time     int64        `json:"t"` // Timestamp (ms)
		UpdateID int64        `json:"u"` // Update ID
	} `json:"result"`
}

// IndexPriceDetail defines the structure for index price info.
type IndexPriceDetail struct {
	Price  string `json:"p"` // Price
	Symbol string `json:"s"` // Trading pair
	Time   int64  `json:"t"` // Time (ms)
}

// IndexPriceResult defines the structure for the single index price response.
type IndexPriceResult struct {
	CommonResponse
	Result IndexPriceDetail `json:"result"`
}

// AllIndexPriceResult defines the structure for the all index prices response.
type AllIndexPriceResult struct {
	CommonResponse
	Result []IndexPriceDetail `json:"result"`
}

// MarkPriceDetail defines the structure for mark price info.
type MarkPriceDetail struct {
	Price  string `json:"p"` // Price
	Symbol string `json:"s"` // Trading pair
	Time   int64  `json:"t"` // Time (ms)
}

// SingleMarkPriceResult defines the structure for the single mark price response.
type SingleMarkPriceResult struct {
	CommonResponse
	Result MarkPriceDetail `json:"result"`
}

// MarkPriceResult defines the structure for the mark price response when fetching all symbols.
type MarkPriceResult struct {
	CommonResponse
	Result []MarkPriceDetail `json:"result"`
}

// Kline defines the structure for a single candlestick
type Kline struct {
	Amount string `json:"a"` // Volume (Turnover in quote currency?)
	Close  string `json:"c"` // Close price
	High   string `json:"h"` // Highest price
	Low    string `json:"l"` // Lowest price
	Open   string `json:"o"` // Open price
	Symbol string `json:"s"` // Trading pair
	Time   int64  `json:"t"` // Time (ms)
	Volume string `json:"v"` // Turnover (Volume in base currency?)
}

// KlinesResult defines the structure for the klines/candlestick response
type KlinesResult struct {
	CommonResponse
	Result []Kline `json:"result"`
}

// AggTickerDetail defines the structure for aggregated ticker information.
type AggTickerDetail struct {
	Timestamp   int64  `json:"t"`  // Timestamp (ms)
	Symbol      string `json:"s"`  // Trading pair
	Close       string `json:"c"`  // Last price
	High        string `json:"h"`  // 24h High
	Low         string `json:"l"`  // 24h Low
	Amount      string `json:"a"`  // 24h Volume (Quote currency?)
	Volume      string `json:"v"`  // 24h Volume (Base currency?)
	Open        string `json:"o"`  // 24h Open
	ChangeRatio string `json:"r"`  // 24h Change Ratio
	IndexPrice  string `json:"i"`  // Index Price
	MarkPrice   string `json:"m"`  // Mark Price
	BidPrice    string `json:"bp"` // Best Bid price
	AskPrice    string `json:"ap"` // Best Ask price
}

// AggTickerResult defines the structure for the single aggregated ticker response.
type AggTickerResult struct {
	CommonResponse
	Result AggTickerDetail `json:"result"`
}

// AllAggTickerResult defines the structure for the all aggregated tickers response.
type AllAggTickerResult struct {
	CommonResponse
	Result []AggTickerDetail `json:"result"`
}

// FundingRateDetail defines the structure for a single funding rate record.
type FundingRateDetail struct {
	Symbol             string  `json:"symbol"`
	FundingRate        string  `json:"fundingRate"` // Use string for precision
	NextCollectionTime *int64  `json:"nextCollectionTime,omitempty"`
	CollectionInternal *int    `json:"collectionInternal,omitempty"`
	ID                 *string `json:"id,omitempty"`          // Only in record list
	CreatedTime        *int64  `json:"createdTime,omitempty"` // Only in record list
}

// FundingRateResult defines the structure for the GetFundRate response.
type FundingRateResult struct {
	CommonResponse
	Result FundingRateDetail `json:"result"` // API returns single object
}

// FundRateRecordResult defines the structure for the GetFundRateRecord response.
type FundRateRecordResult struct {
	CommonResponse
	Result struct {
		HasPrev bool                `json:"hasPrev"`
		HasNext bool                `json:"hasNext"`
		Items   []FundingRateDetail `json:"items"`
	} `json:"result"`
}

// BookTickerDetail defines the structure for ask/bid ticker info.
type BookTickerDetail struct {
	AskPrice  string `json:"ap"` // ask price
	AskQty    string `json:"aq"` // ask amount
	BidPrice  string `json:"bp"` // bid price
	BidQty    string `json:"bq"` // bid amount
	Symbol    string `json:"s"`  // Trading pair
	Timestamp int64  `json:"t"`  // Time (ms)
}

// BookTickerResult defines the structure for the single book ticker response.
type BookTickerResult struct {
	CommonResponse
	Result BookTickerDetail `json:"result"`
}

// AllBookTickerResult defines the structure for the all book tickers response.
type AllBookTickerResult struct {
	CommonResponse
	Result []BookTickerDetail `json:"result"`
}

// RiskBalanceDetail defines the structure for risk balance information.
type RiskBalanceDetail struct {
	ID          string `json:"id"`          // ID is string in response
	Coin        string `json:"coin"`        // Coin
	Amount      string `json:"amount"`      // Amount (use string for precision)
	CreatedTime int64  `json:"createdTime"` // Time (ms)
}

// RiskBalanceResult defines the structure for the risk balance response.
type RiskBalanceResult struct {
	CommonResponse
	Result struct {
		HasPrev bool                `json:"hasPrev"`
		HasNext bool                `json:"hasNext"`
		Items   []RiskBalanceDetail `json:"items"`
	} `json:"result"`
}

// OpenInterestDetail defines the structure for open interest information.
type OpenInterestDetail struct {
	Symbol          string `json:"symbol"`          // Trading pair
	OpenInterest    string `json:"openInterest"`    // open position
	OpenInterestUsd string `json:"openInterestUsd"` // open value (use string)
	Time            int64  `json:"time"`            // time (ms)
}

// OpenInterestResult defines the structure for the open interest response.
type OpenInterestResult struct {
	CommonResponse
	Result OpenInterestDetail `json:"result"`
}

// --- Private Account/User Structs ---

// BalanceDetail defines the structure for a single asset's balance.
type BalanceDetail struct {
	Coin                  string `json:"coin"`                  // e.g., "usdt"
	AvailableBalance      string `json:"availableBalance"`      // Available balance
	IsolatedMargin        string `json:"isolatedMargin"`        // Frozen isolated margin
	OpenOrderMarginFrozen string `json:"openOrderMarginFrozen"` // Frozen order margin
	CrossedMargin         string `json:"crossedMargin"`         // Crossed Margin
	Bonus                 string `json:"bonus"`                 // Bonus
	Coupon                string `json:"coupon"`                // Coupon
	WalletBalance         string `json:"walletBalance"`         // Balance
}

// BalanceListResult defines the structure for the list of all asset balances.
type BalanceListResult struct {
	CommonResponse
	Result []BalanceDetail `json:"result"`
}

// CompatBalanceDetail defines the structure from the compat balance endpoint.
type CompatBalanceDetail struct {
	AccountID             int64  `json:"accountId"`
	UserID                int64  `json:"userId"`
	Coin                  string `json:"coin"`
	UnderlyingType        int    `json:"underlyingType"` // 1: Coin-M, 2: USDT-M
	WalletBalance         string `json:"walletBalance"`
	OpenOrderMarginFrozen string `json:"openOrderMarginFrozen"`
	IsolatedMargin        string `json:"isolatedMargin"`
	CrossedMargin         string `json:"crossedMargin"`
	Amount                string `json:"amount"`      // Net asset balance
	TotalAmount           string `json:"totalAmount"` // Margin balance
	ConvertBtcAmount      string `json:"convertBtcAmount"`
	ConvertUsdtAmount     string `json:"convertUsdtAmount"`
	Profit                string `json:"profit"`    // Realized PNL?
	NotProfit             string `json:"notProfit"` // Unrealized PNL?
	Bonus                 string `json:"bonus"`
	Coupon                string `json:"coupon"`
}

// CompatBalanceListResult defines the structure for the compat balance list response.
type CompatBalanceListResult struct {
	CommonResponse
	Result []CompatBalanceDetail `json:"result"`
}

// AccountInfoResult defines the structure for the account information response.
type AccountInfoResult struct {
	CommonResponse
	Result struct {
		AccountID         int64  `json:"accountId"`
		AllowOpenPosition bool   `json:"allowOpenPosition"`
		AllowTrade        bool   `json:"allowTrade"`
		AllowTransfer     bool   `json:"allowTransfer"`
		OpenTime          string `json:"openTime"` // Assuming string, might be int64 timestamp
		State             int    `json:"state"`
		UserID            int64  `json:"userId"`
	} `json:"result"`
}

// ListenKeyResult defines the structure for the listen key response.
type ListenKeyResult struct {
	CommonResponse
	Result struct {
		ListenKey string `json:"listenKey"`
	} `json:"result"`
}

// AccountOpenResult defines the structure for the account open response.
type AccountOpenResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// GetBalanceResult defines the structure for the single currency balance response.
type GetBalanceResult struct {
	CommonResponse
	Result BalanceDetail `json:"result"`
}

// BalanceBillDetail defines the structure for a single balance bill entry.
type BalanceBillDetail struct {
	AfterAmount string `json:"afterAmount"` // Balance after change
	Amount      string `json:"amount"`      // Quantity
	Coin        string `json:"coin"`        // Currency
	CreatedTime int64  `json:"createdTime"` // Time (ms)
	ID          int64  `json:"id"`          // id
	Side        string `json:"side"`        // ADD:transfer in;SUB:transfer out
	Symbol      string `json:"symbol"`      // Trading pair
	Type        string `json:"type"`        // EXCHANGE:transfer;CLOSE_POSITION:Offset profit and loss;TAKE_OVER:position takeover;QIANG_PING_MANAGER:Liquidation management fee (fee);FUND:Fund Fee;FEE:Fee(Open position, liquidation, Forced liquidation);ADL:Adl;TAKE_OVER:position takeover;MERGE:Position Merge
}

// GetBalanceBillsResult defines the structure for the balance bills response.
type GetBalanceBillsResult struct {
	CommonResponse
	Result struct {
		HasPrev bool                `json:"hasPrev"`
		HasNext bool                `json:"hasNext"`
		Items   []BalanceBillDetail `json:"items"`
	} `json:"result"`
}

// UserFundingRateDetail defines the structure for user funding rate entries.
type UserFundingRateDetail struct {
	Cast         string `json:"cast"`         // Fund fee
	Coin         string `json:"coin"`         // Currency
	CreatedTime  int64  `json:"createdTime"`  // Time (ms)
	ID           int64  `json:"id"`           // id
	PositionSide string `json:"positionSide"` // Direction
	Symbol       string `json:"symbol"`       // Trading pair
}

// GetUserFundingRateListResult defines the structure for the user funding fees response.
type GetUserFundingRateListResult struct {
	CommonResponse
	Result struct {
		HasPrev bool                    `json:"hasPrev"`
		HasNext bool                    `json:"hasNext"`
		Items   []UserFundingRateDetail `json:"items"`
	} `json:"result"`
}

// PositionDetail defines the structure for a single open position.
type PositionDetail struct {
	AutoMargin            bool    `json:"autoMargin"`            // Whether to automatically call margin
	AvailableCloseSize    string  `json:"availableCloseSize"`    // Available quantity (Cont)
	BreakPrice            string  `json:"breakPrice"`            // Blowout price (Liquidation price?)
	CalMarkPrice          string  `json:"calMarkPrice"`          // Calculated mark price
	CloseOrderSize        string  `json:"closeOrderSize"`        // Quantity of open order (Cont)
	ContractType          string  `json:"contractType"`          // Contract Types: PERPETUAL (Perpetual Contract), PREDICT (Predict Contract)
	EntryPrice            string  `json:"entryPrice"`            // Average opening price
	FloatingPL            string  `json:"floatingPL"`            // Unrealized profit or loss
	IsolatedMargin        string  `json:"isolatedMargin"`        // Warehouse-by-warehouse margin
	Leverage              int     `json:"leverage"`              // Leverage ratio (use int based on response example)
	OpenOrderMarginFrozen string  `json:"openOrderMarginFrozen"` // Occupation of deposit for opening order
	OpenOrderSize         string  `json:"openOrderSize"`         // Opening warehouse orders occupied (Not in /list response?)
	PositionSide          string  `json:"positionSide"`          // Position direction
	PositionSize          string  `json:"positionSize"`          // Position quantity (Cont)
	PositionType          string  `json:"positionType"`          // Position type: CROSSED (full position); ISOLATED (warehouse by warehouse)
	ProfitID              *int64  `json:"profitId"`              // Take profit and stop loss id (nullable)
	RealizedProfit        string  `json:"realizedProfit"`        // Realized profit and loss
	Symbol                string  `json:"symbol"`                // trading pair
	TriggerPriceType      *string `json:"triggerPriceType"`      // Trigger price type (nullable)
	TriggerProfitPrice    *string `json:"triggerProfitPrice"`    // Take profit trigger price (nullable)
	TriggerStopPrice      *string `json:"triggerStopPrice"`      // Stop loss trigger price (nullable)
	WelfareAccount        *bool   `json:"welfareAccount"`        // Nullable?
}

// GetPositionsResult defines the structure for the get positions response.
type GetPositionsResult struct {
	CommonResponse
	Result []PositionDetail `json:"result"`
}

// StepRateResult defines the structure for the user step rate response.
type StepRateResult struct {
	CommonResponse
	Result struct {
		MakerFee string `json:"makerFee"`
		TakerFee string `json:"takerFee"`
	} `json:"result"`
}

// AdjustLeverageResult defines the structure for the adjust leverage response.
type AdjustLeverageResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// UpdatePositionMarginResult defines the structure for the update margin response.
type UpdatePositionMarginResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// AllPositionCloseResult defines the structure for the close all positions response.
type AllPositionCloseResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// PositionADLDetail defines the structure for ADL information for a symbol.
type PositionADLDetail struct {
	LongQuantile  int    `json:"longQuantile"`  // long position adl
	ShortQuantile int    `json:"shortQuantile"` // Short position adl
	Symbol        string `json:"symbol"`        // Trading pair
}

// PositionADLResult defines the structure for the ADL response.
type PositionADLResult struct {
	CommonResponse
	Result []PositionADLDetail `json:"result"`
}

// CollectionAddResult defines the structure for the add collection response.
type CollectionAddResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// CollectionCancelResult defines the structure for the cancel collection response.
type CollectionCancelResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// CollectionListResult defines the structure for the list collection response.
type CollectionListResult struct {
	CommonResponse
	Result []string `json:"result"`
}

// ChangePositionTypeResult defines the structure for the change position type response.
type ChangePositionTypeResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// BreakPositionDetail defines the structure for margin call info.
type BreakPositionDetail struct {
	BreakPrice     string `json:"breakPrice"`     // Margin call price. 0 means no margin call
	CalMarkPrice   string `json:"calMarkPrice"`   // Mark price
	ContractType   string `json:"contractType"`   // Futures type: PERPETUAL;PREDICT
	EntryPrice     string `json:"entryPrice"`     // Open position average price
	IsolatedMargin string `json:"isolatedMargin"` // Isolated Margin
	Leverage       int    `json:"leverage"`       // Leverage
	PositionSide   string `json:"positionSide"`   // Position side:LONG;SHORT
	PositionSize   string `json:"positionSize"`   // Position quantity (Cont)
	PositionType   string `json:"positionType"`   // Position type:CROSSED;ISOLATED
	Symbol         string `json:"symbol"`         // Symbol
}

// BreakListResult defines the structure for the margin call list response.
type BreakListResult struct {
	CommonResponse
	Result []BreakPositionDetail `json:"result"`
}

// --- Private Trading Structs ---

// PlaceOrderResult defines the structure for the place order response.
type PlaceOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// OrderDetail defines the structure for detailed order information.
type OrderDetail struct {
	ClientOrderID      *string `json:"clientOrderId"`      // Client order ID (nullable)
	AvgPrice           string  `json:"avgPrice"`           // Average price
	ClosePosition      *bool   `json:"closePosition"`      // Whether to close all when order condition is triggered (nullable)
	CloseProfit        string  `json:"closeProfit"`        // Offset profit and loss
	CreatedTime        int64   `json:"createdTime"`        // Create time (ms)
	ExecutedQty        string  `json:"executedQty"`        // Volume (Cont)
	ForceClose         *bool   `json:"forceClose"`         // Is it a liquidation order (nullable)
	MarginFrozen       string  `json:"marginFrozen"`       // Occupied margin
	OrderID            int64   `json:"orderId"`            // Order ID
	OrderSide          string  `json:"orderSide"`          // Order side
	OrderType          string  `json:"orderType"`          // Order type
	OrigQty            string  `json:"origQty"`            // Quantity (Cont)
	PositionSide       string  `json:"positionSide"`       // Position side
	Price              string  `json:"price"`              // Order price
	SourceID           *int64  `json:"sourceId"`           // Triggering conditions ID (nullable)
	State              string  `json:"state"`              // Order state:NEW,PARTIALLY_FILLED,PARTIALLY_CANCELED,FILLED,CANCELED,REJECTED,EXPIRED
	Symbol             string  `json:"symbol"`             // Trading pair
	TimeInForce        string  `json:"timeInForce"`        // Valid type
	TriggerProfitPrice *string `json:"triggerProfitPrice"` // TP trigger price (nullable)
	TriggerStopPrice   *string `json:"triggerStopPrice"`   // SL trigger price (nullable)
}

// GetOrderResult defines the structure for the get order response.
type GetOrderResult struct {
	CommonResponse
	Result OrderDetail `json:"result"`
}

// GetOrderListResult defines the structure for the get order list response.
type GetOrderListResult struct {
	CommonResponse
	Result struct {
		Items []OrderDetail `json:"items"`
		Page  int           `json:"page"`
		Ps    int           `json:"ps"` // Page size?
		Total int           `json:"total"`
	} `json:"result"`
}

// GetHistoryListResult defines the structure for the order history response.
type GetHistoryListResult struct {
	CommonResponse
	Result struct {
		HasNext bool          `json:"hasNext"`
		HasPrev bool          `json:"hasPrev"`
		Items   []OrderDetail `json:"items"`
	} `json:"result"`
}

// TradeDetail defines the structure for a single trade detail.
type TradeDetail struct {
	Fee        string `json:"fee"`        // Fee
	FeeCoin    string `json:"feeCoin"`    // Currency of fee
	OrderID    int64  `json:"orderId"`    // Order ID
	ExecID     string `json:"execId"`     // Trade ID (string in response)
	Price      string `json:"price"`      // Price
	Quantity   string `json:"quantity"`   // Volume
	Symbol     string `json:"symbol"`     // Trading pair
	Timestamp  int64  `json:"timestamp"`  // Time (ms)
	TakerMaker string `json:"takerMaker"` // TAKER or MAKER
}

// GetTradeListResult defines the structure for the trade list response.
type GetTradeListResult struct {
	CommonResponse
	Result struct {
		Items []TradeDetail `json:"items"`
		Page  int           `json:"page"`
		Ps    int           `json:"ps"`
		Total int           `json:"total"`
	} `json:"result"`
}

// UpdateOrderResult defines the structure for the update order response.
type UpdateOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// PlaceBatchOrderResult defines the structure for the batch place order response.
type PlaceBatchOrderResult struct {
	CommonResponse
	Result bool `json:"result"` // API doc shows boolean
}

// CancelOrderResult defines the structure for the cancel order response.
type CancelOrderResult struct {
	CommonResponse
	Result string `json:"result"` // Order ID as string
}

// CancelBatchOrderResult defines the structure for the cancel all orders response.
type CancelBatchOrderResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// PlanOrderDetail defines the structure for trigger orders.
type PlanOrderDetail struct {
	ClientOrderID    *string `json:"clientOrderId"`    // Client order ID (nullable)
	ClosePosition    *bool   `json:"closePosition"`    // Whether triggered to close all (nullable)
	CreatedTime      int64   `json:"createdTime"`      // Create time (ms)
	EntrustID        int64   `json:"entrustId"`        // Order ID
	EntrustType      string  `json:"entrustType"`      // Order type
	MarketOrderLevel *int    `json:"marketOrderLevel"` // Best market price (nullable?)
	OrderSide        string  `json:"orderSide"`        // Order side
	Ordinary         *bool   `json:"ordinary"`         // Nullable?
	OrigQty          string  `json:"origQty"`          // Quantity (Cont)
	PositionSide     string  `json:"positionSide"`     // Position side
	Price            string  `json:"price"`            // Order price
	State            string  `json:"state"`            // Order state: NOT_TRIGGERED,TRIGGERING,TRIGGERED,USER_REVOCATION,PLATFORM_REVOCATION,EXPIRED
	StopPrice        string  `json:"stopPrice"`        // Trigger price
	Symbol           string  `json:"symbol"`           // Trading pair
	TimeInForce      string  `json:"timeInForce"`      // Valid way
	TriggerPriceType string  `json:"triggerPriceType"` // Trigger price type
}

// CreatePlanOrderResult defines the structure for creating trigger orders.
type CreatePlanOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// CancelPlanOrderResult defines the structure for canceling trigger orders.
type CancelPlanOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// CancelAllPlanOrderResult defines the structure for canceling all trigger orders.
type CancelAllPlanOrderResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// GetPlanOrderListResult defines the structure for listing trigger orders.
type GetPlanOrderListResult struct {
	CommonResponse
	Result struct {
		Items []PlanOrderDetail `json:"items"`
		Page  int               `json:"page"`
		Ps    int               `json:"ps"`
		Total int               `json:"total"`
	} `json:"result"`
}

// GetPlanOrderDetailResult defines the structure for single trigger order detail.
type GetPlanOrderDetailResult struct {
	CommonResponse
	Result PlanOrderDetail `json:"result"`
}

// GetPlanHistoryListResult defines the structure for trigger order history.
type GetPlanHistoryListResult struct {
	CommonResponse
	Result struct {
		HasNext bool              `json:"hasNext"`
		HasPrev bool              `json:"hasPrev"`
		Items   []PlanOrderDetail `json:"items"`
	} `json:"result"`
}

// ProfitStopDetail defines the structure for stop limit orders.
type ProfitStopDetail struct {
	CreatedTime        int64  `json:"createdTime"`        // Time (ms)
	EntryPrice         string `json:"entryPrice"`         // Open position average price
	ExecutedQty        string `json:"executedQty"`        // Actual transaction
	IsolatedMargin     string `json:"isolatedMargin"`     // Isolated Margin
	OrigQty            string `json:"origQty"`            // Quantity (Cont)
	PositionSide       string `json:"positionSide"`       // Position side
	PositionSize       string `json:"positionSize"`       // Position quantity (Cont)
	ProfitID           int64  `json:"profitId"`           // Order ID
	State              string `json:"state"`              // Order state: NOT_TRIGGERED,TRIGGERING,TRIGGERED,USER_REVOCATION,PLATFORM_REVOCATION,EXPIRED
	Symbol             string `json:"symbol"`             // Trading pair
	TriggerProfitPrice string `json:"triggerProfitPrice"` // Stop profit price
	TriggerStopPrice   string `json:"triggerStopPrice"`   // Stop loss price
}

// CreateProfitStopResult defines the structure for creating stop limit orders.
type CreateProfitStopResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// CancelProfitStopResult defines the structure for canceling stop limit orders.
type CancelProfitStopResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// CancelAllProfitStopResult defines the structure for canceling all stop limit orders.
type CancelAllProfitStopResult struct {
	CommonResponse
	Result bool `json:"result"`
}

// GetProfitStopListResult defines the structure for listing stop limit orders.
type GetProfitStopListResult struct {
	CommonResponse
	Result struct {
		Items []ProfitStopDetail `json:"items"`
		Page  int                `json:"page"`
		Ps    int                `json:"ps"`
		Total int                `json:"total"`
	} `json:"result"`
}

// GetProfitStopDetailResult defines the structure for single stop limit order detail.
type GetProfitStopDetailResult struct {
	CommonResponse
	Result ProfitStopDetail `json:"result"`
}

// UpdateProfitStopResult defines the structure for altering stop limit orders.
type UpdateProfitStopResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// TrackOrderDetail defines the structure for track orders.
type TrackOrderDetail struct {
	ActivationPrice  string `json:"activationPrice"`  // Activation price
	AvgPrice         string `json:"avgPrice"`         // Average price
	Callback         string `json:"callback"`         // Callback range configuration 1:PROPORTION 2:FIXED
	CallbackVal      string `json:"callbackVal"`      // Callback value (use string)
	ConfigActivation bool   `json:"configActivation"` // Whether to configure activation price
	CreatedTime      int64  `json:"createdTime"`      // Create time (ms)
	CurrentPrice     string `json:"currentPrice"`     // Real-time price
	Desc             string `json:"desc"`             // Describe
	ExecutedQty      string `json:"executedQty"`      // Actual transaction quantity
	OrderSide        string `json:"orderSide"`        // Order side
	Ordinary         bool   `json:"ordinary"`
	OrigQty          string `json:"origQty"`          // Quantity (Cont)
	PositionSide     string `json:"positionSide"`     // Position side
	Price            string `json:"price"`            // Order price
	State            string `json:"state"`            // Order state: NOT_ACTIVATION,NOT_TRIGGERED,TRIGGERING,TRIGGERED,USER_REVOCATION,PLATFORM_REVOCATION,EXPIRED,DELEGATION_FAILED
	StopPrice        string `json:"stopPrice"`        // Trigger price
	Symbol           string `json:"symbol"`           // Symbol
	TrackID          int64  `json:"trackId"`          // Track id
	TriggerPriceType string `json:"triggerPriceType"` // Trigger price type
	UpdatedTime      int64  `json:"updatedTime"`      // Update time (ms)
}

// CreateTrackOrderResult defines the structure for creating track orders.
type CreateTrackOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// CancelTrackOrderResult defines the structure for canceling track orders.
type CancelTrackOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// GetTrackOrderDetailResult defines the structure for single track order detail.
type GetTrackOrderDetailResult struct {
	CommonResponse
	Result TrackOrderDetail `json:"result"`
}

// GetTrackOrderListResult defines the structure for listing track orders.
type GetTrackOrderListResult struct {
	CommonResponse
	Result struct {
		Items []TrackOrderDetail `json:"items"`
		Page  int                `json:"page"`
		Ps    int                `json:"ps"`
		Total int                `json:"total"`
	} `json:"result"`
}

// CancelAllTrackOrderResult defines the structure for canceling all track orders.
type CancelAllTrackOrderResult struct {
	CommonResponse
	Result map[string]interface{} `json:"result"` // Empty object {} on success
}

// GetTrackHistoryListResult defines the structure for track order history.
type GetTrackHistoryListResult struct {
	CommonResponse
	Result struct {
		HasNext bool               `json:"hasNext"`
		HasPrev bool               `json:"hasPrev"`
		Items   []TrackOrderDetail `json:"items"`
	} `json:"result"`
}
