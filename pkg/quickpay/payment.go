package quickpay

import "time"

type Payment struct {
	ID              int64        `json:"id"`
	ULID            string       `json:"ulid"`
	MerchantID      int64        `json:"merchant_id"`
	OrderID         string       `json:"order_id"`
	Accepted        bool         `json:"accepted"`
	Type            string       `json:"type"`
	TextOnStatement *string      `json:"text_on_statement"`
	BrandingID      *int64       `json:"branding_id"`
	Variables       interface{}  `json:"variables"`
	Currency        string       `json:"currency"`
	State           string       `json:"state"`
	Metadata        Metadata     `json:"metadata"`
	Link            *PaymentLink `json:"link"`
	ShippingAddress *Address     `json:"shipping_address"`
	InvoiceAddress  *Address     `json:"invoice_address"`
	Basket          []*Basket    `json:"basket"`
	Shipping        *Shipping    `json:"shipping"`
	Operations      []*Operation `json:"operations"`
	TestMode        bool         `json:"test_mode"`
	Acquirer        *string      `json:"acquirer"`
	Facilitator     *string      `json:"facilitator"`
	CreatedAt       *time.Time   `json:"created_at"`
	UpdatedAt       *time.Time   `json:"updated_at"`
	RetentedAt      *time.Time   `json:"retented_at"`
	Balance         int64        `json:"balance"`
	Fee             *int64       `json:"fee"`
	SubscriptionID  int64        `json:"subscription_id"`
	DeadlineAt      *time.Time   `json:"deadline_at"`
	ResellerID      *int64       `json:"reseller_id"`
}
