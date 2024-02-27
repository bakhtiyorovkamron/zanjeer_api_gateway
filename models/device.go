package models

type DeviceType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type GetDeviceTypeListRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
}
