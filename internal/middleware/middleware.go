package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Bearer token is required"})
			c.Abort()
			return
		}

		secret := cast.ToString(os.Getenv("JWT_SECRET"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Invalid token" + err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
		} else {
			c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()

	}
}
