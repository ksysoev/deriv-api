// Code generated by bin/generate_call.go, DO NOT EDIT.

package deriv

import "github.com/ksysoev/deriv-api/schema"

// AccountList Returns all accounts belonging to the authorized user.
func (a *DerivAPI) AccountList(r schema.AccountList) (resp schema.AccountListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ActiveSymbols Retrieve a list of all currently active symbols (underlying markets upon which contracts are available for trading).
func (a *DerivAPI) ActiveSymbols(r schema.ActiveSymbols) (resp schema.ActiveSymbolsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ApiToken This call manages API tokens
func (a *DerivAPI) ApiToken(r schema.ApiToken) (resp schema.ApiTokenResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppDelete The request for deleting an application.
func (a *DerivAPI) AppDelete(r schema.AppDelete) (resp schema.AppDeleteResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppGet To get the information of the OAuth application specified by 'app_id'
func (a *DerivAPI) AppGet(r schema.AppGet) (resp schema.AppGetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppList List all of the account's OAuth applications
func (a *DerivAPI) AppList(r schema.AppList) (resp schema.AppListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppMarkupDetails Retrieve details of `app_markup` according to criteria specified.
func (a *DerivAPI) AppMarkupDetails(r schema.AppMarkupDetails) (resp schema.AppMarkupDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppMarkupStatistics Retrieve statistics of `app_markup`.
func (a *DerivAPI) AppMarkupStatistics(r schema.AppMarkupStatistics) (resp schema.AppMarkupStatisticsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppRegister Register a new OAuth application
func (a *DerivAPI) AppRegister(r schema.AppRegister) (resp schema.AppRegisterResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AppUpdate Update a new OAuth application
func (a *DerivAPI) AppUpdate(r schema.AppUpdate) (resp schema.AppUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// AssetIndex Retrieve a list of all available underlyings and the corresponding contract types and duration boundaries. If the user is logged in, only the assets available for that user's landing company will be returned.
func (a *DerivAPI) AssetIndex(r schema.AssetIndex) (resp schema.AssetIndexResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Authorize Authorize current WebSocket session to act on behalf of the owner of a given token. Must precede requests that need to access client account, for example purchasing and selling contracts or viewing portfolio.
func (a *DerivAPI) Authorize(r schema.Authorize) (resp schema.AuthorizeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Balance Get user account balance
func (a *DerivAPI) Balance(r schema.Balance) (resp schema.BalanceResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Buy Buy a Contract
func (a *DerivAPI) Buy(r schema.Buy) (resp schema.BuyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// BuyContractForMultipleAccounts Buy a Contract for multiple Accounts specified by the `tokens` parameter. Note, although this is an authorized call, the contract is not bought for the authorized account.
func (a *DerivAPI) BuyContractForMultipleAccounts(r schema.BuyContractForMultipleAccounts) (resp schema.BuyContractForMultipleAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Cancel Cancel contract with contract id
func (a *DerivAPI) Cancel(r schema.Cancel) (resp schema.CancelResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Cashier Request the cashier info for the specified type.
func (a *DerivAPI) Cashier(r schema.Cashier) (resp schema.CashierResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ConfirmEmail Verifies the email for the user using verification code passed in the request object
func (a *DerivAPI) ConfirmEmail(r schema.ConfirmEmail) (resp schema.ConfirmEmailResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ContractUpdate Update a contract condition.
func (a *DerivAPI) ContractUpdate(r schema.ContractUpdate) (resp schema.ContractUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ContractUpdateHistory Request for contract update history.
func (a *DerivAPI) ContractUpdateHistory(r schema.ContractUpdateHistory) (resp schema.ContractUpdateHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ContractsFor For a given symbol, get the list of currently available contracts, and the latest barrier and duration limits for each contract.
func (a *DerivAPI) ContractsFor(r schema.ContractsFor) (resp schema.ContractsForResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ContractsForCompany Get the list of currently available contracts for a given landing company.
func (a *DerivAPI) ContractsForCompany(r schema.ContractsForCompany) (resp schema.ContractsForCompanyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CopyStart Start copy trader bets
func (a *DerivAPI) CopyStart(r schema.CopyStart) (resp schema.CopyStartResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CopyStop Stop copy trader bets
func (a *DerivAPI) CopyStop(r schema.CopyStop) (resp schema.CopyStopResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CopytradingList Retrieves a list of active copiers and/or traders for Copy Trading
func (a *DerivAPI) CopytradingList(r schema.CopytradingList) (resp schema.CopytradingListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CopytradingStatistics Retrieve performance, trading, risk and copiers statistics of trader.
func (a *DerivAPI) CopytradingStatistics(r schema.CopytradingStatistics) (resp schema.CopytradingStatisticsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CryptoConfig The request for cryptocurrencies configuration.
func (a *DerivAPI) CryptoConfig(r schema.CryptoConfig) (resp schema.CryptoConfigResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// CryptoEstimations Get the current estimations for cryptocurrencies. E.g. Withdrawal fee.
func (a *DerivAPI) CryptoEstimations(r schema.CryptoEstimations) (resp schema.CryptoEstimationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// DocumentUpload Request KYC information from client
func (a *DerivAPI) DocumentUpload(r schema.DocumentUpload) (resp schema.DocumentUploadResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// EconomicCalendar Specify a currency to receive a list of events related to that specific currency. For example, specifying USD will return a list of USD-related events. If the currency is omitted, you will receive a list for all currencies.
func (a *DerivAPI) EconomicCalendar(r schema.EconomicCalendar) (resp schema.EconomicCalendarResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ExchangeRates Retrieves the exchange rate from a base currency to a target currency supported by the system.
func (a *DerivAPI) ExchangeRates(r schema.ExchangeRates) (resp schema.ExchangeRatesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Forget Immediately cancel the real-time stream of messages with a specific ID.
func (a *DerivAPI) Forget(r schema.Forget) (resp schema.ForgetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ForgetAll Immediately cancel the real-time streams of messages of given type.
func (a *DerivAPI) ForgetAll(r schema.ForgetAll) (resp schema.ForgetAllResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// GetAccountStatus Get Account Status
func (a *DerivAPI) GetAccountStatus(r schema.GetAccountStatus) (resp schema.GetAccountStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// GetFinancialAssessment This call gets the financial assessment details. The 'financial assessment' is a questionnaire that clients of certain Landing Companies need to complete, due to regulatory and KYC (know your client) requirements.
func (a *DerivAPI) GetFinancialAssessment(r schema.GetFinancialAssessment) (resp schema.GetFinancialAssessmentResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// GetLimits Trading and Withdrawal Limits for a given user
func (a *DerivAPI) GetLimits(r schema.GetLimits) (resp schema.GetLimitsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// GetSelfExclusion Allows users to exclude themselves from the website for certain periods of time, or to set limits on their trading activities. This facility is a regulatory requirement for certain Landing Companies.
func (a *DerivAPI) GetSelfExclusion(r schema.GetSelfExclusion) (resp schema.GetSelfExclusionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// GetSettings Get User Settings (email, date of birth, address etc)
func (a *DerivAPI) GetSettings(r schema.GetSettings) (resp schema.GetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// IdentityVerificationDocumentAdd Adds document information such as issuing country, id and type for identity verification processes.
func (a *DerivAPI) IdentityVerificationDocumentAdd(r schema.IdentityVerificationDocumentAdd) (resp schema.IdentityVerificationDocumentAddResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// KycAuthStatus Get KYC Authentication Status
func (a *DerivAPI) KycAuthStatus(r schema.KycAuthStatus) (resp schema.KycAuthStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// LandingCompany The company has a number of licensed subsidiaries in various jurisdictions, which are called Landing Companies. This call will return the appropriate Landing Company for clients of a given country. The landing company may differ for derived contracts (Synthetic Indices) and Financial contracts (Forex, Stock Indices, Commodities).
func (a *DerivAPI) LandingCompany(r schema.LandingCompany) (resp schema.LandingCompanyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// LandingCompanyDetails The company has a number of licensed subsidiaries in various jurisdictions, which are called Landing Companies (and which are wholly owned subsidiaries of the Deriv Group). This call provides information about each Landing Company.
func (a *DerivAPI) LandingCompanyDetails(r schema.LandingCompanyDetails) (resp schema.LandingCompanyDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// LoginHistory Retrieve a summary of login history for user.
func (a *DerivAPI) LoginHistory(r schema.LoginHistory) (resp schema.LoginHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Logout Logout the session
func (a *DerivAPI) Logout(r schema.Logout) (resp schema.LogoutResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5Deposit This call allows deposit into MT5 account from Binary account.
func (a *DerivAPI) Mt5Deposit(r schema.Mt5Deposit) (resp schema.Mt5DepositResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5GetSettings Get MT5 user account settings
func (a *DerivAPI) Mt5GetSettings(r schema.Mt5GetSettings) (resp schema.Mt5GetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5LoginList Get list of MT5 accounts for client
func (a *DerivAPI) Mt5LoginList(r schema.Mt5LoginList) (resp schema.Mt5LoginListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5NewAccount This call creates new MT5 user, either demo or real money user.
func (a *DerivAPI) Mt5NewAccount(r schema.Mt5NewAccount) (resp schema.Mt5NewAccountResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5PasswordChange To change passwords of the MT5 account.
func (a *DerivAPI) Mt5PasswordChange(r schema.Mt5PasswordChange) (resp schema.Mt5PasswordChangeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5PasswordCheck This call validates the main password for the MT5 user
func (a *DerivAPI) Mt5PasswordCheck(r schema.Mt5PasswordCheck) (resp schema.Mt5PasswordCheckResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5PasswordReset To reset the password of MT5 account.
func (a *DerivAPI) Mt5PasswordReset(r schema.Mt5PasswordReset) (resp schema.Mt5PasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Mt5Withdrawal This call allows withdrawal from MT5 account to Binary account.
func (a *DerivAPI) Mt5Withdrawal(r schema.Mt5Withdrawal) (resp schema.Mt5WithdrawalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// NewAccountMaltainvest This call opens a new real-money account with the `maltainvest` Landing Company. This call can be made from a virtual-money account or real-money account at Deriv (Europe) Limited. If it is the latter, client information fields in this call will be ignored and data from your existing real-money account will be used.
func (a *DerivAPI) NewAccountMaltainvest(r schema.NewAccountMaltainvest) (resp schema.NewAccountMaltainvestResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// NewAccountReal This call opens a new real-money account. This call can be made from a virtual-money or a real-money account. If it is the latter, client information fields in this call will be ignored and data from your existing real-money account will be used.
func (a *DerivAPI) NewAccountReal(r schema.NewAccountReal) (resp schema.NewAccountRealResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// NewAccountVirtual Create a new virtual-money account.
func (a *DerivAPI) NewAccountVirtual(r schema.NewAccountVirtual) (resp schema.NewAccountVirtualResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// OauthApps List all my used OAuth applications.
func (a *DerivAPI) OauthApps(r schema.OauthApps) (resp schema.OauthAppsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertCreate Creates a P2P (Peer to Peer) advert. Can only be used by an approved P2P advertiser.
func (a *DerivAPI) P2PAdvertCreate(r schema.P2PAdvertCreate) (resp schema.P2PAdvertCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertInfo Retrieve information about a P2P advert.
func (a *DerivAPI) P2PAdvertInfo(r schema.P2PAdvertInfo) (resp schema.P2PAdvertInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertList Returns available adverts for use with `p2p_order_create` .
func (a *DerivAPI) P2PAdvertList(r schema.P2PAdvertList) (resp schema.P2PAdvertListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertUpdate Updates a P2P advert. Can only be used by the advertiser.
func (a *DerivAPI) P2PAdvertUpdate(r schema.P2PAdvertUpdate) (resp schema.P2PAdvertUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserAdverts Returns all P2P adverts created by the authorized client. Can only be used by a registered P2P advertiser.
func (a *DerivAPI) P2PAdvertiserAdverts(r schema.P2PAdvertiserAdverts) (resp schema.P2PAdvertiserAdvertsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserCreate Registers the client as a P2P advertiser.
func (a *DerivAPI) P2PAdvertiserCreate(r schema.P2PAdvertiserCreate) (resp schema.P2PAdvertiserCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserInfo Retrieve information about a P2P advertiser.
func (a *DerivAPI) P2PAdvertiserInfo(r schema.P2PAdvertiserInfo) (resp schema.P2PAdvertiserInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserList Retrieve advertisers has/had trade with the current advertiser.
func (a *DerivAPI) P2PAdvertiserList(r schema.P2PAdvertiserList) (resp schema.P2PAdvertiserListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserPaymentMethods Manage or list P2P advertiser payment methods.
func (a *DerivAPI) P2PAdvertiserPaymentMethods(r schema.P2PAdvertiserPaymentMethods) (resp schema.P2PAdvertiserPaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserRelations Updates and returns favourite and blocked advertisers of the current user.
func (a *DerivAPI) P2PAdvertiserRelations(r schema.P2PAdvertiserRelations) (resp schema.P2PAdvertiserRelationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PAdvertiserUpdate Update the information of the P2P advertiser for the current account. Can only be used by an approved P2P advertiser.
func (a *DerivAPI) P2PAdvertiserUpdate(r schema.P2PAdvertiserUpdate) (resp schema.P2PAdvertiserUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PChatCreate Creates a P2P chat for the specified order.
func (a *DerivAPI) P2PChatCreate(r schema.P2PChatCreate) (resp schema.P2PChatCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PCountryList List all or specific country and its payment methods.
func (a *DerivAPI) P2PCountryList(r schema.P2PCountryList) (resp schema.P2PCountryListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderCancel Cancel a P2P order.
func (a *DerivAPI) P2POrderCancel(r schema.P2POrderCancel) (resp schema.P2POrderCancelResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderConfirm Confirm a P2P order.
func (a *DerivAPI) P2POrderConfirm(r schema.P2POrderConfirm) (resp schema.P2POrderConfirmResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderCreate Creates a P2P order for the specified advert.
func (a *DerivAPI) P2POrderCreate(r schema.P2POrderCreate) (resp schema.P2POrderCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderDispute Dispute a P2P order.
func (a *DerivAPI) P2POrderDispute(r schema.P2POrderDispute) (resp schema.P2POrderDisputeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderInfo Retrieves the information about a P2P order.
func (a *DerivAPI) P2POrderInfo(r schema.P2POrderInfo) (resp schema.P2POrderInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderList List active orders.
func (a *DerivAPI) P2POrderList(r schema.P2POrderList) (resp schema.P2POrderListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2POrderReview Creates a review for the specified order.
func (a *DerivAPI) P2POrderReview(r schema.P2POrderReview) (resp schema.P2POrderReviewResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PPaymentMethods List all P2P payment methods.
func (a *DerivAPI) P2PPaymentMethods(r schema.P2PPaymentMethods) (resp schema.P2PPaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PPing Keeps the connection alive and updates the P2P advertiser's online status. The advertiser will be considered offline 60 seconds after a call is made.
func (a *DerivAPI) P2PPing(r schema.P2PPing) (resp schema.P2PPingResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// P2PSettings Request P2P Settings information.
func (a *DerivAPI) P2PSettings(r schema.P2PSettings) (resp schema.P2PSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentMethods Will return a list payment methods available for the given country. If the request is authenticated the client's residence country will be used.
func (a *DerivAPI) PaymentMethods(r schema.PaymentMethods) (resp schema.PaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentCreate Saves client's payment agent details.
func (a *DerivAPI) PaymentagentCreate(r schema.PaymentagentCreate) (resp schema.PaymentagentCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentDetails Gets client's payment agent details.
func (a *DerivAPI) PaymentagentDetails(r schema.PaymentagentDetails) (resp schema.PaymentagentDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentList Will return a list of Payment Agents for a given country for a given currency. Payment agents allow users to deposit and withdraw funds using local payment methods that might not be available via the main website's cashier system.
func (a *DerivAPI) PaymentagentList(r schema.PaymentagentList) (resp schema.PaymentagentListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentTransfer Payment Agent Transfer - this call is available only to accounts that are approved Payment Agents.
func (a *DerivAPI) PaymentagentTransfer(r schema.PaymentagentTransfer) (resp schema.PaymentagentTransferResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentWithdraw Initiate a withdrawal to an approved Payment Agent.
func (a *DerivAPI) PaymentagentWithdraw(r schema.PaymentagentWithdraw) (resp schema.PaymentagentWithdrawResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PaymentagentWithdrawJustification Provide justification to perform withdrawal using a Payment Agent.
func (a *DerivAPI) PaymentagentWithdrawJustification(r schema.PaymentagentWithdrawJustification) (resp schema.PaymentagentWithdrawJustificationResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// PayoutCurrencies Retrieve a list of available option payout currencies. If a user is logged in, only the currencies available for the account will be returned.
func (a *DerivAPI) PayoutCurrencies(r schema.PayoutCurrencies) (resp schema.PayoutCurrenciesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Ping To send the ping request to the server. Mostly used to test the connection or to keep it alive.
func (a *DerivAPI) Ping(r schema.Ping) (resp schema.PingResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Portfolio Receive information about my current portfolio of outstanding options
func (a *DerivAPI) Portfolio(r schema.Portfolio) (resp schema.PortfolioResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ProfitTable Retrieve a summary of account Profit Table, according to given search criteria
func (a *DerivAPI) ProfitTable(r schema.ProfitTable) (resp schema.ProfitTableResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Proposal Gets latest price for a specific contract.
func (a *DerivAPI) Proposal(r schema.Proposal) (resp schema.ProposalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ProposalOpenContract Get latest price (and other information) for a contract in the user's portfolio
func (a *DerivAPI) ProposalOpenContract(r schema.ProposalOpenContract) (resp schema.ProposalOpenContractResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// RealityCheck Retrieve summary of client's trades and account for the Reality Check facility. A 'reality check' means a display of time elapsed since the session began, and associated client profit/loss. The Reality Check facility is a regulatory requirement for certain landing companies.
func (a *DerivAPI) RealityCheck(r schema.RealityCheck) (resp schema.RealityCheckResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// ResidenceList This call returns a list of countries and 2-letter country codes, suitable for populating the account opening form.
func (a *DerivAPI) ResidenceList(r schema.ResidenceList) (resp schema.ResidenceListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// RevokeOauthApp Used for revoking access of particular app.
func (a *DerivAPI) RevokeOauthApp(r schema.RevokeOauthApp) (resp schema.RevokeOauthAppResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Sell Sell a Contract as identified from a previous `portfolio` call.
func (a *DerivAPI) Sell(r schema.Sell) (resp schema.SellResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SellContractForMultipleAccounts Sell contracts for multiple accounts simultaneously. Uses the shortcode response from `buy_contract_for_multiple_accounts` to identify the contract, and authorisation tokens to select which accounts to sell those contracts on. Note that only the accounts identified by the tokens will be affected. This will not sell the contract on the currently-authorised account unless you include the token for the current account.
func (a *DerivAPI) SellContractForMultipleAccounts(r schema.SellContractForMultipleAccounts) (resp schema.SellContractForMultipleAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SellExpired This call will try to sell any expired contracts and return the number of sold contracts.
func (a *DerivAPI) SellExpired(r schema.SellExpired) (resp schema.SellExpiredResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SetAccountCurrency Set account currency, this will be default currency for your account i.e currency for trading, deposit. Please note that account currency can only be set once, and then can never be changed.
func (a *DerivAPI) SetAccountCurrency(r schema.SetAccountCurrency) (resp schema.SetAccountCurrencyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SetFinancialAssessment This call sets the financial assessment details based on the client's answers to analyze whether they possess the experience and knowledge to understand the risks involved with binary options trading.
func (a *DerivAPI) SetFinancialAssessment(r schema.SetFinancialAssessment) (resp schema.SetFinancialAssessmentResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SetSelfExclusion Set Self-Exclusion (this call should be used in conjunction with `get_self_exclusion`)
func (a *DerivAPI) SetSelfExclusion(r schema.SetSelfExclusion) (resp schema.SetSelfExclusionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// SetSettings Set User Settings (this call should be used in conjunction with `get_settings`)
func (a *DerivAPI) SetSettings(r schema.SetSettings) (resp schema.SetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Statement Retrieve a summary of account transactions, according to given search criteria
func (a *DerivAPI) Statement(r schema.Statement) (resp schema.StatementResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// StatesList For a given country, returns a list of States of that country. This is useful to populate the account opening form.
func (a *DerivAPI) StatesList(r schema.StatesList) (resp schema.StatesListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Ticks Initiate a continuous stream of spot price updates for a given symbol.
func (a *DerivAPI) Ticks(r schema.Ticks) (resp schema.TicksResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TicksHistory Get historic tick data for a given symbol.
func (a *DerivAPI) TicksHistory(r schema.TicksHistory) (resp schema.TicksHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Time Request back-end server epoch time.
func (a *DerivAPI) Time(r schema.Time) (resp schema.TimeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TinValidations Get the validations for Tax Identification Numbers (TIN)
func (a *DerivAPI) TinValidations(r schema.TinValidations) (resp schema.TinValidationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TncApproval To approve the latest version of terms and conditions.
func (a *DerivAPI) TncApproval(r schema.TncApproval) (resp schema.TncApprovalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TopupVirtual When a virtual-money's account balance becomes low, it can be topped up using this call.
func (a *DerivAPI) TopupVirtual(r schema.TopupVirtual) (resp schema.TopupVirtualResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingDurations Retrieve a list of all available underlyings and the corresponding contract types and trading duration boundaries. If the user is logged in, only the assets available for that user's landing company will be returned.
func (a *DerivAPI) TradingDurations(r schema.TradingDurations) (resp schema.TradingDurationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingPlatformInvestorPasswordReset Reset the investor password of a Trading Platform Account
func (a *DerivAPI) TradingPlatformInvestorPasswordReset(r schema.TradingPlatformInvestorPasswordReset) (resp schema.TradingPlatformInvestorPasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingPlatformPasswordReset Reset the password of a Trading Platform Account
func (a *DerivAPI) TradingPlatformPasswordReset(r schema.TradingPlatformPasswordReset) (resp schema.TradingPlatformPasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingPlatformStatus Request trading platform status
func (a *DerivAPI) TradingPlatformStatus(r schema.TradingPlatformStatus) (resp schema.TradingPlatformStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingServers Get the list of servers for a trading platform.
func (a *DerivAPI) TradingServers(r schema.TradingServers) (resp schema.TradingServersResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TradingTimes Receive a list of market opening times for a given date.
func (a *DerivAPI) TradingTimes(r schema.TradingTimes) (resp schema.TradingTimesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// Transaction Subscribe to transaction notifications
func (a *DerivAPI) Transaction(r schema.Transaction) (resp schema.TransactionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// TransferBetweenAccounts This call allows transfers between accounts held by a given user. Transfer funds between your fiat and cryptocurrency accounts (for a fee). Please note that account_from should be same as current authorized account.
func (a *DerivAPI) TransferBetweenAccounts(r schema.TransferBetweenAccounts) (resp schema.TransferBetweenAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// UnsubscribeEmail It unsubscribe user from the email subscription.
func (a *DerivAPI) UnsubscribeEmail(r schema.UnsubscribeEmail) (resp schema.UnsubscribeEmailResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// VerifyEmail Verify an email address for various purposes. The system will send an email to the address containing a security code for verification.
func (a *DerivAPI) VerifyEmail(r schema.VerifyEmail) (resp schema.VerifyEmailResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// WebsiteConfig Request server config.
func (a *DerivAPI) WebsiteConfig(r schema.WebsiteConfig) (resp schema.WebsiteConfigResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

// WebsiteStatus Request server status.
func (a *DerivAPI) WebsiteStatus(r schema.WebsiteStatus) (resp schema.WebsiteStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

