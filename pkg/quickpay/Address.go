package quickpay

type Address struct {
	Name           string `json:"name,omitempty"`
	Att            string `json:"att,omitempty"`
	CompanyName    string `json:"company_name,omitempty"`
	Street         string `json:"street,omitempty"`
	HouseNumber    string `json:"house_number,omitempty"`
	HouseExtension string `json:"house_extension,omitempty"`
	City           string `json:"city,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	Region         string `json:"region,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
	VATNo          string `json:"vat_no,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	MobileNumber   string `json:"mobile_number,omitempty"`
	Email          string `json:"email,omitempty"`
}
