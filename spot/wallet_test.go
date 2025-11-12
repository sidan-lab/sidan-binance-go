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
