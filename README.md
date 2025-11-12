# Sidan Binance Go SDK

A Go SDK for Binance API, providing comprehensive support for Binance sub-account management operations.

## Features

- Full implementation of Binance Sub-Account API
- HMAC SHA256 request signing
- Type-safe API client
- Parameter validation
- Clean and idiomatic Go code

## Installation

```bash
go get github.com/sidan-lab/sidan-binance-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/sidan-lab/sidan-binance-go/spot"
)

func main() {
    // Initialize the client with your API credentials
    client := spot.NewSubAccountClient("YOUR_API_KEY", "YOUR_API_SECRET")

    // Example: Query sub-account list
    response, err := client.SubAccountList(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(response))
}
```

## API Coverage

This SDK currently implements the complete Binance Sub-Account API as defined in the [Python reference implementation](https://github.com/sidan-lab/sidan-binance-py/blob/main/binance/spot/_sub_account.py).

### Sub-Account Management

- `SubAccountCreate` - Create a virtual sub-account
- `SubAccountList` - Query sub-account list
- `SubAccountStatus` - Get sub-account's status on Margin/Futures
- `SubAccountEnableMargin` - Enable margin for sub-account
- `SubAccountEnableFutures` - Enable futures for sub-account
- `SubAccountEnableLeverageToken` - Enable leverage token for sub-account
- `EnableOptionsForSubAccount` - Enable options for sub-account

### Asset Management

- `SubAccountAssets` - Query sub-account assets (V3)
- `QuerySubAccountAssets` - Query sub-account assets (V4)
- `SubAccountDepositAddress` - Get sub-account deposit address
- `SubAccountDepositHistory` - Get sub-account deposit history
- `SubAccountMarginAccount` - Get detail on sub-account's margin account
- `SubAccountMarginAccountSummary` - Get summary of sub-account's margin account
- `SubAccountFuturesAccount` - Get detail on sub-account's futures account (V2)
- `SubAccountFuturesAccountSummary` - Get summary of sub-account's futures account (V2)
- `DetailOnSubAccountSFuturesAccount` - Get detail on sub-account's futures account
- `SummaryOfSubAccountSFuturesAccount` - Get summary of sub-account's futures account
- `FuturesPositionRiskOfSubAccount` - Get futures position-risk of sub-account
- `SubAccountFuturesPositionRisk` - Get futures position-risk of sub-account (V2)
- `SubAccountSpotSummary` - Query sub-account spot assets summary

### Transfer Operations

- `SubAccountFuturesTransfer` - Futures transfer for sub-account
- `SubAccountMarginTransfer` - Margin transfer for sub-account
- `SubAccountTransferToSub` - Transfer to sub-account of same master
- `SubAccountTransferToMaster` - Transfer to master account
- `SubAccountUniversalTransfer` - Universal transfer
- `SubAccountUniversalTransferHistory` - Query universal transfer history
- `SubAccountFuturesAssetTransfer` - Sub-account futures asset transfer
- `SubAccountFuturesAssetTransferHistory` - Query sub-account futures asset transfer history
- `SubAccountTransferSubAccountHistory` - Sub-account transfer history
- `SubAccountSpotTransferHistory` - Query sub-account spot asset transfer history

### Managed Sub-Accounts

- `ManagedSubAccountDeposit` - Deposit assets into managed sub-account
- `ManagedSubAccountWithdraw` - Withdraw assets from managed sub-account
- `ManagedSubAccountAssets` - Query managed sub-account asset details
- `ManagedSubAccountGetSnapshot` - Query managed sub-account snapshot
- `ManagedSubAccountDepositAddress` - Get managed sub-account deposit address
- `ManagedSubAccountInvestorTransLog` - Query managed sub-account transfer log (Investor)
- `ManagedSubAccountTradingTransLog` - Query managed sub-account transfer log (Trading Team)
- `QueryManagedSubAccountList` - Query managed sub-account list
- `QueryManagedSubAccountMarginAssetDetails` - Query managed sub-account margin asset details
- `QueryManagedSubAccountFuturesAssetDetails` - Query managed sub-account futures asset details
- `QueryManagedSubAccountTransferLog` - Query managed sub-account transfer log

### API Management

- `SubAccountUpdateIPRestriction` - Update IP restriction for sub-account API key
- `SubAccountAPIGetIPRestriction` - Get IP restriction for sub-account API key
- `SubAccountAPIDeleteIP` - Delete IP list for sub-account API key

### Statistics

- `QuerySubAccountTransactionStatistics` - Query sub-account transaction statistics

## Usage Examples

### Create a Virtual Sub-Account

```go
client := spot.NewSubAccountClient("API_KEY", "API_SECRET")

response, err := client.SubAccountCreate("mysubaccount", map[string]interface{}{
    "recvWindow": 5000,
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(response))
```

### Query Sub-Account Assets

```go
response, err := client.SubAccountAssets("sub@account.com", nil)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(response))
```

### Perform Universal Transfer

```go
response, err := client.SubAccountUniversalTransfer(
    "SPOT",           // fromAccountType
    "USDT_FUTURE",    // toAccountType
    "USDT",           // asset
    100.0,            // amount
    map[string]interface{}{
        "fromEmail": "from@account.com",
        "toEmail":   "to@account.com",
    },
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(response))
```

### Enable Margin for Sub-Account

```go
response, err := client.SubAccountEnableMargin("sub@account.com", nil)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(response))
```

### Query Managed Sub-Account List

```go
response, err := client.QueryManagedSubAccountList(map[string]interface{}{
    "page":  1,
    "limit": 20,
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(response))
```

## Parameter Handling

All API methods accept optional parameters through a `map[string]interface{}` parameter. You can pass `nil` if no optional parameters are needed.

```go
// No optional parameters
response, err := client.SubAccountList(nil)

// With optional parameters
response, err := client.SubAccountList(map[string]interface{}{
    "email":    "sub@account.com",
    "page":     1,
    "limit":    10,
    "isFreeze": "false",
})
```

## Error Handling

The SDK includes comprehensive parameter validation. Required parameters are validated before making API requests:

```go
response, err := client.SubAccountAssets("", nil)
if err != nil {
    // Error: required parameter email is empty
    log.Fatal(err)
}
```

## Testing

The SDK includes comprehensive tests that can be run against the actual Binance API.

### Setting Up Credentials

You can provide API credentials in two ways:

**Option 1: Using a .env file (Recommended)**

```bash
# Copy the example file
cp .env.example .env

# Edit .env and add your credentials
# BINANCE_API_KEY=your_api_key
# BINANCE_SECRET_KEY=your_api_secret
```

**Option 2: Using environment variables**

```bash
export BINANCE_API_KEY="your_api_key"
export BINANCE_SECRET_KEY="your_api_secret"
```

### Running Tests

```bash
# Run all tests with verbose output
go test ./spot -v

# Run a specific test
go test ./spot -v -run TestSubAccountList
```

Tests will automatically skip if credentials are not set. The test suite includes:

- `TestNewSubAccountClient` - Verifies client initialization with real credentials
- `TestSubAccountList` - Tests querying sub-account list and logs the response
- `TestSubAccountStatus` - Tests querying sub-account status and logs the response
- `TestSubAccountSpotSummary` - Tests querying spot summary and logs the response
- `TestSubAccountClientMethods` - Verifies client structure (no API call required)

All API tests log the full response in pretty-printed JSON format for easy debugging.

## API Documentation

For detailed information about each endpoint, parameters, and responses, please refer to the [official Binance API documentation](https://developers.binance.com/docs/sub_account).

## Project Structure

```
sidan-binance-go/
├── client/          # Core HTTP client with request signing
│   └── client.go
├── spot/            # Spot trading endpoints
│   └── sub_account.go
├── utils/           # Utility functions
│   └── validation.go
├── examples/        # Usage examples
│   └── sub_account_example.go
├── go.mod
├── .gitignore
└── README.md
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Related Projects

- [sidan-binance-py](https://github.com/sidan-lab/sidan-binance-py) - Python implementation (reference)

## Support

For issues and questions, please use the GitHub issue tracker.
