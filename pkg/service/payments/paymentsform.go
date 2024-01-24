package payments

import "goquickpay/pkg/quickpay"

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
