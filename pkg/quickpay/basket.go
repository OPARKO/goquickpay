package quickpay

type Basket struct {
	ItemNo    string  `json:"item_no,omitempty"`
	ItemName  string  `json:"item_name,omitempty"`
	Quantity  int64   `json:"qty,omitempty"`
	ItemPrice float64 `json:"item_price,omitempty"`
	VatRate   float64 `json:"vat_rate,omitempty"`
}
