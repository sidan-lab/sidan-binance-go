package spot

import (
	"github.com/sidan-lab/sidan-binance-go/client"
)

// WalletClient handles wallet related API endpoints
type WalletClient struct {
	*client.Client
}

// NewWalletClient creates a new WalletClient
func NewWalletClient(apiKey, apiSecret string) *WalletClient {
	return &WalletClient{
		Client: client.NewClient(apiKey, apiSecret),
	}
}

// Balance queries user wallet balance (USER_DATA)
//
// Weight(IP): 60
//
// GET /sapi/v1/asset/wallet/balance
//
// https://developers.binance.com/docs/wallet/asset/query-user-wallet-balance
//
// Optional parameters:
//   - quoteAsset: Currency for balance valuation (e.g., "BTC", "USDT", "ETH", "USDC", "BNB"). Default: "BTC"
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) Balance(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/asset/wallet/balance", params)
}

// UserAsset gets user assets, just for positive data (USER_DATA)
//
// Weight(IP): 5
//
// POST /sapi/v3/asset/getUserAsset
//
// https://developers.binance.com/docs/wallet/asset/user-assets
//
// Optional parameters:
//   - asset: If asset is blank, then query all positive assets user have
//   - needBtcValuation: true or false
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) UserAsset(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("POST", "/sapi/v3/asset/getUserAsset", params)
}
