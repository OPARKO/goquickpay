package subscriptions

import (
	"encoding/json"
	"fmt"
	"goquickpay/pkg/httpmethod"
	"goquickpay/pkg/quickpay"
	"goquickpay/pkg/service"
	"goquickpay/pkg/service/constants"
	"io"
	"net/http"
)

type Service struct {
	Client service.QuickpayClient
}

func NewService(client service.QuickpayClient) Service {
	return Service{client}
}

// GET /subscriptions GetSubscriptions

func (s Service) CreateSubscription(form CreateForm) (*quickpay.Subscription, error) {
	response, err := s.Client.CallEndpoint(httpmethod.Post, constants.SUBSCRIPTIONS, form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf("quickpay unknown response code: %d", statusCode)
	}

	return DecodeSubscriptioFrom(response.Body)
}

func (s Service) CreateOrUpdatePaymentLink(id int64, form CreateOrUpdateLinkForm) (*quickpay.PaymentLinkUrl, error) {
	response, err := s.Client.CallEndpoint(httpmethod.Put, fmt.Sprintf(constants.SUBSCRIPTIONS_LINK, id), form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusOK {
		return nil, fmt.Errorf("quickpay unknown response code: %d", statusCode)
	}

	var linkUrl quickpay.PaymentLinkUrl

	err = json.NewDecoder(response.Body).Decode(&linkUrl)
	if err != nil {
		return nil, err
	}

	return &linkUrl, nil
}

// DELETE /subscriptions/{id}/link DeletePaymentLink
// GET /subscriptions/{id} GetSubscription
// PATCH /subscriptions/{id} UpdateSubscription
// POST /subscriptions/{id}/session CreateSubscriptionSession
// POST /subscriptions/{id}/authorize AuthorizeAsubscription

func (s Service) CancelSubscription(id int64, callback *string) (*quickpay.Subscription, error) {
	request, err := s.Client.PrepareEndPoint(httpmethod.Post, fmt.Sprintf(constants.SUBSCRIPTIONS_CANCEL, id), nil)
	if err != nil {
		return nil, err
	}

	if callback != nil {
		request.Header.Add(constants.SUBSCRIPTIONS_HEADER_CALLBACK_URL, *callback)
	}

	response, err := s.Client.CallEndPointWith(request)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf("quickpay unknown response code: %d", statusCode)
	}

	return DecodeSubscriptioFrom(response.Body)
}

func (s Service) CreateSubscriptionRecurringPayment(id int64, form RecurringForm, callback *string) (*quickpay.Payment, error) {
	request, err := s.Client.PrepareEndPoint(httpmethod.Get, fmt.Sprintf(constants.SUBSCRIPTIONS_RECURRING, id), form)
	if err != nil {
		return nil, err
	}

	if callback != nil {
		request.Header.Add(constants.SUBSCRIPTIONS_HEADER_CALLBACK_URL, *callback)
	}

	response, err := s.Client.CallEndPointWith(request)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf("quickpay unknown response code: %d", statusCode)
	}

	var payment quickpay.Payment

	err = json.NewDecoder(response.Body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

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
