package subscriptions

type RecurringForm struct {
	TextOnStatement string  `json:"text_on_statement,omitempty"`
	Description     string  `json:"description,omitempty"`
	OrderID         string  `json:"order_id" binding:"required"`
	FeeVAT          float64 `json:"fee_vat,omitempty"`
	Amount          int64   `json:"amount" binding:"required"`
	AutoCapture     bool    `json:"auto_capture,omitempty"`
	Autofee         bool    `json:"autofee,omitempty"`
	ZeroAuth        bool    `json:"zero_auth,omitempty"`
}
