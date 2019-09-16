package util

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JwtUtil struct{}

type Claims struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

var (
	jwtSecret = []byte("theSalt")
)

func (jwtUtil *JwtUtil) CreateToken(userId uint, username string) (string, error) {
	claims := Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 86400),
			Issuer:    "x-shares",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}

func (jwtUtil *JwtUtil) ParseToekn(token string) *Claims {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil
	}
	log.Println(tokenClaim)
	if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
		return claims
	}
	return nil
}
