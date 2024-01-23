package quickpay

type Metadata struct {
	Type                   string   `json:"type"`
	Origin                 string   `json:"origin"`
	Brand                  string   `json:"brand"`
	BIN                    string   `json:"bin"`
	Corporate              bool     `json:"corporate"`
	Last4                  string   `json:"last4"`
	ExpMonth               int      `json:"exp_month"`
	ExpYear                int      `json:"exp_year"`
	Country                string   `json:"country"`
	Is3DSecure             bool     `json:"is_3d_secure"`
	Secure3dType           string   `json:"3d_secure_type"`
	IssuedTo               string   `json:"issued_to"`
	Hash                   string   `json:"hash"`
	Moto                   bool     `json:"moto"`
	Number                 any      `json:"number"`
	CustomerIP             string   `json:"customer_ip"`
	CustomerCountry        string   `json:"customer_country"`
	FraudSuspected         bool     `json:"fraud_suspected"`
	FraudRemarks           []string `json:"fraud_remarks"`
	FraudReported          bool     `json:"fraud_reported"`
	FraudReportDescription string   `json:"fraud_report_description"`
	FraudReportedAt        string   `json:"fraud_reported_at"`
	NINNumber              string   `json:"nin_number"`
	NINCountryCode         string   `json:"nin_country_code"`
	NINGender              string   `json:"nin_gender"`
	ShopsystemName         string   `json:"shopsystem_name"`
	ShopsystemVersion      string   `json:"shopsystem_version"`
	ClearHouseMerchantId   string   `json:"ch_mid"`
}
