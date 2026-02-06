package middleware

import (
	"net/http"
	"strings"

	"ToDoApi/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			ctx.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			ctx.Abort()
			return
		}
		tokenString := parts[1]
		claims, err := jwtService.ExtractClaims(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserId)
		ctx.Set("username", claims.Username)
		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
