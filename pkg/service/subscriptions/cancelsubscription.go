package subscriptions

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
	"github.com/parkeringskompagniet/goquickpay/pkg/service/constants"
)

func (s SubscriptionService) CancelSubscription(subscriptionID int64, callback *string) (*quickpay.Subscription, error) {
	request, err := s.Client.PrepareWithPath(quickpay.Post, fmt.Sprintf(constants.SUBSCRIPTIONS_CANCEL, subscriptionID), nil)
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

	if statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return DecodeSubscriptioFrom(response.Body)
}
