package subscription

type RecurringForm struct {
	TextOnStatement string  `json:"text_on_statement,omitempty"`
	Description     string  `json:"description,omitempty"`
	OrderID         string  `json:"order_id" binding:"required"`
	FeeVAT          float64 `json:"fee_vat,omitempty"`
	Amount          int     `json:"amount" binding:"required"`
	AutoCapture     bool    `json:"auto_capture,omitempty"`
	Autofee         bool    `json:"autofee,omitempty"`
	ZeroAuth        bool    `json:"zero_auth,omitempty"`
}

// schema
// type RecurringForm struct {
// 	Amount          int      `schema:"amount,required"`
// 	OrderID         string   `schema:"order_id,required"`
// 	AutoCapture     *bool    `schema:"auto_capture,omitempty"`
// 	Autofee         *bool    `schema:"autofee,omitempty"`
// 	ZeroAuth        *bool    `schema:"zero_auth,omitempty"`
// 	TextOnStatement *string  `schema:"text_on_statement,omitempty"`
// 	FeeVAT          *float64 `schema:"fee_vat,omitempty"`
// 	Description     *string  `schema:"description,omitempty"`
// }
