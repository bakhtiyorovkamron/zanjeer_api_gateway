package models

type Admin struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}
type GetAdmins struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
}
