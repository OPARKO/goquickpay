package subscriptions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parkeringskompagniet/goquickpay/pkg/constants"
	"github.com/parkeringskompagniet/goquickpay/pkg/quickpay"
)

type CreateOrUpdateLinkForm struct {
	Amount                    int64                  `schema:"amount"`
	AgreementID               int64                  `schema:"agreement_id,omitempty"`
	Language                  string                 `schema:"language,omitempty"`
	ContinueURL               string                 `schema:"continue_url,omitempty"`
	CancelURL                 string                 `schema:"cancel_url,omitempty"`
	CallbackURL               string                 `schema:"callback_url,omitempty"`
	PaymentMethods            string                 `schema:"payment_methods,omitempty"`
	AutoFee                   bool                   `schema:"auto_fee,omitempty"`
	BrandingID                int64                  `schema:"branding_id,omitempty"`
	GoogleAnalyticsTrackingID string                 `schema:"google_analytics_tracking_id,omitempty"`
	GoogleAnalyticsClientID   string                 `schema:"google_analytics_client_id,omitempty"`
	Acquirer                  string                 `schema:"acquirer,omitempty"`
	Deadline                  int64                  `schema:"deadline,omitempty"`
	Framed                    bool                   `schema:"framed,omitempty"`
	BrandingConfig            map[string]interface{} `schema:"branding_config,omitempty"`
	FeeVAT                    float64                `schema:"fee_vat,omitempty"`
	Moto                      bool                   `schema:"moto,omitempty"`
	CustomerEmail             string                 `schema:"customer_email,omitempty"`
	InvoiceAddressSelection   bool                   `schema:"invoice_address_selection,omitempty"`
	ShippingAddressSelection  bool                   `schema:"shipping_address_selection,omitempty"`
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
