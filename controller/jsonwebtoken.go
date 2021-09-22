package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "ji3g45/4u;6"
)

type Claims struct {
	ID string
	jwt.StandardClaims
}

func SetToken(userid string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	claims := Claims{
		ID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETKEY))
	return token, err
}

func AuthJWT(token string, userid string) bool {
	var auth bool
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		auth = false
		return auth
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			temp := claims.ID
			if temp == userid {
				auth = true
			}
			return auth
		}
	}

	return auth
}
