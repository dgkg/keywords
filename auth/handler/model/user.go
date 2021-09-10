package model

import (
	"math/rand"

	xRand "golang.org/x/exp/rand"

	srcRand "keywords/src/rand"
)

type User struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"pass"`
	AccessLevel int    `json:"access_level"`
}

func (u *User) Random() *User {
	return &User{
		ID:          xRand.Uint64(),
		Name:        srcRand.String(10),
		Password:    srcRand.String(10),
		AccessLevel: rand.Intn(3),
	}
}
