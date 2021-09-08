package json_test

import (
	"encoding/json"
	"testing"

	internalJSON "keywords/src/json"
)

func TestUserUnmarshalJSON(t *testing.T) {
	data := []byte(`{"name": "Rob", "birth_date": "2000-05-22"}`)
	var u internalJSON.User
	err := json.Unmarshal(data, &u)
	if err != nil {
		t.Errorf("error unmarshal %v", err)
	}
}

func TestUserMarshalJSON(t *testing.T) {

}
