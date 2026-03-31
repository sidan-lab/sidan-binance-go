package futures

import (
	"github.com/sidan-lab/sidan-binance-go/client"
)

const FuturesBaseURL = "https://fapi.binance.com"

// AccountClient handles USDⓈ-M Futures account related API endpoints
type AccountClient struct {
	*client.Client
}

// NewAccountClient creates a new Futures AccountClient
func NewAccountClient(apiKey, apiSecret string) *AccountClient {
	c := client.NewClient(apiKey, apiSecret)
	c.BaseURL = FuturesBaseURL
	return &AccountClient{Client: c}
}

// Balance queries futures account balance (USER_DATA)
//
// Weight(IP): 5
//
// GET /fapi/v3/balance
//
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V3
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (a *AccountClient) Balance(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return a.SignRequest("GET", "/fapi/v3/balance", params)
}
