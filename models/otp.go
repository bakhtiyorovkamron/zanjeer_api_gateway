package models

type SmsOtp struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
