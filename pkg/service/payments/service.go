package payments

import (
	"encoding/json"
	"fmt"
	"goquickpay/pkg/httpmethod"
	"goquickpay/pkg/quickpay"
	"goquickpay/pkg/service"
	"goquickpay/pkg/service/constants"
	"net/http"
)

type Service struct {
	Client service.QuickpayClient
}

func NewService(client service.QuickpayClient) Service {
	return Service{client}
}

func (s Service) CreatePayment(form PaymentsForm) (*quickpay.Payment, error) {
	response, err := s.Client.CallEndpoint(httpmethod.Post, constants.PAYMENTS, form)
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
