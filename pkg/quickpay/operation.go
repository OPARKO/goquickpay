package quickpay

import "time"

type Operation struct {
	ID                   int         `json:"id"`
	Type                 string      `json:"type"`
	Amount               int         `json:"amount"`
	Pending              bool        `json:"pending"`
	QPStatusCode         *string     `json:"qp_status_code"`
	QPStatusMsg          *string     `json:"qp_status_msg"`
	AQStatusCode         *string     `json:"aq_status_code"`
	AQStatusMsg          *string     `json:"aq_status_msg"`
	Data                 interface{} `json:"data"`
	CallbackURL          string      `json:"callback_url"`
	CallbackSuccess      *bool       `json:"callback_success"`
	CallbackResponseCode *string     `json:"callback_response_code"`
	CallbackDuration     *int        `json:"callback_duration"`
	Acquirer             string      `json:"acquirer"`
	SecureStatus3D       *string     `json:"3d_secure_status"`
	CallbackAt           *time.Time  `json:"callback_at"`
	CreatedAt            time.Time   `json:"created_at"`
}
