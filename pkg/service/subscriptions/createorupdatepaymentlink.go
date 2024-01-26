package subscriptions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type CreateOrUpdateLinkForm struct {
	Amount                    int                    `json:"amount"`
	AgreementID               int                    `json:"agreement_id,omitempty"`
	Language                  string                 `json:"language,omitempty"`
	ContinueURL               string                 `json:"continue_url,omitempty"`
	CancelURL                 string                 `json:"cancel_url,omitempty"`
	CallbackURL               string                 `json:"callback_url,omitempty"`
	PaymentMethods            string                 `json:"payment_methods,omitempty"`
	AutoFee                   bool                   `json:"auto_fee,omitempty"`
	BrandingID                int                    `json:"branding_id,omitempty"`
	GoogleAnalyticsTrackingID string                 `json:"google_analytics_tracking_id,omitempty"`
	GoogleAnalyticsClientID   string                 `json:"google_analytics_client_id,omitempty"`
	Acquirer                  string                 `json:"acquirer,omitempty"`
	Deadline                  int                    `json:"deadline,omitempty"`
	Framed                    bool                   `json:"framed,omitempty"`
	BrandingConfig            map[string]interface{} `json:"branding_config,omitempty"`
	FeeVAT                    float64                `json:"fee_vat,omitempty"`
	Moto                      bool                   `json:"moto,omitempty"`
	CustomerEmail             string                 `json:"customer_email,omitempty"`
	InvoiceAddressSelection   bool                   `json:"invoice_address_selection,omitempty"`
	ShippingAddressSelection  bool                   `json:"shipping_address_selection,omitempty"`
}

func (s SubscriptionService) CreateOrUpdatePaymentLink(transactionID int64, form CreateOrUpdateLinkForm) (*quickpay.PaymentLinkUrl, error) {
	response, err := s.Client.CallWithPath(quickpay.Put, fmt.Sprintf(constants.SUBSCRIPTIONS_LINK, transactionID), form)
	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode == http.StatusBadRequest {
		return nil, fmt.Errorf(response.Status)
	} else if statusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrNotExpectedResponseCode, statusCode)
	}

	var linkUrl quickpay.PaymentLinkUrl

	err = json.NewDecoder(response.Body).Decode(&linkUrl)
	if err != nil {
		return nil, err
	}

	return &linkUrl, nil
}
