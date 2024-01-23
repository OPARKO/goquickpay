package quickpay

type Shipping struct {
	Method         string  `json:"method,omitempty"`
	Company        string  `json:"company,omitempty"`
	TrackingNumber string  `json:"tracking_number,omitempty"`
	TrackingURL    string  `json:"tracking_url,omitempty"`
	Amount         int     `json:"amount,omitempty"`
	VatRate        float64 `json:"vat_rate,omitempty"`
}
