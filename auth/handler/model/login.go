package model

type PayloadLogin struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}
