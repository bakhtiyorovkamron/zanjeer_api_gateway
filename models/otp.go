package models

type SmsOtp struct {
	SmsId string `json:"sms_id"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
type ConfirmOTP struct {
	SmsId string  `json:"sms_id"`
	Phone *string `json:"phone"`
	Code  string  `json:"code"`
}
