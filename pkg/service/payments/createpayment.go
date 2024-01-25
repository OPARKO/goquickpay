package payments

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
	"github.com/parkeringskompagniet/goquickpay/pkg/service/constants"
)

type PaymentsForm struct{}

type PaymentRequestForm struct {
	InvoiceAddress  quickpay.Address    `json:"invoice_address,omitempty"`
	ShippingAddress quickpay.Address    `json:"shipping_address,omitempty"`
	Shipping        quickpay.Shipping   `json:"shipping,omitempty"`
	Shopsystem      quickpay.ShopSystem `json:"shopsystem,omitempty"`
	Variables       map[string]any      `json:"variables,omitempty"`
	Currency        string              `json:"currency"`
	OrderID         string              `json:"order_id"`
	TextOnStatement string              `json:"text_on_statement,omitempty"`
	Basket          []quickpay.Basket   `json:"basket"`
	BrandingID      int                 `json:"branding_id,omitempty"`
}

func (s PaymentService) CreatePayment(form PaymentsForm) (*quickpay.Payment, error) {
	response, err := s.Client.CallWithPath(quickpay.Post, constants.PAYMENTS, form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusCreated {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return DecodePaymentFrom(response.Body)
}
