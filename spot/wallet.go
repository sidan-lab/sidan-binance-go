package spot

import (
	"github.com/sidan-lab/sidan-binance-go/client"
	"github.com/sidan-lab/sidan-binance-go/utils"
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

// DepositHistory queries user deposit history (USER_DATA)
//
// Weight(IP): 10
//
// GET /sapi/v1/capital/deposit/hisrec
//
// https://developers.binance.com/docs/wallet/asset/deposit-history
//
// Optional parameters:
//   - coin: Coin symbol
//   - status: 0:pending, 6:credited, 1:success
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - offset: Default 0
//   - limit: Default 1000, max 1000
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) DepositHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/capital/deposit/hisrec", params)
}

// WithdrawalHistory queries user withdrawal history (USER_DATA)
//
// Weight(IP): 10
//
// GET /sapi/v1/capital/withdraw/history
//
// https://developers.binance.com/docs/wallet/asset/withdraw-history
//
// Optional parameters:
//   - coin: Coin symbol
//   - status: 0:Email Sent, 1:Cancelled, 2:Awaiting Approval, 3:Rejected, 4:Processing, 5:Failure, 6:Completed
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - offset: Default 0
//   - limit: Default 1000, max 1000
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) WithdrawalHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/capital/withdraw/history", params)
}

// MyTrades queries account trade history (USER_DATA)
//
// Weight(IP): 10
//
// GET /api/v3/myTrades
//
// https://developers.binance.com/docs/spot/trade/account-trade-list
//
// Parameters:
//   - symbol: Trading pair symbol (required)
//
// Optional parameters:
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - orderId: Order ID
//   - fromId: Trade ID to start from (inclusive)
//   - limit: Default 500, max 1000
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) MyTrades(symbol string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(symbol, "symbol"); err != nil {
		return nil, err
	}
	if params == nil {
		params = make(map[string]interface{})
	}
	params["symbol"] = symbol
	return w.SignRequest("GET", "/api/v3/myTrades", params)
}

// UniversalTransferHistory queries user universal transfer history (USER_DATA)
//
// Weight(IP): 1
//
// GET /sapi/v1/asset/transfer
//
// https://developers.binance.com/docs/wallet/asset/query-user-universal-transfer
//
// Parameters:
//   - type: Transfer type (required). Examples:
//     MAIN_UMFUTURE: Spot to USDⓈ-M Futures
//     MAIN_CMFUTURE: Spot to COIN-M Futures
//     MAIN_MARGIN: Spot to Cross Margin
//     UMFUTURE_MAIN: USDⓈ-M Futures to Spot
//     CMFUTURE_MAIN: COIN-M Futures to Spot
//     MARGIN_MAIN: Cross Margin to Spot
//     MAIN_FUNDING: Spot to Funding
//     FUNDING_MAIN: Funding to Spot
//
// Optional parameters:
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - current: Current page, default 1
//   - size: Page size, default 10, max 100
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) UniversalTransferHistory(transferType string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(transferType, "type"); err != nil {
		return nil, err
	}
	if params == nil {
		params = make(map[string]interface{})
	}
	params["type"] = transferType
	return w.SignRequest("GET", "/sapi/v1/asset/transfer", params)
}

// SubAccountTransferHistory queries sub-account's own transfer history (For Sub-account)
//
// Weight(IP): 1
//
// GET /sapi/v1/sub-account/transfer/subUserHistory
//
// https://developers.binance.com/docs/sub_account/asset-management/Sub-account-Transfer-History
//
// Note: This endpoint is for sub-accounts to query their own transfer history
// with master account or other sub-accounts. For PnL calculations:
//   - type 1: Transfer IN (master/sub-account → this sub-account) = deposit
//   - type 2: Transfer OUT (this sub-account → master) = withdrawal
//
// Optional parameters:
//   - asset: Asset symbol (e.g., "USDT", "BTC")
//   - type: 1: transfer in, 2: transfer out
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - limit: Default 500, max 500
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) SubAccountTransferHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/sub-account/transfer/subUserHistory", params)
}

// AccountSnapshot queries daily account snapshots (USER_DATA)
//
// Weight(IP): 2400
//
// GET /sapi/v1/accountSnapshot
//
// https://developers.binance.com/docs/wallet/account/daily-account-snapshot
//
// Note: This endpoint is useful for tracking historical balance changes,
// including deposits/withdrawals that may not appear in transfer history APIs.
// The snapshot is taken daily and shows balances at the end of each day.
//
// Parameters:
//   - type: Account type (required). Values: "SPOT", "MARGIN", "FUTURES"
//
// Optional parameters:
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - limit: Default 7, max 30
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) AccountSnapshot(accountType string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(accountType, "type"); err != nil {
		return nil, err
	}
	if params == nil {
		params = make(map[string]interface{})
	}
	params["type"] = accountType
	return w.SignRequest("GET", "/sapi/v1/accountSnapshot", params)
}

// MasterSubAccountTransferHistory queries sub-account transfer history (For Master Account)
//
// Weight(IP): 1
//
// GET /sapi/v1/sub-account/universalTransfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Universal-Transfer-History
//
// Note: This endpoint is for MASTER accounts to query transfer history
// with their sub-accounts. Use this for more complete transfer records
// compared to SubAccountTransferHistory.
//
// Optional parameters:
//   - fromEmail: Sub-account email
//   - toEmail: Sub-account email
//   - clientTranId: Client transfer ID
//   - startTime: Start time in milliseconds
//   - endTime: End time in milliseconds
//   - page: Default 1
//   - limit: Default 500, max 500
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) MasterSubAccountTransferHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/sub-account/universalTransfer", params)
}

// MasterSubAccountList queries sub-account list (For Master Account)
//
// Weight(IP): 1
//
// GET /sapi/v1/sub-account/list
//
// https://developers.binance.com/docs/sub_account/account-management/Query-Sub-account-List
//
// Note: This endpoint is for MASTER accounts to list their sub-accounts.
//
// Optional parameters:
//   - email: Sub-account email
//   - isFreeze: true or false
//   - page: Default 1
//   - limit: Default 1, max 200
//   - recvWindow: The value cannot be greater than 60000
func (w *WalletClient) MasterSubAccountList(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return w.SignRequest("GET", "/sapi/v1/sub-account/list", params)
}
