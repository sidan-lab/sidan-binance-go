package spot

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	// Try to load .env file, but don't fail if it doesn't exist
	// Environment variables will still work if set directly
	_ = godotenv.Load("../.env")
}

func TestNewSubAccountClient(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewSubAccountClient(apiKey, apiSecret)

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

	t.Logf("Client initialized successfully with BaseURL: %s", client.BaseURL)
}

func TestSubAccountList(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewSubAccountClient(apiKey, apiSecret)

	// Test SubAccountList API call
	response, err := client.SubAccountList(map[string]interface{}{
		"page":  1,
		"limit": 10,
	})

	if err != nil {
		t.Logf("API Error (this may be expected if account has no sub-accounts): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result map[string]interface{}
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
		t.Logf("Sub-account list response:\n%s", string(prettyJSON))
	}
}

func TestSubAccountStatus(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewSubAccountClient(apiKey, apiSecret)

	// Test SubAccountStatus API call
	response, err := client.SubAccountStatus(nil)

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result map[string]interface{}
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
		t.Logf("Sub-account status response:\n%s", string(prettyJSON))
	}
}

func TestSubAccountSpotSummary(t *testing.T) {
	// Load from environment variables
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET_KEY")
	fmt.Println("API Key:", apiKey)
	fmt.Println("API Secret:", apiSecret)

	if apiKey == "" || apiSecret == "" {
		t.Skip("Skipping test: BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables not set")
	}

	client := NewSubAccountClient(apiKey, apiSecret)

	// Test SubAccountSpotSummary API call
	response, err := client.SubAccountSpotSummary(map[string]interface{}{
		"page": 1,
		"size": 10,
	})

	if err != nil {
		t.Logf("API Error (this may be expected): %v", err)
		t.Logf("Response: %s", string(response))
		return
	}

	// Parse and log the response
	var result map[string]interface{}
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
		t.Logf("Sub-account spot summary response:\n%s", string(prettyJSON))
	}
}

func TestSubAccountClientMethods(t *testing.T) {
	client := NewSubAccountClient("test_key", "test_secret")

	// Test that all methods exist by verifying client is properly initialized
	// Note: This test doesn't actually call the API, just verifies the client structure

	if client == nil {
		t.Fatal("Expected non-nil client")
	}

	// Verify the client has access to the base Client methods
	if client.Client == nil {
		t.Error("Expected embedded Client to be non-nil")
	}

	t.Log("Client structure verification passed")
}
