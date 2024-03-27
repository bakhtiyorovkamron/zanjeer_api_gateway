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
type GetAdminsRequest struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
}
type GetAdminsResponse struct {
	Count  int     `json:"count"`
	Admins []Admin `json:"admins"`
}
type EditAdminsResponse struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}
