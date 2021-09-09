package json

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02"

type User struct {
	Name      string
	BirthDate time.Time
}

func (u *User) UnmarshalJSON(b []byte) error {
	var aux userAux
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	if len(aux.BirthDate) == 0 {
		aux.BirthDate = "0000-00-00"
	}

	u.Name = aux.Name

	t, err := time.ParseInLocation(DateFormat, aux.BirthDate, time.FixedZone("UTC+2", 2*60*60))

	if err != nil {
		return err
	}
	u.BirthDate = t

	return nil
}

type userAux struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthdate"`
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := userAux{
		Name:      u.Name,
		BirthDate: u.BirthDate.Format(DateFormat),
	}
	return json.Marshal(aux)
}
