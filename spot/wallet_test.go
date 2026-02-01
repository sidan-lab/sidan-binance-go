package spot

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	// Try to load .env file, but don't fail if it doesn't exist
	// Environment variables will still work if set directly
	_ = godotenv.Load("../.env")
}

func TestNewWalletClient(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	if client == nil {
		t.Fatal("Expected non-nil client")
	}

	if client.APIKey != apiKey {
		t.Errorf("Expected APIKey %s, got %s", apiKey, client.APIKey)
	}

	if client.APISecret != apiSecret {
		t.Errorf("Expected APISecret %s, got %s", apiSecret, client.APISecret)
	}

	if client.BaseURL == "" {
		t.Error("Expected non-empty BaseURL")
	}

	if client.HTTPClient == nil {
		t.Error("Expected non-nil HTTPClient")
	}

	t.Logf("WalletClient initialized successfully with BaseURL: %s", client.BaseURL)
}

func TestBalance(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test Balance API call
	response, err := client.Balance(nil)

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("Wallet balance response:\n%s", string(prettyJSON))
	}
}

func TestBalanceWithRecvWindow(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test Balance API call with recvWindow parameter
	response, err := client.Balance(map[string]interface{}{
		"recvWindow": 5000,
	})

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("Wallet balance response (with recvWindow):\n%s", string(prettyJSON))
	}
}

func TestBalanceWithQuoteAsset(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test Balance API call with quoteAsset parameter
	response, err := client.Balance(map[string]interface{}{
		"quoteAsset": "USDT",
	})

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("Wallet balance response (USDT valuation):\n%s", string(prettyJSON))
	}
}

func TestUserAsset(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test UserAsset API call without parameters (get all positive assets)
	response, err := client.UserAsset(nil)

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("User assets response:\n%s", string(prettyJSON))
	}
}

func TestUserAssetWithAsset(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test UserAsset API call with specific asset
	response, err := client.UserAsset(map[string]interface{}{
		"asset": "BTC",
	})

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("User assets response (BTC):\n%s", string(prettyJSON))
	}
}

func TestUserAssetWithBtcValuation(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test UserAsset API call with BTC valuation
	response, err := client.UserAsset(map[string]interface{}{
		"needBtcValuation": "true",
	})

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Logf("Result: %+v", result)
	} else {
		t.Logf("User assets response (with BTC valuation):\n%s", string(prettyJSON))
	}
}

func TestWalletClientMethods(t *testing.T) {
	client := NewWalletClient("test_key", "test_secret")

	// Test that all methods exist by verifying client is properly initialized
	// Note: This test doesn't actually call the API, just verifies the client structure

	if client == nil {
		t.Fatal("Expected non-nil client")
	}

	// Verify the client has access to the base Client methods
	if client.Client == nil {
		t.Error("Expected embedded Client to be non-nil")
	}

	t.Log("WalletClient structure verification passed")
}

func TestDepositHistory(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.DepositHistory(nil)
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Deposit history response:\n%s", string(prettyJSON))
}

func TestDepositHistoryWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.DepositHistory(map[string]interface{}{
		"status": 1,
		"limit":  10,
	})
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Deposit history (with params) response:\n%s", string(prettyJSON))
}

func TestWithdrawalHistory(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.WithdrawalHistory(nil)
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Withdrawal history response:\n%s", string(prettyJSON))
}

func TestWithdrawalHistoryWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.WithdrawalHistory(map[string]interface{}{
		"status": 6,
		"limit":  10,
	})
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Withdrawal history (with params) response:\n%s", string(prettyJSON))
}

func TestMyTrades(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.MyTrades("BTCUSDT", nil)
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("My trades response:\n%s", string(prettyJSON))
}

func TestMyTradesWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.MyTrades("BTCUSDT", map[string]interface{}{
		"limit": 10,
	})
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("My trades (with params) response:\n%s", string(prettyJSON))
}

func TestMyTradesRequiresSymbol(t *testing.T) {
	client := NewWalletClient("test_key", "test_secret")

	_, err := client.MyTrades("", nil)
	if err == nil {
		t.Error("Expected error for empty symbol, got nil")
	}
	t.Logf("Correctly returned error for empty symbol: %v", err)
}

func TestUniversalTransferHistory(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test transfer from Futures to Spot (transfer in)
	response, err := client.UniversalTransferHistory("UMFUTURE_MAIN", nil)
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Universal transfer history (UMFUTURE_MAIN) response:\n%s", string(prettyJSON))
}

func TestUniversalTransferHistoryWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Test transfer from Spot to Futures (transfer out) with params
	response, err := client.UniversalTransferHistory("MAIN_UMFUTURE", map[string]interface{}{
		"size": 10,
	})
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Universal transfer history (MAIN_UMFUTURE with params) response:\n%s", string(prettyJSON))
}

func TestUniversalTransferHistoryRequiresType(t *testing.T) {
	client := NewWalletClient("test_key", "test_secret")

	_, err := client.UniversalTransferHistory("", nil)
	if err == nil {
		t.Error("Expected error for empty type, got nil")
	}
	t.Logf("Correctly returned error for empty type: %v", err)
}

func TestSubAccountTransferHistory(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Get all transfers
	response, err := client.SubAccountTransferHistory(nil)
	if err != nil {
		t.Logf("API Error (this may be expected for non-sub-accounts): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Sub-account transfer history (%d records):\n%s", len(result), string(prettyJSON))
}

func TestSubAccountTransferHistoryWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	// Get only withdrawals (type=2 = transfers OUT to master)
	response, err := client.SubAccountTransferHistory(map[string]interface{}{
		"type":  2,
		"limit": 10,
	})
	if err != nil {
		t.Logf("API Error (this may be expected for non-sub-accounts): %v", err)
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Sub-account withdrawals (type=2, %d records):\n%s", len(result), string(prettyJSON))
}

func TestAccountSnapshot(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.AccountSnapshot("SPOT", nil)
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		t.Logf("Raw response: %s", string(response))
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Account snapshot response:\n%s", string(prettyJSON))
}

func TestAccountSnapshotWithParams(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewWalletClient(apiKey, apiSecret)

	response, err := client.AccountSnapshot("SPOT", map[string]interface{}{
		"limit": 30,
	})
	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Errorf("Failed to parse response: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Account snapshot (30 days) response:\n%s", string(prettyJSON))
}

func TestAccountSnapshotRequiresType(t *testing.T) {
	client := NewWalletClient("test_key", "test_secret")

	_, err := client.AccountSnapshot("", nil)
	if err == nil {
		t.Error("Expected error for empty type, got nil")
	}
	t.Logf("Correctly returned error for empty type: %v", err)
}
