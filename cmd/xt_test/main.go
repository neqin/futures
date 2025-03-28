package main

import (
	"context"
	"log"
	"os" // Added for environment variables
	"time"

	"github.com/joho/godotenv"               // Added for .env loading
	"github.com/neqin/futures/connectors/xt" // Adjust import path if needed
)

func main() {
	// Load .env file from the root directory
	// Assumes the test is run from the project root (e.g., go run ./cmd/xt_test)
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Warning: Could not load .env.local file:", err)
	}

	apiKey := os.Getenv("XT_API_KEY")
	secretKey := os.Getenv("XT_API_SECRET")

	if apiKey == "" || secretKey == "" {
		log.Println("Warning: XT_API_KEY or XT_API_SECRET not found in environment. Skipping private tests.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second) // Increased timeout
	defer cancel()

	// --- Public Client Tests ---
	publicClient := xt.NewPublicOnly(nil)
	log.Println("Testing XT.com Public API Methods...")
	testPublicMethods(ctx, publicClient)

	// --- Private Client Tests ---
	if apiKey != "" && secretKey != "" {
		privateClient := xt.New(apiKey, secretKey, nil)
		log.Println("\nTesting XT.com Private API Methods...")
		testPrivateMethods(ctx, privateClient)
	} else {
		log.Println("\nSkipping XT.com Private API Method Tests (API keys not found).")
	}

	log.Println("\nXT.com API Method Tests Complete.")
}

func testPublicMethods(ctx context.Context, client *xt.Client) {
	symbol := "btc_usdt" // Example symbol

	// --- Test GetServerTime ---
	log.Println("Fetching server time...")
	serverTime, err := client.GetServerTime(ctx)
	if err != nil {
		log.Printf("ERROR fetching server time: %v\n", err)
	} else {
		log.Printf("OK: Server Time (ms): %d\n", serverTime.Result)
	}

	// --- Test GetMarketTickers ---
	log.Println("Fetching all market tickers...")
	tickers, err := client.GetMarketTickers(ctx)
	if err != nil {
		log.Printf("ERROR fetching all tickers: %v\n", err)
	} else if tickers != nil && len(tickers.Result) > 0 {
		log.Printf("OK: Fetched %d tickers. First: %s (Close: %s)\n", len(tickers.Result), tickers.Result[0].Symbol, tickers.Result[0].Close)
	} else {
		log.Println("WARN: Fetched tickers, but list is empty or nil.")
	}

	// --- Test GetMarketTicker ---
	log.Printf("Fetching ticker for %s...\n", symbol)
	ticker, err := client.GetMarketTicker(ctx, symbol)
	if err != nil {
		log.Printf("ERROR fetching ticker for %s: %v\n", symbol, err)
	} else if ticker != nil {
		log.Printf("OK: Fetched ticker for %s. Close: %s, High: %s, Low: %s\n", ticker.Result.Symbol, ticker.Result.Close, ticker.Result.High, ticker.Result.Low)
	} else {
		log.Println("WARN: Fetched ticker, but result is nil.")
	}

	// --- Test GetDepth ---
	log.Printf("Fetching depth for %s...\n", symbol)
	depthLevel := 5 // Fetch top 5 levels
	depth, err := client.GetDepth(ctx, symbol, depthLevel)
	if err != nil {
		log.Printf("ERROR fetching depth for %s: %v\n", symbol, err)
	} else if depth != nil && len(depth.Result.Asks) > 0 && len(depth.Result.Bids) > 0 {
		log.Printf("OK: Fetched depth for %s. Ask[0]: %s @ %s, Bid[0]: %s @ %s\n",
			symbol, depth.Result.Asks[0][1], depth.Result.Asks[0][0], depth.Result.Bids[0][1], depth.Result.Bids[0][0])
	} else {
		log.Printf("WARN: Fetched depth for %s, but asks or bids are empty or nil.\n", symbol)
	}

	// --- Test GetKlines ---
	log.Printf("Fetching 1m klines for %s...\n", symbol)
	klineLimit := 5
	interval := "1m"
	klines, err := client.GetKlines(ctx, symbol, interval, nil, nil, &klineLimit)
	if err != nil {
		log.Printf("ERROR fetching klines for %s: %v\n", symbol, err)
	} else if klines != nil && len(klines.Result) > 0 {
		log.Printf("OK: Fetched %d klines for %s. First Kline Time: %d, Open: %s\n", len(klines.Result), symbol, klines.Result[0].Time, klines.Result[0].Open)
	} else {
		log.Println("WARN: Fetched klines, but list is empty or nil.")
	}

	// --- Test GetMarketDeal ---
	log.Printf("Fetching recent deals for %s...\n", symbol)
	dealNum := 10
	deals, err := client.GetMarketDeal(ctx, symbol, dealNum)
	if err != nil {
		log.Printf("ERROR fetching deals for %s: %v\n", symbol, err)
	} else if deals != nil && len(deals.Result) > 0 {
		log.Printf("OK: Fetched %d deals for %s. First Deal Time: %d, Price: %s\n", len(deals.Result), symbol, deals.Result[0].Time, deals.Result[0].Price)
	} else {
		log.Println("WARN: Fetched deals, but list is empty or nil.")
	}

	// --- Test GetAllMarketConfigV3 ---
	log.Println("Fetching all market config (v3)...")
	config, err := client.GetAllMarketConfigV3(ctx)
	if err != nil {
		log.Printf("ERROR fetching all market config: %v\n", err)
	} else if config != nil && len(config.Result.Symbols) > 0 { // Access nested Symbols field
		log.Printf("OK: Fetched %d market configs. First: %s\n", len(config.Result.Symbols), config.Result.Symbols[0].Symbol)
	} else {
		log.Println("WARN: Fetched market configs, but list is empty or nil.")
	}
}

func testPrivateMethods(ctx context.Context, client *xt.Client) {
	symbol := "btc_usdt" // Example symbol

	// --- Test GetAccountInfo ---
	log.Println("Fetching account info...")
	accountInfo, err := client.GetAccountInfo(ctx)
	if err != nil {
		log.Printf("ERROR fetching account info: %v\n", err)
	} else if accountInfo != nil {
		log.Printf("OK: Fetched account info. UserID: %d, AllowTrade: %t\n", accountInfo.Result.UserID, accountInfo.Result.AllowTrade)
	} else {
		log.Println("WARN: Fetched account info, but result is nil.")
	}

	// --- Test GetBalanceList ---
	log.Println("Fetching balance list...")
	balanceList, err := client.GetBalanceList(ctx)
	if err != nil {
		log.Printf("ERROR fetching balance list: %v\n", err)
	} else if balanceList != nil && len(balanceList.Result) > 0 {
		log.Printf("OK: Fetched %d balances. First: %s (Wallet: %s)\n", len(balanceList.Result), balanceList.Result[0].Coin, balanceList.Result[0].WalletBalance)
	} else {
		log.Println("WARN: Fetched balance list, but list is empty or nil.")
	}

	// --- Test GetPositions ---
	log.Println("Fetching positions...")
	positions, err := client.GetPositions(ctx, nil) // Fetch all positions
	if err != nil {
		log.Printf("ERROR fetching positions: %v\n", err)
	} else if positions != nil {
		log.Printf("OK: Fetched %d positions.\n", len(positions.Result))
		if len(positions.Result) > 0 {
			log.Printf("  First position: %s (%s), Size: %s, Entry: %s\n", positions.Result[0].Symbol, positions.Result[0].PositionSide, positions.Result[0].PositionSize, positions.Result[0].EntryPrice)
		}
	} else {
		log.Println("WARN: Fetched positions, but result is nil.")
	}

	// --- Test GetOrderList (Unfinished) ---
	log.Println("Fetching unfinished orders...")
	orderState := "UNFINISHED" // Example state
	orderListReq := xt.GetOrderListRequest{State: &orderState}
	orderList, err := client.GetOrderList(ctx, orderListReq)
	if err != nil {
		log.Printf("ERROR fetching unfinished orders: %v\n", err)
	} else if orderList != nil {
		log.Printf("OK: Fetched %d unfinished orders (Total: %d).\n", len(orderList.Result.Items), orderList.Result.Total)
		if len(orderList.Result.Items) > 0 {
			log.Printf("  First unfinished order ID: %d, Symbol: %s\n", orderList.Result.Items[0].OrderID, orderList.Result.Items[0].Symbol)
		}
	} else {
		log.Println("WARN: Fetched unfinished orders, but result is nil.")
	}

	// --- Test GetTradeList ---
	log.Println("Fetching recent trade list...")
	tradeListSize := 5
	tradeListReq := xt.GetTradeListRequest{Size: &tradeListSize, Symbol: &symbol}
	tradeList, err := client.GetTradeList(ctx, tradeListReq)
	if err != nil {
		log.Printf("ERROR fetching trade list for %s: %v\n", symbol, err)
	} else if tradeList != nil {
		log.Printf("OK: Fetched %d trades for %s (Total: %d).\n", len(tradeList.Result.Items), symbol, tradeList.Result.Total)
		if len(tradeList.Result.Items) > 0 {
			log.Printf("  First trade ExecID: %s, Price: %s\n", tradeList.Result.Items[0].ExecID, tradeList.Result.Items[0].Price)
		}
	} else {
		log.Println("WARN: Fetched trade list, but result is nil.")
	}

	// Add more private tests here if needed (e.g., GetListenKey, GetBalanceBills, etc.)
	// Be cautious with order placement/cancellation tests.
}
