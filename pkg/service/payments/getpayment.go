package payments

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

func (s PaymentService) GetPayment(paymentID int64, operationSize uint) (*quickpay.Payment, error) {
	values := url.Values{}
	if operationSize > 0 {
		values.Add("operation_size", strconv.FormatUint(uint64(operationSize), 10))
	}

	u, err := s.Client.CreateBaseUrl(fmt.Sprintf(constants.PAYMENTS_ID, paymentID), values)
	if err != nil {
		return nil, err
	}

	response, err := s.Client.CallWithURL(quickpay.Get, u, nil)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode
	if statusCode == http.StatusForbidden || statusCode == http.StatusNotFound {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	return DecodePaymentFrom(response.Body)
}
