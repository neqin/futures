# XT.com Futures Connector (`connectors/xt`)

This directory contains the Go client implementation for interacting with the XT.com Futures API (v1).

**Note:** This implementation is based on the provided examples (`xt.txt`) and documentation (`xt2.txt`). Some details, especially around response structures and specific parameter requirements (like content-type for POST requests), might need further verification against the live API.

## Structure

-   `xt.go`: Provides helper functions (`New`, `NewPublicOnly`) to create client instances.
-   `client.go`: Contains the core `Client` struct, authentication logic (signature generation based on `xt2.txt`), and request sending methods. Handles `application/x-www-form-urlencoded` and `application/json` request bodies.
-   `types.go`: Defines Go structs corresponding to the JSON data structures returned by the API endpoints.
-   `market_public.go`: Implements public API methods related to market data (symbols, tickers, k-lines, depth, etc.). These do not require API keys.
-   `account_private.go`: Implements private API methods related to user account details, balances, positions, and history. Requires API keys.
-   `trading_private.go`: Implements private API methods related to placing and managing orders (spot, trigger, stop-limit, track). Requires API keys.

## Installation

Assuming your project uses Go modules, the connector can be used directly via its import path within the project (e.g., `github.com/neqin/futures/connectors/xt`).

## Usage

### Creating a Client

**1. Public-Only Client:**

```go
package main

import (
	"context"
	"log"

	"github.com/neqin/futures/connectors/xt"
)

func main() {
	publicClient := xt.NewPublicOnly(nil)

	// Example: Get Server Time
	serverTime, err := publicClient.GetServerTime(context.Background())
	if err != nil {
		log.Fatalf("Error getting server time: %v", err)
	}
	log.Printf("XT Server Time (ms): %d", serverTime.Result)

	// Example: Get BTC_USDT Ticker
	ticker, err := publicClient.GetMarketTicker(context.Background(), "btc_usdt")
	if err != nil {
		log.Fatalf("Error getting ticker: %v", err)
	}
	log.Printf("BTC_USDT Last Price: %s", ticker.Result.Close)
}
```

**2. Private Client (Requires API Keys):**

Ensure API keys are available as environment variables (e.g., loaded from `.env.local` using `godotenv`). The expected variable names are `XT_API_KEY` and `XT_API_SECRET`.

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neqin/futures/connectors/xt"
)

func main() {
	// Load API keys
	err := godotenv.Load(".env.local") // Adjust path if needed
	if err != nil {
		log.Println("Warning: Could not load .env.local file:", err)
	}

	apiKey := os.Getenv("XT_API_KEY") // Make sure these match your .env file
	secretKey := os.Getenv("XT_API_SECRET")

	if apiKey == "" || secretKey == "" {
		log.Fatal("Error: XT_API_KEY or XT_API_SECRET not found in environment.")
	}

	// Create a client with API keys
	privateClient := xt.New(apiKey, secretKey, nil)

	// Example: Get Account Info
	accountInfo, err := privateClient.GetAccountInfo(context.Background())
	if err != nil {
		log.Fatalf("Error getting account info: %v", err)
	}
	log.Printf("XT Account User ID: %d, Allow Trade: %t", accountInfo.Result.UserID, accountInfo.Result.AllowTrade)

	// Example: Get USDT Balance
	balance, err := privateClient.GetBalance(context.Background(), "usdt")
	if err != nil {
		log.Fatalf("Error getting USDT balance: %v", err)
	}
	log.Printf("USDT Wallet Balance: %s", balance.Result.WalletBalance)
}
```

### Available Methods

Refer to the method definitions and comments within:
- `market_public.go`
- `account_private.go`
- `trading_private.go`

### Signature Generation

The signature generation follows the process described in `xt2.txt`:
1.  Headers `validate-appkey` and `validate-timestamp` are combined.
2.  Path, sorted query parameters (for GET/DELETE), and request body (sorted form data or raw JSON string for POST/PUT) are combined.
3.  The header string and data string are concatenated.
4.  The final string is signed using HMAC-SHA256 with the `secretKey`.
5.  The signature is added to the `validate-signature` header.

### Content Types

-   GET/DELETE requests do not typically have a Content-Type.
-   POST/PUT requests default to `application/json` if the `bodyParams` argument to `SendPrivateRequest` is a struct or map (excluding `map[string]string`).
-   If `bodyParams` is specifically `map[string]string`, it's treated as form data, sorted, encoded, and sent with `Content-Type: application/x-www-form-urlencoded`. This matches the requirement for endpoints like batch order creation.

## Testing

No dedicated test script is provided yet for XT.com. You can adapt the `cmd/gateio_test/main.go` script or create a new one (`cmd/xt_test/main.go`) to test the implemented methods. Remember to set `XT_API_KEY` and `XT_API_SECRET` environment variables or in your `.env.local` file.