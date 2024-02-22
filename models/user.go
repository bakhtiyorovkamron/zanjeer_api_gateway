package models

type Sms struct {
	Phone string `json:"phone"`
}
type VerifyNumber struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
