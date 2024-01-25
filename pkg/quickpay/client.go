package quickpay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"

	"github.com/gorilla/schema"
)

type QuickpayClient struct {
	BaseUrl string
	ApiKey  string
}

func NewClient(apiKey string) QuickpayClient {
	return QuickpayClient{constants.DEFAULT_QUICKPAY_URL, apiKey}
}

func (c QuickpayClient) setupRequest(method HTTPMethod, path *url.URL, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(string(method), path.String(), body)
	if err != nil {
		return nil, errors.New("there was an error setting up base request")
	}

	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(":" + c.ApiKey))

	request.Header.Add("Accept-Version", "v10")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+encodedAPIKey)

	return request, nil
}

func (c QuickpayClient) CreateBaseUrl(path string, params url.Values) (*url.URL, error) {
	u, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, err
	}

	u.Path = path
	u.RawQuery = params.Encode()

	return u, nil
}

func (c QuickpayClient) PrepareWithURL(method HTTPMethod, u *url.URL, data any) (*http.Request, error) {
	if data == nil {
		return c.setupRequest(method, u, strings.NewReader(""))
	}

	body, err := c.EncodeBody(data)
	if err != nil {
		return nil, err
	}

	return c.setupRequest(method, u, body)
}

func (c QuickpayClient) PrepareWithPath(method HTTPMethod, path string, data any) (*http.Request, error) {
	u, err := c.CreateBaseUrl(path, nil)
	if err != nil {
		return nil, err
	}

	return c.PrepareWithURL(method, u, data)
}

func (c QuickpayClient) CallWithURL(method HTTPMethod, u *url.URL, body any) (*http.Response, error) {
	request, err := c.PrepareWithURL(method, u, body)
	if err != nil {
		return nil, err
	}

	return c.CallWithRequest(request)
}

func (c QuickpayClient) CallWithPath(method HTTPMethod, path string, body any) (*http.Response, error) {
	u, err := c.CreateBaseUrl(path, nil)
	if err != nil {
		return nil, err
	}

	return c.CallWithURL(method, u, body)
}

func (c QuickpayClient) CallWithRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}

	fmt.Println(request)

	return client.Do(request)
}

func (c QuickpayClient) ConverToURLValues(data any) (url.Values, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	err := encoder.Encode(data, values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (c QuickpayClient) EncodeBody(data any) (io.Reader, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(bytes)), nil
}
