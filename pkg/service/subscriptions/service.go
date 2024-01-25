package subscriptions

import (
	"encoding/json"
	"io"

	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type SubscriptionService struct {
	Client quickpay.QuickpayClient
}

func NewSubscriptionService(client quickpay.QuickpayClient) SubscriptionService {
	return SubscriptionService{client}
}

//TODO:
//
// GET /subscriptions GetSubscriptions
// DELETE /subscriptions/{id}/link DeletePaymentLink
// GET /subscriptions/{id} GetSubscription
// PATCH /subscriptions/{id} UpdateSubscription
// POST /subscriptions/{id}/session CreateSubscriptionSession
// POST /subscriptions/{id}/authorize AuthorizeAsubscription
// POST /subscriptions/{id}/fraud-reportCreate FraudConfirmationReport
// GET /subscriptions/{id}/operations/{operation_id} GetOperation
// PATCH /subscriptions/{id}/operations/{operation_id} UpdateOperation
// GET /subscriptions/{id}/payments GetAllSubscriptionPayments

func DecodeSubscriptioFrom(body io.ReadCloser) (*quickpay.Subscription, error) {
	var subscription quickpay.Subscription

	err := json.NewDecoder(body).Decode(&subscription)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}
