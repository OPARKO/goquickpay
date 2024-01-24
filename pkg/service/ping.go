package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/httpmethod"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
	"github.com/parkeringskompagniet/goquickpay/pkg/service/constants"
)

// NOTE: add query params?

func (c QuickpayClient) Ping(method httpmethod.HTTPMethod) (*quickpay.Pong, error) {
	response, err := c.CallEndpoint(method, "/ping", nil)
	if err != nil {
		return nil, err
	}

	if method == httpmethod.Get && response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, response.StatusCode)
	} else if method == httpmethod.Post && response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, response.StatusCode)
	}

	var pong quickpay.Pong
	err = json.NewDecoder(response.Body).Decode(&pong)
	if err != nil {
		return nil, err
	}

	return &pong, nil
}
