package middlewares

import (
	"fmt"
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
		if len(jwtKey) == 0 {
			jwtKey = []byte(os.Getenv("SECRET_KEY"))
			if len(jwtKey) == 0 {
				jwtKey = []byte("tu_clave_secreta")
			}
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = tokenString[7:] 
		}
		tokenString = strings.TrimSpace(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("CLAIMS EXTRAÍDOS: ID=%v, User=%v\n", claims["user_id"], claims["username"])
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
					c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a este perfil"})
					c.Abort()
					return
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
