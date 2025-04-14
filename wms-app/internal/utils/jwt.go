package utils

import (
	"time"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

func GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// In-memory blacklist
var tokenBlacklist = make(map[string]bool)

func IsTokenBlacklisted(token string) bool {
	return tokenBlacklist[token]
}

func BlacklistToken(token string) {
	tokenBlacklist[token] = true
}


func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization token missing or invalid"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		if IsTokenBlacklisted(tokenStr) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token has been logged out"})
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("user_id", claims["sub"])
		c.Set("email", claims["email"])
		c.Set("token_string", tokenStr)
		c.Next()
	}
}

