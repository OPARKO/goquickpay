package quickpay

type PaymentLinkUrl struct {
	Url string `json:"url" binding:"required"`
}
