package json

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// DateFormat is the standard way to communicate date with JSON.
	DateFormat = "2006-01-02"
	// LocationParis by default the application uses the localisation time
	// of Paris UTC+1.
	LocationParis = "Europe/Paris"
)

var locParis *time.Location

func init() {
	var err error
	locParis, err = time.LoadLocation(LocationParis)
	if err != nil {
		panic(err)
	}
}

// BirthDate is used to communicate with JSON format the time.Time
// with the UTC+1.
type BirthDate time.Time

// String implement the Stringer interface in the pkg fmt.
func (bd BirthDate) String() string {
	return time.Time(bd).Format(DateFormat)
}

// UnmarshalJSON implements the Unmarshaler interface of the encoding/json std package.
func (bd *BirthDate) UnmarshalJSON(b []byte) error {
	if bd == nil {
		return nil
	}

	var aux string
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	if len(aux) == 0 {
		aux = "0000-00-00"
	}

	t, err := time.ParseInLocation(DateFormat, aux, locParis)
	if err != nil {
		return err
	}
	fmt.Println("time got in Unmarshal", t.String())
	*bd = BirthDate(t)

	return nil
}

// MarshalJSON is implementing the encoding/json interface Marshaler.
func (bd BirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd.String())
}

// User is a regular user un our application.
type User struct {
	// Name is use for the concatenation firstname and lastname.
	Name string
	// BirthDate is the birthdate of the user.
	BirthDate BirthDate
}
