package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "test_salt"
	signingKey = "sdflfasdq235f"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	GroupId  int    `json:"group_id"`
}

func ParseToken(authToken string) (int, error) {
	token, err := jwt.ParseWithClaims(authToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type tokenClaims")
	}
	return claims.GroupId, nil
}
