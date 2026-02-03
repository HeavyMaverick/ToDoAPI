package service

import (
	"ToDoApi/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(*model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractClaims(tokenString string) (*jwt.Claims, error)
	RefreshToken(tokenString string) (string, error)
}

type jwtService struct {
	secretKey []byte
}

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func NewJwtService(secretKey string) *jwtService {
	if secretKey == "" {
		panic("secretKey can't be empty")
	}
	return &jwtService{
		secretKey: []byte(secretKey),
	}
}

func (s *jwtService) GenerateToken(user *model.User) (string, error) {
	if user == nil {
		return "", nil
	}
	return "", nil
}
func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error)
func (s *jwtService) ExtractClaims(tokenString string) (*jwt.Claims, error)
func (s *jwtService) RefreshToken(tokenString string) (string, error)
