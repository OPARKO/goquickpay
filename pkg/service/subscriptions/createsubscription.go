package subscriptions

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type CreateForm struct {
	Variables       map[string]string   `json:"variables,omitempty"`
	BrandingID      int64               `json:"branding_id,omitempty"`
	InvoiceAddress  quickpay.Address    `json:"invoice_address,omitempty"`
	ShippingAddress quickpay.Address    `json:"shipping_address,omitempty"`
	ShopSystem      quickpay.ShopSystem `json:"shopsystem,omitempty"`
	OrderID         string              `json:"order_id" binding:"required,len=4:20"`
	Currency        string              `json:"currency" binding:"required"`
	Description     string              `json:"description" binding:"required"`
	TextOnStatement string              `json:"text_on_statement,omitempty"`
	ThreeDSV2       quickpay.ThreedsV2  `json:"threeds_v2,omitempty"`
	GroupIDs        []int               `json:"group_ids,omitempty"`
	Unscheduled     bool                `json:"unscheduled,omitempty"`
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
