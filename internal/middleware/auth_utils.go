package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotAuthenticated = errors.New("user not authenticated")
	ErrInvalidUserIdType    = errors.New("invalid user_id type")
	ErrInvalidUsernameType  = errors.New("invalid username type")
)

func GetIdFromContext(ctx *gin.Context) (int, error) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		return 0, ErrUserNotAuthenticated
	}
	id, ok := userId.(int)
	if !ok {
		return 0, ErrInvalidUserIdType
	}
	return id, nil
}

func MustGetUserId(ctx *gin.Context) int {
	userId, err := GetIdFromContext(ctx)
	if err != nil {
		panic("AuthMiddleware must be used")
	}
	return userId
}

func GetUsernameFromContext(ctx *gin.Context) (string, error) {
	username, exists := ctx.Get("username")
	if !exists {
		return "", ErrUserNotAuthenticated
	}
	name, ok := username.(string)
	if !ok {
		return "", ErrInvalidUsernameType
	}
	return name, nil
}
