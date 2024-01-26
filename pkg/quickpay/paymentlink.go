package quickpay

type PaymentLink struct {
	URL                       string      `json:"url"`
	AgreementID               int64       `json:"agreement_id"`
	Language                  string      `json:"language"`
	Amount                    int64       `json:"amount"`
	ContinueURL               string      `json:"continue_url"`
	CancelURL                 string      `json:"cancel_url"`
	CallbackURL               string      `json:"callback_url"`
	PaymentMethods            string      `json:"payment_methods"`
	AutoFee                   bool        `json:"auto_fee"`
	AutoCapture               bool        `json:"auto_capture"`
	BrandingID                int64       `json:"branding_id"`
	GoogleAnalyticsClientID   string      `json:"google_analytics_client_id"`
	GoogleAnalyticsTrackingID string      `json:"google_analytics_tracking_id"`
	Version                   string      `json:"version"`
	Acquirer                  string      `json:"acquirer"`
	Deadline                  int64       `json:"deadline"`
	Framed                    bool        `json:"framed"`
	BrandingConfig            interface{} `json:"branding_config"`
	InvoiceAddressSelection   bool        `json:"invoice_address_selection"`
	ShippingAddressSelection  bool        `json:"shipping_address_selection"`
	FeeVat                    int64       `json:"fee_vat"`
	Moto                      bool        `json:"moto"`
	CustomerEmail             string      `json:"customer_email"`
}
