package json_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	internalJSON "keywords/src/json"
)

func TestUserMarshalJSONUnmarshalJSON(t *testing.T) {
	date := time.Date(2007, 1, 2, 0, 0, 0, 0, time.FixedZone("UTC+1", 1*60*60))
	var u internalJSON.User = internalJSON.User{
		Name:      "Bob",
		BirthDate: internalJSON.BirthDate(date),
		Password:  []byte("this-is-my-pass"),
	}

	data, err := json.Marshal(&u)
	if err != nil {
		t.Errorf("error marshal %v", err)
	}

	var u2 internalJSON.User
	err = json.Unmarshal(data, &u2)
	if err != nil {
		t.Errorf("error unmarshal %v", err)
	}

	if u2.Name != "Bob" {
		t.Errorf("wrong name %v", u2.Name)
	}

	if !strings.Contains(u2.BirthDate.String(), "2007-01-02") {
		t.Errorf("wrong bithdate %v", u2.BirthDate)
	}

	if u2.BirthDate.String() != "2007-01-02" {
		t.Errorf("wrong bithdate %v", u2.BirthDate)
	}

	if time.Time(u2.BirthDate).UnixMilli() != date.UnixMilli() {
		t.Errorf("wrong bithdate %v wait for %v", u2.BirthDate, date)
	}

	refPass := base64.StdEncoding.EncodeToString([]byte("this-is-my-pass"))
	resRefPass, _ := internalJSON.HashPassword([]byte(refPass), []byte(`this-is-my-salt`))
	ok, err := internalJSON.Authenticate(u2.Password, []byte(refPass), resRefPass)
	fmt.Println("ok", ok)
	fmt.Println("err", err)

}

func TestBirthDate_string(t *testing.T) {
	var bd internalJSON.BirthDate = internalJSON.BirthDate(time.Date(2007, 1, 2, 0, 0, 0, 0, time.FixedZone("UTC+1", 1*60*60)))
	if bd.String() != "2007-01-02" {
		t.Errorf("wrong bithdate %v", bd)
	}
}
