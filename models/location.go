package models

type GetDeviceLocationRequest struct {
}
type GetDeviceLocationResponse struct {
	Imei string
	// Location  Location
	Time      string
	Angle     int16
	Speed     int16
	Longitude string
	Latitude  string
}
type Location struct {
	Type        string
	Coordinates []int32
}
