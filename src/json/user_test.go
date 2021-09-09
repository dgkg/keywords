package json_test

import (
	"crypto/sha256"
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
		Password:  "this-is-my-pass",
	}

	data, err := json.Marshal(&u)
	if err != nil {
		t.Errorf("error marshal %v", err)
	}

	fmt.Println(string(data))

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

	h := sha256.New()
	h.Write([]byte("this-is-my-pass"))
	refPass := fmt.Sprintf("%x", h.Sum(nil))

	h = sha256.New()
	h.Write([]byte("this-is-my-pass"))
	u2.Password = internalJSON.Password(fmt.Sprintf("%x", h.Sum(nil)))

	if internalJSON.Password(refPass) != u2.Password {
		t.Errorf("wrong password %v wait for %v", string(u2.Password), string(refPass))
	}

}

func TestBirthDate_string(t *testing.T) {
	var bd internalJSON.BirthDate = internalJSON.BirthDate(time.Date(2007, 1, 2, 0, 0, 0, 0, time.FixedZone("UTC+1", 1*60*60)))
	if bd.String() != "2007-01-02" {
		t.Errorf("wrong bithdate %v", bd)
	}
}
