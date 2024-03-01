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
type CreateDeviceRequest struct {
	Name      string `json:"name"`
	Imei      string `json:"imei"`
	DriverId  string `json:"driver_id"`
	Type      string `json:"device_type_id"`
	IpAddress string `json:"ip_address"`
}
type CreateDeviceResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Imei      string `json:"imei"`
	DriverId  string `json:"driver_id"`
	Type      string `json:"device_type_id"`
	IpAddress string `json:"ip_address"`
	CreatedAt string `json:"created_at"`
}
