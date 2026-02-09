package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

var ignoreRoutesOfVerification = []string{
	"/users/id/:id",
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtKey == nil {
			c.Next()
			return
			// jwtKey = []byte("tu_clave_secreta")
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, "Bearer ")
		if len(parts) > 1 {
			tokenString = parts[1]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			username := claims["username"].(string)
			c.Set("userId", userID)
			c.Set("username", username)

			// Ignore verification for some routes
			for _, route := range ignoreRoutesOfVerification {
				if c.FullPath() == route {
					paramId := c.Param("id")
					userIdStr := strconv.Itoa(int(userID))

					if paramId == userIdStr {
						c.Next()
						return
					}
				}
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
