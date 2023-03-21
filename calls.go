package deriv


func (a *DerivAPI) ActiveSymbols(r ActiveSymbols) (resp ActiveSymbolsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ApiToken(r ApiToken) (resp ApiTokenResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppDelete(r AppDelete) (resp AppDeleteResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppGet(r AppGet) (resp AppGetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppList(r AppList) (resp AppListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppMarkupDetails(r AppMarkupDetails) (resp AppMarkupDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppMarkupStatistics(r AppMarkupStatistics) (resp AppMarkupStatisticsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppRegister(r AppRegister) (resp AppRegisterResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AppUpdate(r AppUpdate) (resp AppUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) AssetIndex(r AssetIndex) (resp AssetIndexResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Authorize(r Authorize) (resp AuthorizeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Balance(r Balance) (resp BalanceResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Buy(r Buy) (resp BuyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) BuyContractForMultipleAccounts(r BuyContractForMultipleAccounts) (resp BuyContractForMultipleAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Cancel(r Cancel) (resp CancelResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Cashier(r Cashier) (resp CashierResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ContractUpdate(r ContractUpdate) (resp ContractUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ContractUpdateHistory(r ContractUpdateHistory) (resp ContractUpdateHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ContractsFor(r ContractsFor) (resp ContractsForResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) CopyStart(r CopyStart) (resp CopyStartResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) CopyStop(r CopyStop) (resp CopyStopResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) CopytradingList(r CopytradingList) (resp CopytradingListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) CopytradingStatistics(r CopytradingStatistics) (resp CopytradingStatisticsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) CryptoConfig(r CryptoConfig) (resp CryptoConfigResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) DocumentUpload(r DocumentUpload) (resp DocumentUploadResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) EconomicCalendar(r EconomicCalendar) (resp EconomicCalendarResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ExchangeRates(r ExchangeRates) (resp ExchangeRatesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Forget(r Forget) (resp ForgetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ForgetAll(r ForgetAll) (resp ForgetAllResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) GetAccountStatus(r GetAccountStatus) (resp GetAccountStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) GetFinancialAssessment(r GetFinancialAssessment) (resp GetFinancialAssessmentResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) GetLimits(r GetLimits) (resp GetLimitsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) GetSelfExclusion(r GetSelfExclusion) (resp GetSelfExclusionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) GetSettings(r GetSettings) (resp GetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) IdentityVerificationDocumentAdd(r IdentityVerificationDocumentAdd) (resp IdentityVerificationDocumentAddResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) LandingCompany(r LandingCompany) (resp LandingCompanyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) LandingCompanyDetails(r LandingCompanyDetails) (resp LandingCompanyDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) LoginHistory(r LoginHistory) (resp LoginHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Logout(r Logout) (resp LogoutResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5Deposit(r Mt5Deposit) (resp Mt5DepositResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5GetSettings(r Mt5GetSettings) (resp Mt5GetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5LoginList(r Mt5LoginList) (resp Mt5LoginListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5NewAccount(r Mt5NewAccount) (resp Mt5NewAccountResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5PasswordChange(r Mt5PasswordChange) (resp Mt5PasswordChangeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5PasswordCheck(r Mt5PasswordCheck) (resp Mt5PasswordCheckResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5PasswordReset(r Mt5PasswordReset) (resp Mt5PasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Mt5Withdrawal(r Mt5Withdrawal) (resp Mt5WithdrawalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) NewAccountMaltainvest(r NewAccountMaltainvest) (resp NewAccountMaltainvestResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) NewAccountReal(r NewAccountReal) (resp NewAccountRealResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) NewAccountVirtual(r NewAccountVirtual) (resp NewAccountVirtualResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) OauthApps(r OauthApps) (resp OauthAppsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertCreate(r P2PAdvertCreate) (resp P2PAdvertCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertInfo(r P2PAdvertInfo) (resp P2PAdvertInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertList(r P2PAdvertList) (resp P2PAdvertListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertUpdate(r P2PAdvertUpdate) (resp P2PAdvertUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserAdverts(r P2PAdvertiserAdverts) (resp P2PAdvertiserAdvertsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserCreate(r P2PAdvertiserCreate) (resp P2PAdvertiserCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserInfo(r P2PAdvertiserInfo) (resp P2PAdvertiserInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserList(r P2PAdvertiserList) (resp P2PAdvertiserListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserPaymentMethods(r P2PAdvertiserPaymentMethods) (resp P2PAdvertiserPaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserRelations(r P2PAdvertiserRelations) (resp P2PAdvertiserRelationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PAdvertiserUpdate(r P2PAdvertiserUpdate) (resp P2PAdvertiserUpdateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PChatCreate(r P2PChatCreate) (resp P2PChatCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderCancel(r P2POrderCancel) (resp P2POrderCancelResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderConfirm(r P2POrderConfirm) (resp P2POrderConfirmResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderCreate(r P2POrderCreate) (resp P2POrderCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderDispute(r P2POrderDispute) (resp P2POrderDisputeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderInfo(r P2POrderInfo) (resp P2POrderInfoResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderList(r P2POrderList) (resp P2POrderListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2POrderReview(r P2POrderReview) (resp P2POrderReviewResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PPaymentMethods(r P2PPaymentMethods) (resp P2PPaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) P2PPing(r P2PPing) (resp P2PPingResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentMethods(r PaymentMethods) (resp PaymentMethodsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentCreate(r PaymentagentCreate) (resp PaymentagentCreateResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentDetails(r PaymentagentDetails) (resp PaymentagentDetailsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentList(r PaymentagentList) (resp PaymentagentListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentTransfer(r PaymentagentTransfer) (resp PaymentagentTransferResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentWithdraw(r PaymentagentWithdraw) (resp PaymentagentWithdrawResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PaymentagentWithdrawJustification(r PaymentagentWithdrawJustification) (resp PaymentagentWithdrawJustificationResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) PayoutCurrencies(r PayoutCurrencies) (resp PayoutCurrenciesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Ping(r Ping) (resp PingResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Portfolio(r Portfolio) (resp PortfolioResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ProfitTable(r ProfitTable) (resp ProfitTableResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Proposal(r Proposal) (resp ProposalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ProposalOpenContract(r ProposalOpenContract) (resp ProposalOpenContractResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) RealityCheck(r RealityCheck) (resp RealityCheckResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) ResidenceList(r ResidenceList) (resp ResidenceListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) RevokeOauthApp(r RevokeOauthApp) (resp RevokeOauthAppResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Sell(r Sell) (resp SellResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SellContractForMultipleAccounts(r SellContractForMultipleAccounts) (resp SellContractForMultipleAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SellExpired(r SellExpired) (resp SellExpiredResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SetAccountCurrency(r SetAccountCurrency) (resp SetAccountCurrencyResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SetFinancialAssessment(r SetFinancialAssessment) (resp SetFinancialAssessmentResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SetSelfExclusion(r SetSelfExclusion) (resp SetSelfExclusionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) SetSettings(r SetSettings) (resp SetSettingsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Statement(r Statement) (resp StatementResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) StatesList(r StatesList) (resp StatesListResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Ticks(r Ticks) (resp TicksResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TicksHistory(r TicksHistory) (resp TicksHistoryResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Time(r Time) (resp TimeResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TncApproval(r TncApproval) (resp TncApprovalResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TopupVirtual(r TopupVirtual) (resp TopupVirtualResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TradingDurations(r TradingDurations) (resp TradingDurationsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TradingPlatformInvestorPasswordReset(r TradingPlatformInvestorPasswordReset) (resp TradingPlatformInvestorPasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TradingPlatformPasswordReset(r TradingPlatformPasswordReset) (resp TradingPlatformPasswordResetResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TradingServers(r TradingServers) (resp TradingServersResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TradingTimes(r TradingTimes) (resp TradingTimesResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) Transaction(r Transaction) (resp TransactionResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) TransferBetweenAccounts(r TransferBetweenAccounts) (resp TransferBetweenAccountsResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) UnsubscribeEmail(r UnsubscribeEmail) (resp UnsubscribeEmailResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) VerifyEmail(r VerifyEmail) (resp VerifyEmailResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) VerifyEmailCellxpert(r VerifyEmailCellxpert) (resp VerifyEmailCellxpertResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

func (a *DerivAPI) WebsiteStatus(r WebsiteStatus) (resp WebsiteStatusResp, err error) {
	id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return
}

