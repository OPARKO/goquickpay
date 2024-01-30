package subscriptions

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
	"github.com/parkeringskompagniet/goquickpay/pkg/service/payments"
)

type RecurringForm struct {
	TextOnStatement string  `schema:"text_on_statement,omitempty"`
	Description     string  `schema:"description,omitempty"`
	OrderID         string  `schema:"order_id" binding:"required"`
	FeeVAT          float64 `schema:"fee_vat,omitempty"`
	Amount          int64   `schema:"amount" binding:"required"`
	AutoCapture     bool    `schema:"auto_capture,omitempty"`
	Autofee         bool    `schema:"autofee,omitempty"`
	ZeroAuth        bool    `schema:"zero_auth,omitempty"`
}

func (s SubscriptionService) CreateSubscriptionRecurringPayment(subscriptionID int64, form RecurringForm, callback *string) (*quickpay.Payment, error) {
	request, err := s.Client.PrepareWithPath(quickpay.Post, fmt.Sprintf(constants.SUBSCRIPTIONS_RECURRING, subscriptionID), form)
	if err != nil {
		return nil, err
	}

	if callback != nil {
		request.Header.Add(constants.SUBSCRIPTIONS_HEADER_CALLBACK_URL, *callback)
	}

	response, err := s.Client.CallWithRequest(request)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return payments.DecodePaymentFrom(response.Body)
}
