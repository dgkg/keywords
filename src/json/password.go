package json

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Password string

func (p *Password) UnmarshalJSON(b []byte) error {
	if p == nil {
		return nil
	}

	var aux string
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(aux))

	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))

	return nil
}

// MarshalJSON is implementing the encoding/json interface Marshaler.
func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}
