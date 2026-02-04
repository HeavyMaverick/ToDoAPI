package service

import (
	"ToDoApi/internal/model"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(*model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractClaims(tokenString string) (*jwt.Claims, error)
	RefreshToken(tokenString string) (string, error)
}

type jwtService struct {
	secretKey     []byte
	tokenDuration time.Duration
	issuer        string
}

type JWTConfig struct {
	SecretKey     string        `json:"sercet_key"`
	TokenDuration time.Duration `json:"token_duration"`
	Issuer        string        `json:"issuer"`
}

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var (
	ErrUserNil = errors.New("user can't be nil")
)

func NewJwtService(config JWTConfig) (*jwtService, error) {
	if config.SecretKey == "" {
		panic("secretKey can't be empty")
	}
	if config.TokenDuration == 0 {
		config.TokenDuration = 24 * time.Hour
	}
	if config.Issuer == "" {
		config.Issuer = "todoapi"
	}
	return &jwtService{
		secretKey:     []byte(config.SecretKey),
		tokenDuration: config.TokenDuration,
		issuer:        config.Issuer,
	}, nil
}

func (s *jwtService) GenerateToken(user *model.User) (string, error) {
	if user == nil {
		return "", ErrUserNil
	}
	claims := Claims{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer,
			Subject:   user.Username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func (s *jwtService) ExtractClaims(tokenString string) (*Claims, error) {
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return nil, errors.New("invalid user_id in token")
		}
		return &Claims{
			UserId:   int(userIDFloat),
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
		}, nil
	}
	return nil, errors.New("invalid token claims")
}

func (s *jwtService) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})
	if err != nil {
		if !strings.Contains(err.Error(), "token is expired") {
			return "", err
		}
	}
	if claims, ok := token.Claims.(*Claims); ok {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.tokenDuration))
		claims.IssuedAt = jwt.NewNumericDate(time.Now())
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return newToken.SignedString(s.secretKey)
	}
	return "", errors.New("invalid token claims")
}
