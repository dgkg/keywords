package jwt

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

type CustomPayload struct {
	jwt.Payload
	AccessLevel string `json:"access_level,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	UUIDUser    string `json:"uuid_user,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

func New(uuid, level, userName string) (string, error) {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         "Atos",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://xxx.org", "https://zzz.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		UUIDUser:    uuid,
		UserName:    userName,
		AccessLevel: level,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func Valid(token string) error {
	var pl CustomPayload
	_, err := jwt.Verify([]byte(token), hs, &pl)
	if err != nil {
		return err
	}
	return nil
}
