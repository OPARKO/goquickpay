package subscriptions

type CreateOrUpdateLinkForm struct {
	Amount                    int            `json:"amount"`
	AgreementID               int            `json:"agreement_id,omitempty"`
	Language                  string         `json:"language,omitempty"`
	ContinueURL               string         `json:"continue_url,omitempty"`
	CancelURL                 string         `json:"cancel_url,omitempty"`
	CallbackURL               string         `json:"callback_url,omitempty"`
	PaymentMethods            string         `json:"payment_methods,omitempty"`
	AutoFee                   bool           `json:"auto_fee,omitempty"`
	BrandingID                int            `json:"branding_id,omitempty"`
	GoogleAnalyticsTrackingID string         `json:"google_analytics_tracking_id,omitempty"`
	GoogleAnalyticsClientID   string         `json:"google_analytics_client_id,omitempty"`
	Acquirer                  string         `json:"acquirer,omitempty"`
	Deadline                  int            `json:"deadline,omitempty"`
	Framed                    bool           `json:"framed,omitempty"`
	BrandingConfig            map[string]any `json:"branding_config,omitempty"`
	FeeVAT                    float64        `json:"fee_vat,omitempty"`
	Moto                      bool           `json:"moto,omitempty"`
	CustomerEmail             string         `json:"customer_email,omitempty"`
	InvoiceAddressSelection   bool           `json:"invoice_address_selection,omitempty"`
	ShippingAddressSelection  bool           `json:"shipping_address_selection,omitempty"`
}
