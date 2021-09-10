package moke

import (
	"log"

	"github.com/dgkg/keywords/auth/handler/model"
	"github.com/dgkg/keywords/db"
)

type DB struct {
	users map[string]*model.User
}

func New() *DB {
	db := &DB{
		users: make(map[string]*model.User),
	}
	initDB(db)

	return db
}

func initDB(db *DB) {
	var u model.User
	for i := 0; i < 100; i++ {
		u2 := u.Random()
		db.users[u2.Name] = u2
	}

	dataUser := map[string]*model.User{
		"casper": {
			ID:          1,
			Name:        "Casper",
			Password:    "Tatata",
			AccessLevel: 1,
		},
		"boss": {
			ID:          2,
			Name:        "The Boss",
			Password:    "bibibi",
			AccessLevel: 10,
		},
	}
	for k := range dataUser {
		db.users[dataUser[k].Name] = dataUser[k]
	}
}

func (d *DB) GetUserByLogin(login string) (*model.User, error) {
	log.Println("GetUserByLogin", login)
	u, ok := d.users[login]
	if !ok {
		log.Println("GetUserByLogin not found")
		return nil, db.NewErrNotFound(nil, login)
	}
	log.Println("GetUserByLogin found", u)
	return u, nil
}

func (d *DB) CreateUser(u *model.User) error {
	d.users[u.Name] = u
	return nil
}
