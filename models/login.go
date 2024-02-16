package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type StandardResponse struct {
	StatusId string `json:"status_id"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
}
