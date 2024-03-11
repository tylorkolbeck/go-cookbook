package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthConfig struct {
	SigningKey []byte
}

func NewAuthConfig(signingKey []byte) *AuthConfig {
	return &AuthConfig{SigningKey: signingKey}
}

func (ac *AuthConfig) GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(ac.SigningKey)

	return tokenString, err
}
