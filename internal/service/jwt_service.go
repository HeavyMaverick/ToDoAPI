package service

import (
	"ToDoApi/internal/model"
	"errors"
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
		return s.secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	if token.Method.Alg() != jwt.SigningMethodHS256.Name {
		return nil, errors.New("invalid signing method")
	}
	return token, nil
}

func (s *jwtService) ExtractClaims(tokenString string) (*jwt.Claims, error)
func (s *jwtService) RefreshToken(tokenString string) (string, error)
