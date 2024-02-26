package models

type Driver struct {
	Id        string `json:"id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}