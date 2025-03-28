# Gate.io Futures Connector (`connectors/gateio`)

This directory contains the Go client implementation for interacting with the Gate.io Futures API (v4).

## Structure

-   `gateio.go`: Provides helper functions (`New`, `NewPublicOnly`) to create client instances.
-   `client.go`: Contains the core `Client` struct, authentication logic (signature generation), and request sending methods.
-   `types.go`: Defines Go structs corresponding to the JSON data structures returned by the API endpoints.
-   `market_public.go`: Implements public API methods related to market data (contracts, order book, tickers, k-lines, etc.). These do not require API keys.
-   `account_private.go`: Implements private API methods related to user account details, positions, and history. Requires API keys.
-   `trading_private.go`: Implements private API methods related to placing and managing orders. Requires API keys.

## Installation

Assuming your project uses Go modules, the connector can be used directly via its import path within the project (e.g., `github.com/neqin/futures/connectors/gateio`).

## Usage

### Creating a Client

**1. Public-Only Client:**
For accessing only public market data endpoints that don't require authentication.

```go
package main

import (
	"context"
	"log"

	"github.com/neqin/futures/connectors/gateio"
)

func main() {
	// Create a client using the default HTTP client
	publicClient := gateio.NewPublicOnly(nil)

	// Example: List USDT contracts
	contracts, err := publicClient.ListFuturesContracts(context.Background(), "usdt")
	if err != nil {
		log.Fatalf("Error fetching contracts: %v", err)
	}
	log.Printf("Fetched %d USDT contracts.", len(*contracts))
}

```

**2. Private Client (Requires API Keys):**
For accessing private endpoints related to your account, positions, and trading.

Create a `.env.local` file in your project root (or ensure environment variables are set):

```dotenv
GATE_API_KEY=YOUR_API_KEY_HERE
GATE_API_SECRET=YOUR_API_SECRET_HERE
```

Then, load these keys and create the client:

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv" // go get github.com/joho/godotenv
	"github.com/neqin/futures/connectors/gateio"
)

func main() {
	// Load API keys from .env.local
	err := godotenv.Load(".env.local") // Or load from root: godotenv.Load("../.env.local") if running from cmd/
	if err != nil {
		log.Println("Warning: Could not load .env.local file:", err)
	}

	apiKey := os.Getenv("GATE_API_KEY")
	secretKey := os.Getenv("GATE_API_SECRET")

	if apiKey == "" || secretKey == "" {
		log.Fatal("Error: GATE_API_KEY or GATE_API_SECRET not found in environment.")
	}

	// Create a client with API keys
	privateClient := gateio.New(apiKey, secretKey, nil)

	// Example: Get USDT Futures Account Details
	account, err := privateClient.GetFuturesAccount(context.Background(), "usdt")
	if err != nil {
		log.Fatalf("Error fetching account details: %v", err)
	}
	log.Printf("Account User ID: %d, Total Balance: %s USDT", account.User, account.Total)

	// Example: List Open Positions
	positions, err := privateClient.ListPositions(context.Background(), "usdt", nil)
	if err != nil {
		log.Fatalf("Error fetching positions: %v", err)
	}
	log.Printf("Found %d open positions.", len(*positions))
}
```

### Available Methods

Refer to the method definitions and comments within:
- `market_public.go` for public endpoints.
- `account_private.go` for private account/position endpoints.
- `trading_private.go` for private order/trade endpoints.

**Example (Public): Get Order Book**

```go
	contract := "BTC_USDT"
	limit := 5
	interval := "0" // No aggregation
	withID := true
	orderBook, err := publicClient.ListFuturesOrderBook(context.Background(), "usdt", contract, &interval, &limit, &withID)
	// ... handle error and use orderBook ...
```

**Example (Private): Get Open Orders**

```go
	contract := "BTC_USDT"
	openOrders, err := privateClient.ListFuturesOrders(context.Background(), "usdt", "open", &contract, nil, nil, nil, nil, nil)
	// ... handle error and use openOrders ...
```

**Note:** Be cautious when calling private methods that modify state (e.g., `CreateFuturesOrder`, `CancelFuturesOrder`, `UpdatePositionMargin`). Ensure you understand the parameters and consequences.

## Testing

A test script is available at `cmd/gateio_test/main.go`. It demonstrates usage of both public and private methods. To run it (from the project root):

1.  Ensure `GATE_API_KEY` and `GATE_API_SECRET` are set in `.env.local`.
2.  Run: `go run ./cmd/gateio_test/main.go`