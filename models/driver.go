package models

type Driver struct {
	Id        string `json:"id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	Verified  bool   `json:"verified"`
}
type DriverList struct {
	Count   int32    `json:"count"`
	Drivers []Driver `json:"drivers"`
}
type DriverSearchRequest struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
