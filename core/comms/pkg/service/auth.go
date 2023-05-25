package service

import (
	"comms-app-service/pkg/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "sdflfasdq235f"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) ParseToken(authToken string) (int, error) {
	token, err := jwt.ParseWithClaims(authToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		// expired
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type tokenClaims")
	}
	return claims.UserId, nil
}
