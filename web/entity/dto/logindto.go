package dto

type LoginDto struct {
	Id    int    `json:"id"`
	User  string `json:"user"`
	Token string `json:"token"`
}
