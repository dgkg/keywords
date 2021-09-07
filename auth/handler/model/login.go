package model

type PayloadLogin struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}
