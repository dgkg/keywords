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
	aux := struct {
		Name      string `json:"name"`
		BirthDate string `json:"birthdate"`
	}{}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	u.Name = aux.Name

	t, err := time.Parse(DateFormat, aux.BirthDate)
	if err != nil {
		return err
	}
	u.BirthDate = t

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		Name      string `json:"name"`
		BirthDate string `json:"birthdate"`
	}{

		Name:      u.Name,
		BirthDate: u.BirthDate.Format(DateFormat),
	}
	return json.Marshal(aux)
}
