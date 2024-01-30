package subscriptions

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type CreateForm struct {
	Variables       map[string]string   `schema:"variables,omitempty"`
	BrandingID      int64               `schema:"branding_id,omitempty"`
	InvoiceAddress  quickpay.Address    `schema:"invoice_address,omitempty"`
	ShippingAddress quickpay.Address    `schema:"shipping_address,omitempty"`
	ShopSystem      quickpay.ShopSystem `schema:"shopsystem,omitempty"`
	OrderID         string              `schema:"order_id" binding:"required,len=4:20"`
	Currency        string              `schema:"currency" binding:"required"`
	Description     string              `schema:"description" binding:"required"`
	TextOnStatement string              `schema:"text_on_statement,omitempty"`
	ThreeDSV2       quickpay.ThreedsV2  `schema:"threeds_v2,omitempty"`
	GroupIDs        []int               `schema:"group_ids,omitempty"`
	Unscheduled     bool                `schema:"unscheduled,omitempty"`
}

func (s SubscriptionService) CreateSubscription(form CreateForm) (*quickpay.Subscription, error) {
	response, err := s.Client.CallWithPath(quickpay.Post, constants.SUBSCRIPTIONS, form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusAccepted {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return DecodeSubscriptioFrom(response.Body)
}
