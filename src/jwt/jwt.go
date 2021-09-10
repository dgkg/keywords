package jwt

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

type CustomPayload struct {
	jwt.Payload
	AccessLevel int    `json:"access_level,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	UUIDUser    uint64 `json:"uuid_user,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

func New(level int, uuid uint64, userName string) (string, error) {
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

func valid(key, token string) error {
	var pl CustomPayload
	_, err := jwt.Verify([]byte(token), hs, &pl)
	if err != nil {
		return err
	}
	return nil
}

func MiddlewareJWT(key string, minLevelAccess int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authValue := ctx.GetHeader("Authorization")
		if len(authValue) == 0 {
			log.Println(errors.New("0 user not authorized"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Println("authValue", authValue)
		if !strings.Contains(authValue, "Baerer") {
			log.Println(errors.New("1 user not authorized"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authValue = strings.ReplaceAll(authValue, "Baerer ", "")
		err := valid(key, authValue)
		if err != nil {
			log.Println(errors.New("2 user not authorized"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var payload CustomPayload
		_, err = jwt.Verify([]byte(authValue), hs, &payload)
		if err != nil {
			log.Println(errors.New("3 user not authorized"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println("payload", payload)
		if payload.AccessLevel < minLevelAccess {
			log.Println(errors.New("user not authorized"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println("authorized !!!!")
	}
}
