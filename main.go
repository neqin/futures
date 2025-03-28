package main

import (
	"context"
	"log"
	"time"

	"github.com/neqin/futures/connectors/gateio" // Adjust import path if needed
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Adjusted timeout
	defer cancel()

	// Create a public-only client
	client := gateio.NewPublicOnly(nil)
	// Optional: Set base URL if needed for testing (e.g., testnet)
	// client.SetBaseURL("https://fx-api-testnet.gateio.ws")

	log.Println("Testing All Gate.io Public API Methods (Post-Fix 2)...")

	settle := "usdt"
	contractName := "BTC_USDT" // Example contract for most tests

	// --- Test ListFuturesContracts ---
	log.Println("Fetching USDT contracts...")
	contracts, err := client.ListFuturesContracts(ctx, settle)
	if err != nil {
		log.Printf("ERROR fetching contracts: %v\n", err)
	} else if contracts != nil && len(*contracts) > 0 {
		log.Printf("OK: Fetched %d USDT contracts. First: %s\n", len(*contracts), (*contracts)[0].Name)
	} else {
		log.Println("WARN: Fetched contracts, but the list is empty or nil.")
	}

	// --- Test ListContractStats ---
	log.Printf("Fetching contract stats for %s...\n", contractName)
	statsLimit := 5
	statsInterval := "5m"
	stats, err := client.ListContractStats(ctx, settle, contractName, &statsInterval, &statsLimit, nil, nil)
	if err != nil {
		log.Printf("ERROR fetching contract stats for %s: %v\n", contractName, err)
	} else if stats != nil && len(*stats) > 0 {
		log.Printf("OK: Fetched %d stats entries for %s. First entry time: %d, MarkPrice: %f\n", len(*stats), contractName, (*stats)[0].Time, (*stats)[0].MarkPrice)
	} else {
		log.Printf("WARN: Fetched contract stats for %s, but list is empty or nil.\n", contractName)
	}

	// --- Test ListFuturesOrderBook ---
	log.Printf("Fetching order book for %s...\n", contractName)
	obLimit := 10
	obWithID := true
	obInterval := "0" // No aggregation
	orderBook, err := client.ListFuturesOrderBook(ctx, settle, contractName, &obInterval, &obLimit, &obWithID)
	if err != nil {
		log.Printf("ERROR fetching order book for %s: %v\n", contractName, err)
	} else if orderBook != nil && len(orderBook.Asks) > 0 && len(orderBook.Bids) > 0 {
		log.Printf("OK: Fetched order book for %s. Ask[0]: %d @ %s, Bid[0]: %d @ %s (ID: %d)\n",
			contractName, orderBook.Asks[0].Size, orderBook.Asks[0].Price, orderBook.Bids[0].Size, orderBook.Bids[0].Price, orderBook.ID)
	} else {
		log.Printf("WARN: Fetched order book for %s, but asks or bids are empty or nil.\n", contractName)
	}

	// --- Test ListFuturesTrades ---
	log.Printf("Fetching recent trades for %s...\n", contractName)
	tradesLimit := 5
	trades, err := client.ListFuturesTrades(ctx, settle, contractName, &tradesLimit, nil, nil, nil, nil)
	if err != nil {
		log.Printf("ERROR fetching trades for %s: %v\n", contractName, err)
	} else if trades != nil && len(*trades) > 0 {
		log.Printf("OK: Fetched %d recent trades for %s. First trade ID: %d\n", len(*trades), contractName, (*trades)[0].ID)
	} else {
		log.Printf("WARN: Fetched trades for %s, but list is empty or nil.\n", contractName)
	}

	// --- Test ListFuturesCandlesticks ---
	log.Printf("Fetching candlesticks for %s...\n", contractName)
	candleLimit := 5
	candleInterval := "1m"
	// Expecting []FuturesCandlestick (which is []CandlestickData)
	candles, err := client.ListFuturesCandlesticks(ctx, settle, contractName, &candleLimit, &candleInterval, nil, nil)
	if err != nil {
		log.Printf("ERROR fetching candlesticks for %s: %v\n", contractName, err)
	} else if candles != nil && len(*candles) > 0 {
		// Access fields using dot notation
		log.Printf("OK: Fetched %d candlesticks for %s. First candle timestamp: %s, Open: %s\n", len(*candles), contractName, (*candles)[0].Timestamp, (*candles)[0].Open)
	} else {
		log.Printf("WARN: Fetched candlesticks for %s, but list is empty or nil.\n", contractName)
	}

	// --- Test ListFuturesPremiumIndex ---
	log.Printf("Fetching premium index K-line for %s...\n", contractName)
	premiumLimit := 5
	premiumInterval := "1m"
	// Expecting []FuturesPremiumIndex (which is []PremiumIndexData)
	premiumIndex, err := client.ListFuturesPremiumIndex(ctx, settle, contractName, &premiumLimit, &premiumInterval, nil, nil)
	if err != nil {
		log.Printf("ERROR fetching premium index for %s: %v\n", contractName, err)
	} else if premiumIndex != nil && len(*premiumIndex) > 0 {
		// Access fields using dot notation - using corrected fields
		log.Printf("OK: Fetched %d premium index entries for %s. First entry timestamp: %d, MarkPrice: %f\n", len(*premiumIndex), contractName, (*premiumIndex)[0].Timestamp, (*premiumIndex)[0].MarkPrice)
	} else {
		log.Printf("WARN: Fetched premium index for %s, but list is empty or nil.\n", contractName)
	}

	// --- Test ListFuturesTickers ---
	log.Printf("Fetching USDT tickers...")
	tickers, err := client.ListFuturesTickers(ctx, settle, nil) // Fetch all USDT tickers
	if err != nil {
		log.Printf("ERROR fetching tickers: %v\n", err)
	} else if tickers != nil && len(*tickers) > 0 {
		log.Printf("OK: Fetched %d USDT tickers. First: %s (Last: %s)\n", len(*tickers), (*tickers)[0].Contract, (*tickers)[0].Last)
	} else {
		log.Println("WARN: Fetched tickers, but the list is empty or nil.")
	}

	// --- Test ListFuturesFundingRateHistory ---
	log.Printf("Fetching funding rate history for %s...\n", contractName)
	fundingLimit := 5
	fundingRates, err := client.ListFuturesFundingRateHistory(ctx, settle, contractName, &fundingLimit)
	if err != nil {
		log.Printf("ERROR fetching funding rates for %s: %v\n", contractName, err)
	} else if fundingRates != nil && len(*fundingRates) > 0 {
		log.Printf("OK: Fetched %d funding rate entries for %s. First entry: %s @ %d\n", len(*fundingRates), contractName, (*fundingRates)[0].Rate, (*fundingRates)[0].Timestamp)
	} else {
		log.Printf("WARN: Fetched funding rates for %s, but list is empty or nil.\n", contractName)
	}

	// --- Test ListFuturesInsuranceLedger ---
	log.Printf("Fetching insurance ledger for %s...\n", settle)
	insuranceLimit := 5
	insuranceLedger, err := client.ListFuturesInsuranceLedger(ctx, settle, &insuranceLimit)
	if err != nil {
		log.Printf("ERROR fetching insurance ledger for %s: %v\n", settle, err)
	} else if insuranceLedger != nil && len(*insuranceLedger) > 0 {
		log.Printf("OK: Fetched %d insurance ledger entries for %s. First entry: %s @ %d\n", len(*insuranceLedger), settle, (*insuranceLedger)[0].Change, (*insuranceLedger)[0].Timestamp)
	} else {
		log.Printf("WARN: Fetched insurance ledger for %s, but list is empty or nil.\n", settle)
	}

	// --- Test GetLiquidationHistory ---
	log.Printf("Fetching liquidation history for %s...\n", settle)
	liqLimit := 5
	liqHistory, err := client.GetLiquidationHistory(ctx, settle, nil, &liqLimit, nil, nil, nil) // Fetch recent 5 for settle currency
	if err != nil {
		log.Printf("ERROR fetching liquidation history for %s: %v\n", settle, err)
	} else if liqHistory != nil && len(*liqHistory) > 0 {
		log.Printf("OK: Fetched %d liquidation history entries for %s. First entry OrderID: %d\n", len(*liqHistory), settle, (*liqHistory)[0].OrderID)
	} else {
		log.Printf("WARN: Fetched liquidation history for %s, but list is empty or nil.\n", settle)
	}

	// --- Test GetRiskLimitTiers ---
	log.Printf("Fetching risk limit tiers for %s...\n", contractName)
	riskTiers, err := client.GetRiskLimitTiers(ctx, settle, contractName)
	if err != nil {
		log.Printf("ERROR fetching risk limit tiers for %s: %v\n", contractName, err)
	} else if riskTiers != nil && len(*riskTiers) > 0 {
		log.Printf("OK: Fetched %d risk limit tiers for %s. First tier: %d (Max Leverage: %s)\n", len(*riskTiers), contractName, (*riskTiers)[0].Tier, (*riskTiers)[0].LeverageMax)
	} else {
		log.Printf("WARN: Fetched risk limit tiers for %s, but list is empty or nil.\n", contractName)
	}

	log.Println("Gate.io Public API Method Tests Complete.")
}
