package db

import "keywords/auth/handler/model"

type Storer interface {
	GetUserByLogin(login string) (*model.User, error)
}
