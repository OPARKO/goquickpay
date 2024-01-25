package payments

import (
	"encoding/json"
	"io"

	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type PaymentService struct {
	Client quickpay.QuickpayClient
}

func NewPaymentService(client quickpay.QuickpayClient) PaymentService {
	return PaymentService{client}
}

// TODO:
//
// GET /payments GetAllPayments
// PUT /payments/{id}/linkCreate or update a payment link
// DELETE /payments/{id}/linkDelete payment link
// GET /payments/{id}Get payment
// PATCH /payments/{id}Update payment
// POST /payments/{id}/sessionCreate payment session
// POST /payments/{id}/authorizeAuthorize payment
// POST /payments/{id}/captureCapture payment
// POST /payments/{id}/refundRefund payment
// POST /payments/{id}/cancelCancel payment
// POST /payments/{id}/renewRenew authorization
// POST /payments/{id}/fraud-reportCreate fraud confirmation report
// GET /payments/{id}/operations/{operation_id}Get Operation
// PATCH /payments/{id}/operations/{operation_id} UpdateOperation

func DecodePaymentFrom(body io.ReadCloser) (*quickpay.Payment, error) {
	var payment quickpay.Payment

	err := json.NewDecoder(body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
