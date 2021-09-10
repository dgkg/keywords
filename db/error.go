package db

import "fmt"

type ErrDB struct {
	params []interface{}
	kind   string
	err    error
}

const (
	// ErrNotFound when the given ressource is not found.
	ErrNotFound = "not found"
)

func (e ErrDB) Error() string {
	return fmt.Sprintf("db: params:%v kind:%v err:%v", e.params, e.kind, e.err)
}

func NewErrNotFound(err error, params ...interface{}) error {
	return &ErrDB{
		params: params,
		kind:   ErrNotFound,
		err:    err,
	}
}
