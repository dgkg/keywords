package model

type User struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Password    string `json:"pass"`
	AccessLevel int    `json:"access_level"`
}
