package spot

import (
	"github.com/sidan-lab/sidan-binance-go/client"
	"github.com/sidan-lab/sidan-binance-go/utils"
)

// SubAccountClient handles sub-account related API endpoints
type SubAccountClient struct {
	*client.Client
}

// NewSubAccountClient creates a new SubAccountClient
func NewSubAccountClient(apiKey, apiSecret string) *SubAccountClient {
	return &SubAccountClient{
		Client: client.NewClient(apiKey, apiSecret),
	}
}

// SubAccountCreate creates a virtual sub-account (For Master Account)
// Generate a virtual sub account under the master account
//
// POST /sapi/v1/sub-account/virtualSubAccount
//
// https://developers.binance.com/docs/sub_account/account-management/Create-a-Virtual-Sub-account
//
// Parameters:
//   - subAccountString: Please input a string. We will create a virtual email using that string for you to register
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountCreate(subAccountString string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(subAccountString, "subAccountString"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["subAccountString"] = subAccountString

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/virtualSubAccount", params)
}

// SubAccountList queries sub-account list (For Master Account)
// Fetch sub account list.
//
// GET /sapi/v1/sub-account/list
//
// https://developers.binance.com/docs/sub_account/account-management/Query-Sub-account-List
//
// Optional parameters:
//   - email: Sub-account email
//   - isFreeze: true or false
//   - page: default 1
//   - limit: default 10, max 200
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountList(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/list", params)
}

// SubAccountAssets queries sub-account assets (For Master Account)
// Fetch sub-account assets
//
// GET /sapi/v3/sub-account/assets
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Assets-V3
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountAssets(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("GET", "/sapi/v3/sub-account/assets", params)
}

// SubAccountDepositAddress gets sub-account deposit address (For Master Account)
// Fetch sub-account deposit address
//
// GET /sapi/v1/capital/deposit/subAddress
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Sub-account-Deposit-Address
//
// Parameters:
//   - email: Sub-account email
//   - coin: Coin symbol
//
// Optional parameters:
//   - network: Network
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountDepositAddress(email, coin string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email": email,
		"coin":  coin,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["coin"] = coin

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/capital/deposit/subAddress", params)
}

// SubAccountDepositHistory gets sub-account deposit history (For Master Account)
// Fetch sub-account deposit history
//
// GET /sapi/v1/capital/deposit/subHisrec
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Sub-account-Deposit-Address
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - coin: Coin symbol
//   - status: Default 0 (0:pending, 6: credited but cannot withdraw, 1:success)
//   - startTime: Start time
//   - endTime: End time
//   - limit: Limit
//   - offset: Default 0
//   - txId: Transaction ID
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountDepositHistory(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/capital/deposit/subHisrec", params)
}

// SubAccountStatus gets sub-account's status on Margin/Futures (For Master Account)
//
// GET /sapi/v1/sub-account/status
//
// https://developers.binance.com/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures
//
// Optional parameters:
//   - email: Sub-account email
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountStatus(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/status", params)
}

// SubAccountEnableMargin enables margin for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/margin/enable
//
// https://developers.binance.com/docs/sub_account/account-management/Enable-Margin-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountEnableMargin(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/margin/enable", params)
}

// SubAccountMarginAccount gets detail on sub-account's margin account (For Master Account)
//
// GET /sapi/v1/sub-account/margin/account
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountMarginAccount(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/margin/account", params)
}

// SubAccountMarginAccountSummary gets summary of sub-account's margin account (For Master Account)
//
// GET /sapi/v1/sub-account/margin/accountSummary
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Margin-Account
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountMarginAccountSummary(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.SignRequest("GET", "/sapi/v1/sub-account/margin/accountSummary", params)
}

// SubAccountEnableFutures enables futures for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/futures/enable
//
// https://developers.binance.com/docs/sub_account/account-management/Enable-Futures-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountEnableFutures(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/futures/enable", params)
}

// SubAccountFuturesTransfer performs futures transfer for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/futures/transfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Futures-Transfer-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//   - asset: Asset symbol
//   - amount: Amount
//   - transferType: Transfer type
func (s *SubAccountClient) SubAccountFuturesTransfer(email, asset string, amount float64, transferType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":  email,
		"asset":  asset,
		"amount": amount,
		"type":   transferType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["asset"] = asset
	params["amount"] = amount
	params["type"] = transferType

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/futures/transfer", params)
}

// SubAccountMarginTransfer performs margin transfer for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/margin/transfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Margin-Transfer-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//   - asset: Asset symbol
//   - amount: Amount
//   - transferType: Transfer type
func (s *SubAccountClient) SubAccountMarginTransfer(email, asset string, amount float64, transferType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":  email,
		"asset":  asset,
		"amount": amount,
		"type":   transferType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["asset"] = asset
	params["amount"] = amount
	params["type"] = transferType

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/margin/transfer", params)
}

// SubAccountTransferToSub transfers to sub-account of same master (For Sub-account)
//
// POST /sapi/v1/sub-account/transfer/subToSub
//
// https://developers.binance.com/docs/sub_account/asset-management/Transfer-to-Sub-account-of-Same-Master
//
// Parameters:
//   - toEmail: Recipient email
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountTransferToSub(toEmail, asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"toEmail": toEmail,
		"asset":   asset,
		"amount":  amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["toEmail"] = toEmail
	params["asset"] = asset
	params["amount"] = amount

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/transfer/subToSub", params)
}

// SubAccountTransferToMaster transfers to master (For Sub-account)
//
// POST /sapi/v1/sub-account/transfer/subToMaster
//
// https://developers.binance.com/docs/sub_account/asset-management/Transfer-to-Master
//
// Parameters:
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountTransferToMaster(asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"asset":  asset,
		"amount": amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["asset"] = asset
	params["amount"] = amount

	return s.SignRequest("POST", "/sapi/v1/sub-account/transfer/subToMaster", params)
}

// SubAccountTransferSubAccountHistory gets sub-account transfer history (For Sub-account)
//
// GET /sapi/v1/sub-account/transfer/subUserHistory
//
// https://developers.binance.com/docs/sub_account/asset-management/Sub-account-Transfer-History
//
// Optional parameters:
//   - asset: If not sent, result of all assets will be returned
//   - transferType: 1: transfer in, 2: transfer out
//   - startTime: Start time
//   - endTime: End time
//   - limit: Default 500
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountTransferSubAccountHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.SignRequest("GET", "/sapi/v1/sub-account/transfer/subUserHistory", params)
}

// SubAccountFuturesAssetTransferHistory queries sub-account futures asset transfer history (For Master Account)
//
// GET /sapi/v1/sub-account/futures/internalTransfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Futures-Asset-Transfer-History
//
// Parameters:
//   - email: Sub-account email
//   - futuresType: 1: USDT-margined Futures, 2: Coin-margined Futures
//
// Optional parameters:
//   - startTime: Default return the history within 100 days
//   - endTime: Default return the history within 100 days
//   - page: Default value: 1
//   - limit: Default value: 50, Max value: 500
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountFuturesAssetTransferHistory(email string, futuresType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":       email,
		"futuresType": futuresType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["futuresType"] = futuresType

	return s.SignRequest("GET", "/sapi/v1/sub-account/futures/internalTransfer", params)
}

// SubAccountFuturesAssetTransfer performs sub-account futures asset transfer (For Master Account)
//
// POST /sapi/v1/sub-account/futures/internalTransfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Sub-account-Futures-Asset-Transfer
//
// Parameters:
//   - fromEmail: Sender email
//   - toEmail: Recipient email
//   - futuresType: 1: USDT-margined Futures, 2: Coin-margined Futures
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountFuturesAssetTransfer(fromEmail, toEmail string, futuresType int, asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"fromEmail":   fromEmail,
		"toEmail":     toEmail,
		"futuresType": futuresType,
		"asset":       asset,
		"amount":      amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["fromEmail"] = fromEmail
	params["toEmail"] = toEmail
	params["futuresType"] = futuresType
	params["asset"] = asset
	params["amount"] = amount

	return s.SignRequest("POST", "/sapi/v1/sub-account/futures/internalTransfer", params)
}

// SubAccountSpotSummary queries sub-account spot assets summary (For Master Account)
//
// GET /sapi/v1/sub-account/spotSummary
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Spot-Assets-Summary
//
// Optional parameters:
//   - email: Sub account email
//   - page: Default: 1
//   - size: Default 10, max 20
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountSpotSummary(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.SignRequest("GET", "/sapi/v1/sub-account/spotSummary", params)
}

// SubAccountUniversalTransfer performs universal transfer (For Master Account)
//
// POST /sapi/v1/sub-account/universalTransfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Universal-Transfer
//
// You need to enable "internal transfer" option for the api key which requests this endpoint.
// Transfer from master account by default if fromEmail is not sent.
// Transfer to master account by default if toEmail is not sent.
//
// Parameters:
//   - fromAccountType: "SPOT", "USDT_FUTURE", "COIN_FUTURE", "MARGIN"(Cross), "ISOLATED_MARGIN"
//   - toAccountType: "SPOT", "USDT_FUTURE", "COIN_FUTURE", "MARGIN"(Cross), "ISOLATED_MARGIN"
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - fromEmail: Sender email
//   - toEmail: Recipient email
//   - clientTranId: Must be unique
//   - symbol: Only supported under ISOLATED_MARGIN type
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountUniversalTransfer(fromAccountType, toAccountType, asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"fromAccountType": fromAccountType,
		"toAccountType":   toAccountType,
		"asset":           asset,
		"amount":          amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["fromAccountType"] = fromAccountType
	params["toAccountType"] = toAccountType
	params["asset"] = asset
	params["amount"] = amount

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/universalTransfer", params)
}

// SubAccountUniversalTransferHistory queries universal transfer history (For Master Account)
//
// GET /sapi/v1/sub-account/universalTransfer
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Universal-Transfer-History
//
// fromEmail and toEmail cannot be sent at the same time.
// Return fromEmail equal master account email by default.
// Only get the latest history of past 30 days.
//
// Optional parameters:
//   - fromEmail: Sender email
//   - toEmail: Recipient email
//   - clientTranId: Transaction ID
//   - startTime: Start time
//   - endTime: End time
//   - page: Page number
//   - limit: Default 10, max 20
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountUniversalTransferHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/universalTransfer", params)
}

// SubAccountFuturesAccount gets detail on sub-account's futures account V2 (For Master Account)
//
// GET /sapi/v2/sub-account/futures/account
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2
//
// Parameters:
//   - email: Sub-account email
//   - futuresType: 1: USDT-margined Futures, 2: Coin-margined Futures
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountFuturesAccount(email string, futuresType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":       email,
		"futuresType": futuresType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["futuresType"] = futuresType

	return s.LimitedEncodedSignRequest("GET", "/sapi/v2/sub-account/futures/account", params)
}

// SubAccountFuturesAccountSummary gets summary of sub-account's futures account V2 (For Master Account)
//
// GET /sapi/v2/sub-account/futures/accountSummary
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Futures-Account-V2
//
// Parameters:
//   - futuresType: 1: USDT-margined Futures, 2: Coin-margined Futures
//
// Optional parameters:
//   - page: Default:1
//   - limit: Default 10, max 20
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountFuturesAccountSummary(futuresType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(futuresType, "futuresType"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["futuresType"] = futuresType

	return s.SignRequest("GET", "/sapi/v2/sub-account/futures/accountSummary", params)
}

// SubAccountFuturesPositionRisk gets futures position-risk of sub-account V2 (For Master Account)
//
// GET /sapi/v2/sub-account/futures/positionRisk
//
// https://developers.binance.com/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2
//
// Parameters:
//   - email: Sub-account email
//   - futuresType: 1: USDT-margined Futures, 2: Coin-margined Futures
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountFuturesPositionRisk(email string, futuresType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":       email,
		"futuresType": futuresType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["futuresType"] = futuresType

	return s.LimitedEncodedSignRequest("GET", "/sapi/v2/sub-account/futures/positionRisk", params)
}

// SubAccountSpotTransferHistory queries sub-account spot asset transfer history (For Master Account)
//
// GET /sapi/v1/sub-account/sub/transfer/history
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Spot-Asset-Transfer-History
//
// Optional parameters:
//   - fromEmail: Sender email
//   - toEmail: Recipient email
//   - startTime: Default return the history within 100 days
//   - endTime: Default return the history within 100 days
//   - page: Default value: 1
//   - limit: Default value: 500
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountSpotTransferHistory(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/sub/transfer/history", params)
}

// SubAccountEnableLeverageToken enables leverage token for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/blvt/enable
//
// https://developers.binance.com/docs/sub_account/account-management/Enable-Leverage-Token-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//   - enableBlvt: Only true for now
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountEnableLeverageToken(email string, enableBlvt bool, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":      email,
		"enableBlvt": enableBlvt,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["enableBlvt"] = enableBlvt

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/sub-account/blvt/enable", params)
}

// ManagedSubAccountDeposit deposits assets into the managed sub-account (For Investor Master Account)
//
// POST /sapi/v1/managed-subaccount/deposit
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Deposit-Assets-Into-The-Managed-Sub-account
//
// Parameters:
//   - toEmail: Sub-account email
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) ManagedSubAccountDeposit(toEmail, asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"toEmail": toEmail,
		"asset":   asset,
		"amount":  amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["toEmail"] = toEmail
	params["asset"] = asset
	params["amount"] = amount

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/managed-subaccount/deposit", params)
}

// ManagedSubAccountAssets queries managed sub-account asset details (For Investor Master Account)
//
// GET /sapi/v1/managed-subaccount/asset
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Asset-Details
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) ManagedSubAccountAssets(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/managed-subaccount/asset", params)
}

// ManagedSubAccountWithdraw withdraws assets from the managed sub-account (For Investor Master Account)
//
// POST /sapi/v1/managed-subaccount/withdraw
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Withdrawl-Assets-From-The-Managed-Sub-account
//
// Parameters:
//   - fromEmail: Sub-account email
//   - asset: Asset symbol
//   - amount: Amount
//
// Optional parameters:
//   - transferDate: Withdrawals automatically occur on the transfer date (UTC0)
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) ManagedSubAccountWithdraw(fromEmail, asset string, amount float64, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"fromEmail": fromEmail,
		"asset":     asset,
		"amount":    amount,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["fromEmail"] = fromEmail
	params["asset"] = asset
	params["amount"] = amount

	return s.LimitedEncodedSignRequest("POST", "/sapi/v1/managed-subaccount/withdraw", params)
}

// SubAccountUpdateIPRestriction updates IP restriction for sub-account API key (For Master Account)
//
// POST /sapi/v2/sub-account/subAccountApi/ipRestriction
//
// https://developers.binance.com/docs/sub_account/api-management/Add-IP-Restriction-for-Sub-Account-API-key
//
// Parameters:
//   - email: Sub-account email
//   - subAccountApiKey: Sub-account API key
//   - status: IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
//
// Optional parameters:
//   - ipAddress: Can be added in batches, separated by commas
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountUpdateIPRestriction(email, subAccountApiKey, status string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":             email,
		"subAccountApiKey":  subAccountApiKey,
		"status":            status,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["subAccountApiKey"] = subAccountApiKey
	params["status"] = status

	return s.LimitedEncodedSignRequest("POST", "/sapi/v2/sub-account/subAccountApi/ipRestriction", params)
}

// SubAccountAPIGetIPRestriction gets IP restriction for a sub-account API key (For Master Account)
//
// GET /sapi/v1/sub-account/subAccountApi/ipRestriction
//
// https://developers.binance.com/docs/sub_account/api-management/Get-IP-Restriction-for-a-Sub-account-API-Key
//
// Parameters:
//   - email: Sub-account email
//   - subAccountApiKey: Sub-account API key
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountAPIGetIPRestriction(email, subAccountApiKey string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":            email,
		"subAccountApiKey": subAccountApiKey,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["subAccountApiKey"] = subAccountApiKey

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/sub-account/subAccountApi/ipRestriction", params)
}

// SubAccountAPIDeleteIP deletes IP list for a sub-account API key (For Master Account)
//
// DELETE /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList
//
// https://developers.binance.com/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key
//
// Parameters:
//   - email: Sub-account email
//   - subAccountApiKey: Sub-account API key
//   - ipAddress: Can be added in batches, separated by commas
//
// Optional parameters:
//   - thirdPartyName: Third party name
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SubAccountAPIDeleteIP(email, subAccountApiKey, ipAddress string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":            email,
		"subAccountApiKey": subAccountApiKey,
		"ipAddress":        ipAddress,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["subAccountApiKey"] = subAccountApiKey
	params["ipAddress"] = ipAddress

	return s.LimitedEncodedSignRequest("DELETE", "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList", params)
}

// ManagedSubAccountGetSnapshot queries managed sub-account snapshot (For Investor Master Account)
//
// GET /sapi/v1/managed-subaccount/accountSnapshot
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Snapshot
//
// Parameters:
//   - email: Sub-account email
//   - snapshotType: "SPOT", "MARGIN" (cross), "FUTURES" (UM)
//
// Optional parameters:
//   - startTime: Start time
//   - endTime: End time
//   - limit: min 7, max 30, default 7
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) ManagedSubAccountGetSnapshot(email, snapshotType string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email": email,
		"type":  snapshotType,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["type"] = snapshotType

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/managed-subaccount/accountSnapshot", params)
}

// ManagedSubAccountInvestorTransLog queries managed sub-account transfer log (Investor)
//
// GET /sapi/v1/managed-subaccount/queryTransLogForInvestor
//
// https://developers.binance.com/en
//
// Parameters:
//   - email: Sub-account email
//   - startTime: Start time
//   - endTime: End time
//   - page: Page number
//   - limit: Limit
//
// Optional parameters:
//   - transfers: Transfer Direction (FROM/TO)
//   - transferFunctionAccountType: Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
func (s *SubAccountClient) ManagedSubAccountInvestorTransLog(email string, startTime, endTime int64, page, limit int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":     email,
		"startTime": startTime,
		"endTime":   endTime,
		"page":      page,
		"limit":     limit,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = page
	params["limit"] = limit

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/managed-subaccount/queryTransLogForInvestor", params)
}

// ManagedSubAccountTradingTransLog queries managed sub-account transfer log (Trading Team)
//
// GET /sapi/v1/managed-subaccount/queryTransLogForTradeParent
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-Account-Transfer-Log-Trading-Team-Master
//
// Parameters:
//   - email: Sub-account email
//   - startTime: Start time
//   - endTime: End time
//   - page: Page number
//   - limit: Limit
//
// Optional parameters:
//   - transfers: Transfer Direction (FROM/TO)
//   - transferFunctionAccountType: Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
func (s *SubAccountClient) ManagedSubAccountTradingTransLog(email string, startTime, endTime int64, page, limit int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email":     email,
		"startTime": startTime,
		"endTime":   endTime,
		"page":      page,
		"limit":     limit,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = page
	params["limit"] = limit

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/managed-subaccount/queryTransLogForTradeParent", params)
}

// ManagedSubAccountDepositAddress gets managed sub-account deposit address (For Investor Master Account)
//
// GET /sapi/v1/managed-subaccount/deposit/address
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Get-Managed-Sub-account-Deposit-Address
//
// Parameters:
//   - email: Sub-account email
//   - coin: Coin symbol
//
// Optional parameters:
//   - network: Network
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) ManagedSubAccountDepositAddress(email, coin string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"email": email,
		"coin":  coin,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email
	params["coin"] = coin

	return s.LimitedEncodedSignRequest("GET", "/sapi/v1/managed-subaccount/deposit/address", params)
}

// QuerySubAccountAssets queries sub-account assets V4 (For Master Account)
//
// GET /sapi/v4/sub-account/assets
//
// https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Assets-V4
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QuerySubAccountAssets(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v4/sub-account/assets", params)
}

// EnableOptionsForSubAccount enables options for sub-account (For Master Account)
//
// POST /sapi/v1/sub-account/eoptions/enable
//
// https://developers.binance.com/docs/sub_account/account-management/Enable-Options-for-Sub-account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) EnableOptionsForSubAccount(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("POST", "/sapi/v1/sub-account/eoptions/enable", params)
}

// QuerySubAccountTransactionStatistics queries sub-account transaction statistics (For Master Account)
//
// GET /sapi/v1/sub-account/transaction-statistics
//
// https://developers.binance.com/docs/sub_account/account-management/Query-Sub-account-Transaction-Statistics
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QuerySubAccountTransactionStatistics(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v1/sub-account/transaction-statistics", params)
}

// QueryManagedSubAccountTransferLog queries managed sub-account transfer log (For Trading Team Sub Account)
//
// GET /sapi/v1/managed-subaccount/query-trans-log
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-Account-Transfer-Log-Trading-Team-Sub
//
// Parameters:
//   - startTime: UTC timestamp in ms
//   - endTime: UTC timestamp in ms
//   - page: Default 1
//   - limit: Default 500; max 1000
//
// Optional parameters:
//   - transfers: Transfer direction
//   - transferFunctionAccountType: Transfer function account type
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QueryManagedSubAccountTransferLog(startTime, endTime int64, page, limit int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameters(map[string]interface{}{
		"startTime": startTime,
		"endTime":   endTime,
		"page":      page,
		"limit":     limit,
	}); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = page
	params["limit"] = limit

	return s.SignRequest("GET", "/sapi/v1/managed-subaccount/query-trans-log", params)
}

// QueryManagedSubAccountList queries managed sub-account list (For Investor)
//
// GET /sapi/v1/managed-subaccount/info
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-List
//
// Optional parameters:
//   - email: Sub-account email
//   - page: Default 1
//   - limit: Default 500; max 1000
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QueryManagedSubAccountList(params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	return s.SignRequest("GET", "/sapi/v1/managed-subaccount/info", params)
}

// QueryManagedSubAccountMarginAssetDetails queries managed sub-account margin asset details (For Investor Master Account)
//
// GET /sapi/v1/managed-subaccount/marginAsset
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Margin-Asset-Details
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QueryManagedSubAccountMarginAssetDetails(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v1/managed-subaccount/marginAsset", params)
}

// QueryManagedSubAccountFuturesAssetDetails queries managed sub-account futures asset details (For Investor Master Account)
//
// GET /sapi/v1/managed-subaccount/fetch-future-asset
//
// https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Futures-Asset-Details
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) QueryManagedSubAccountFuturesAssetDetails(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v1/managed-subaccount/fetch-future-asset", params)
}

// FuturesPositionRiskOfSubAccount gets futures position-risk of sub-account (For Master Account)
//
// GET /sapi/v1/sub-account/futures/positionRisk
//
// https://developers.binance.com/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) FuturesPositionRiskOfSubAccount(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v1/sub-account/futures/positionRisk", params)
}

// SummaryOfSubAccountSFuturesAccount gets summary of sub-account's futures account V2 (For Master Account)
//
// GET /sapi/v2/sub-account/futures/accountSummary
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Futures-Account-V2
//
// Parameters:
//   - futuresType: 1: USDT Margined Futures, 2: COIN Margined Futures
//
// Optional parameters:
//   - page: default 1
//   - limit: default 10, max 20
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) SummaryOfSubAccountSFuturesAccount(futuresType int, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(futuresType, "futuresType"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["futuresType"] = futuresType

	return s.SignRequest("GET", "/sapi/v2/sub-account/futures/accountSummary", params)
}

// DetailOnSubAccountSFuturesAccount gets detail on sub-account's futures account (For Master Account)
//
// GET /sapi/v1/sub-account/futures/account
//
// https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account
//
// Parameters:
//   - email: Sub-account email
//
// Optional parameters:
//   - recvWindow: The value cannot be greater than 60000
func (s *SubAccountClient) DetailOnSubAccountSFuturesAccount(email string, params map[string]interface{}) ([]byte, error) {
	if err := utils.CheckRequiredParameter(email, "email"); err != nil {
		return nil, err
	}

	if params == nil {
		params = make(map[string]interface{})
	}
	params["email"] = email

	return s.SignRequest("GET", "/sapi/v1/sub-account/futures/account", params)
}
