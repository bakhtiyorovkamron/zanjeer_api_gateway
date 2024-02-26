package models

type Sms struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
	SmsId string `json:"sms_id"`
}
type VerifyNumber struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}