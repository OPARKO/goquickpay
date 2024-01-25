package quickpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
)

func (c QuickpayClient) Ping(method HTTPMethod, params url.Values) (*Pong, error) {
	path, err := c.CreateBaseUrl(constants.QUICKPAY_PING, params)
	if err != nil {
		return nil, err
	}

	response, err := c.CallWithURL(method, path, nil)
	if err != nil {
		return nil, err
	}

	if method == Get && response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, response.StatusCode)
	} else if method == Post && response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, response.StatusCode)
	}

	var pong Pong
	err = json.NewDecoder(response.Body).Decode(&pong)
	if err != nil {
		return nil, err
	}

	return &pong, nil
}
