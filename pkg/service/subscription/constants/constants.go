package constants

// TODO: move endpoints
const (
	SUBSCRIPTIONS              = "/subscriptions"
	SUBSCRIPTIONS_LINK         = "/subscriptions/%d/link"
	SUBSCRIPTIONS_ID           = "/subscriptions/%d"
	SUBSCRIPTIONS_SESSION      = "/subscriptions/%d/session"
	SUBSCRIPTIONS_AUTHORIZE    = "/subscriptions/%d/authorize"
	SUBSCRIPTIONS_CANCEL       = "/subscriptions/%d/cancel"
	SUBSCRIPTIONS_RECURRING    = "/subscriptions/%d/recurring"
	SUBSCRIPTIONS_FRAUD_REPORT = "/subscriptions/%d/fraud-report"
	SUBSCRIPTIONS_OPERATION    = "/subscriptions/%d/operation/%d"
	SUBSCRIPTIONS_PAYMENT      = "/subscriptions/%d/payments"
)

const (
	SUBSCRIPTIONS_HEADER_CALLBACK_URL = "QuickPay-Callback-Url"
)
