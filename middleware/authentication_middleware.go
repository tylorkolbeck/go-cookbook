package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("secret")

// AuthMiddleware checks for the presence and validity of the JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		tokenParts := strings.Split(tokenString, " ")
		if (len(tokenParts) != 2) || (tokenParts[0] != "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		tokenString = tokenParts[1]
		println(tokenString)

		token, err := jwt.Parse(strings.Trim(tokenString, " "), func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				println("SOME OTHER ISSUE")
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return mySigningKey, nil
		})

		if err != nil {
			println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			println(claims["user_id"])
			// Set the claims on the context so they can be retrieved in the handler if needed
			c.Set("userID", claims["user_id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
