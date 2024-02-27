package models

type Driver struct {
	Id        string `json:"id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
}
type DriverList struct {
	Count   int32    `json:"count"`
	Drivers []Driver `json:"drivers"`
}
type DriverSearchRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
