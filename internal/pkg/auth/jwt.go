package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("Micplus-tiktok")

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func ReleaseToken(id int64) (string, error) {
	expire := time.Now().Add(24 * time.Hour)
	claims := Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "tiktok",
			Subject:   "user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return ts, nil
}

func ParseToken(token string) (*Claims, bool) {
	t, _ := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, true
	}
	return nil, false
}
