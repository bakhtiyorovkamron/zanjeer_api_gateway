package models

type Admin struct {
	Id        string `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
type GetAdmins struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
