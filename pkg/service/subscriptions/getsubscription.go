package subscriptions

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

func (s SubscriptionService) GetSubscription(subscriptionID int64) (*quickpay.Subscription, error) {
	res, err := s.Client.CallWithPath(quickpay.Get, fmt.Sprintf(constants.SUBSCRIPTIONS_ID, subscriptionID), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(res.Status)

	statusCode := res.StatusCode
	if statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(res.Status)
	} else if statusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return DecodeSubscriptioFrom(res.Body)
}
