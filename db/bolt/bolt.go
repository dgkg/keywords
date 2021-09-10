package bolt

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"

	"keywords/auth/handler/model"
	"keywords/db"
)

var bucketUser []byte = []byte("user")

func initDB(path string) *bolt.DB {
	db := open(path)

	tx, err := db.Begin(true)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.CreateBucket(bucketUser)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Commit the transaction and check for error.
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	var u model.User
	var s *Store = &Store{conn: db}
	for i := 0; i < 100; i++ {
		s.CreateUser(u.Random())
	}

	var dataUser = map[string]model.User{
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

	for _, u := range dataUser {
		s.CreateUser(&u)
	}

	return db
}

func New(path string, init ...int) *Store {
	var db *bolt.DB
	if len(init) > 0 {
		log.Println("init db")
		db = initDB(path)
	} else {
		log.Println("open db")
		db = open(path)
	}

	log.Println("return store")
	return &Store{
		conn: db,
	}
}

type Store struct {
	conn *bolt.DB
}

func (s *Store) GetUserByLogin(login string) (*model.User, error) {

	log.Println("GetUserByLogin")
	if len(login) == 0 {
		return nil, db.NewErrNotFound(nil, login)
	}

	var data []byte
	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketUser)
		log.Println("try to get user", login)
		data = b.Get([]byte(login))
		if len(data) == 0 {
			return db.NewErrNotFound(nil, login)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var u model.User
	return &u, json.Unmarshal(data, &u)
}

// CreateUser saves u to the store. The new user ID is set on u once the data is persisted.
func (s *Store) CreateUser(u *model.User) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket(bucketUser)

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		// bind id.
		u.ID = id

		// Marshal user data into bytes.
		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(u.Name), buf)
	})
}

func open(path string) *bolt.DB {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
