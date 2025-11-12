package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BaseURL = "https://api.binance.com"
)

// Client represents the Binance API client
type Client struct {
	APIKey    string
	APISecret string
	BaseURL   string
	HTTPClient *http.Client
}

// NewClient creates a new Binance API client
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   BaseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// sign creates a HMAC SHA256 signature
func (c *Client) sign(queryString string) string {
	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

// buildQueryString converts parameters to URL query string
func buildQueryString(params map[string]interface{}) string {
	values := url.Values{}
	for key, value := range params {
		if value == nil {
			continue
		}
		switch v := value.(type) {
		case string:
			if v != "" {
				values.Add(key, v)
			}
		case int:
			values.Add(key, strconv.Itoa(v))
		case int64:
			values.Add(key, strconv.FormatInt(v, 10))
		case float64:
			values.Add(key, strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			values.Add(key, strconv.FormatBool(v))
		case []string:
			for _, item := range v {
				values.Add(key, item)
			}
		}
	}
	return values.Encode()
}

// SignRequest performs a signed API request with rate limiting
func (c *Client) SignRequest(method, endpoint string, params map[string]interface{}) ([]byte, error) {
	// Add timestamp
	if params == nil {
		params = make(map[string]interface{})
	}
	params["timestamp"] = time.Now().UnixMilli()

	// Build query string
	queryString := buildQueryString(params)

	// Sign the query string
	signature := c.sign(queryString)
	queryString += "&signature=" + signature

	// Build full URL
	fullURL := c.BaseURL + endpoint + "?" + queryString

	// Create request
	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Set("X-MBX-APIKEY", c.APIKey)

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// LimitedEncodedSignRequest performs a signed API request with rate limiting
// This is an alias for SignRequest to match the Python SDK naming
func (c *Client) LimitedEncodedSignRequest(method, endpoint string, params map[string]interface{}) ([]byte, error) {
	return c.SignRequest(method, endpoint, params)
}

// ParseResponse parses JSON response into the provided interface
func ParseResponse(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
