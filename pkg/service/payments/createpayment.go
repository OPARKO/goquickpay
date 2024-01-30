package payments

import (
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type PaymentsForm struct{}

type PaymentRequestForm struct {
	InvoiceAddress  quickpay.Address       `schema:"invoice_address,omitempty"`
	ShippingAddress quickpay.Address       `schema:"shipping_address,omitempty"`
	Shipping        quickpay.Shipping      `schema:"shipping,omitempty"`
	Shopsystem      quickpay.ShopSystem    `schema:"shopsystem,omitempty"`
	Variables       map[string]interface{} `schema:"variables,omitempty"`
	Currency        string                 `schema:"currency"`
	OrderID         string                 `schema:"order_id"`
	TextOnStatement string                 `schema:"text_on_statement,omitempty"`
	Basket          []quickpay.Basket      `schema:"basket"`
	BrandingID      int64                  `schema:"branding_id,omitempty"`
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
