package payments

import (
	"encoding/json"
	"fmt"
	"goquickpay/pkg/call"
	"goquickpay/pkg/quickpay"
	"goquickpay/pkg/service"
	"goquickpay/pkg/service/payments/constants"
	"net/http"
)

type Service struct {
	Client service.QuickPayClient
}

func (s Service) CreatePayment(form PaymentsForm) (*quickpay.Payment, error) {
	response, err := s.Client.CallEndpoint(call.Post, constants.PAYMENTS, form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("quickpay unknown response code: %d", statusCode)
	}

	var payment quickpay.Payment

	err = json.NewDecoder(response.Body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
