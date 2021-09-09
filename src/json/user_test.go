package json_test

import (
	"encoding/json"
	"testing"
	"time"

	internalJSON "keywords/src/json"
)

func TestUserMarshalJSONUnmarshalJSON(t *testing.T) {

	var u internalJSON.User = internalJSON.User{
		Name:      "Bob",
		BirthDate: time.Date(2007, 1, 2, 0, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60)),
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

	if u2.BirthDate.String() != "2007-01-02" {
		t.Errorf("wrong name %v", u2.Name)
	}
}

func TestUserMarshalJSON(t *testing.T) {

}
