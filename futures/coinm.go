package futures

import (
	"github.com/sidan-lab/sidan-binance-go/client"
)

const CoinMBaseURL = "https://dapi.binance.com"

// CoinMAccountClient handles COIN-M Futures account related API endpoints
type CoinMAccountClient struct {
	*client.Client
}

// NewCoinMAccountClient creates a new COIN-M Futures AccountClient
func NewCoinMAccountClient(apiKey, apiSecret string) *CoinMAccountClient {
	c := client.NewClient(apiKey, apiSecret)
	c.BaseURL = CoinMBaseURL
	return &CoinMAccountClient{Client: c}
}

// Balance queries COIN-M futures account balance (USER_DATA)
//
// Weight(IP): 1
//
// GET /dapi/v1/balance
//
// https://developers.binance.com/docs/derivatives/coin-margined-futures/account/Futures-Account-Balance
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (a *CoinMAccountClient) Balance(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return a.SignRequest("GET", "/dapi/v1/balance", params)
}
