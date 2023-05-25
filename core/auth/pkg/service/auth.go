package service

import (
	"auth-app-service/pkg/metrics"
	authApp "auth-app-service/pkg/model"
	"auth-app-service/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "test_salt"
	signingKey = "sdflfasdq235f"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo    repository.Authorization
	metrics metrics.MetricsCollector
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	GroupId  int    `json:"group_id"`
}

func NewAuthService(repo repository.Authorization, metrics *metrics.Metrics) *AuthService {
	return &AuthService{
		repo:    repo,
		metrics: metrics,
	}
}

func (s *AuthService) CreateUser(user authApp.User) (int, error) {
	user.Password = generateHash(user.Password)
	id, err := s.repo.CreateUser(user)
	if err == nil {
		s.metrics.RegisterUserHandler()
	}
	return id, err
}

func (s *AuthService) GetUser(username, password string) (authApp.Profile, error) {
	return s.repo.GetUser(username, password)
}

func (s *AuthService) GetUserById(id int) (authApp.Profile, error) {
	return s.repo.GetUserById(id)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
		user.Username,
		user.GroupID,
	})
	return token.SignedString([]byte(signingKey))
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(authToken string) (int, error) {
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
	return claims.UserId, nil
}
