package db

import "github.com/dgkg/keywords/auth/handler/model"

type Storer interface {
	GetUserByLogin(login string) (*model.User, error)
	CreateUser(u *model.User) error
}
