package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"goquickpay/pkg/call"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
)

type QuickPayClient struct {
	ApiKey  string
	BaseURL string
}

func (q QuickPayClient) setupRequest(method call.HTTPMethod, endpoint string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(string(method), q.BaseURL+endpoint, body)
	if err != nil {
		return nil, errors.New("there was an error setting up base request")
	}

	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(":" + q.ApiKey))

	request.Header.Add("Accept-Version", "v10")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+encodedAPIKey)

	return request, nil
}

func (q QuickPayClient) CallEndpoint(method call.HTTPMethod, endpoint string, data any) (*http.Response, error) {
	request, err := q.PrepareEndPoint(method, endpoint, data)
	if err != nil {
		return nil, err
	}

	return q.CallEndPointWith(request)
}

func (q QuickPayClient) PrepareEndPoint(method call.HTTPMethod, endpoint string, data any) (*http.Request, error) {
	if data == nil {
		return q.setupRequest(method, endpoint, strings.NewReader(""))
	}

	body, err := q.EncodeBody(data)
	if err != nil {
		return nil, err
	}

	return q.setupRequest(method, endpoint, body)
}

func (q QuickPayClient) CallEndPointWith(request *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(request)
}

func (q QuickPayClient) EncodeQuery(data any) (string, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}
	err := encoder.Encode(data, values)
	if err != nil {
		return "", err
	}

	return values.Encode(), nil
}

// TODO: custom parser for custom schema
func (q QuickPayClient) EncodeBody(data any) (io.Reader, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(bytes)), nil
}
