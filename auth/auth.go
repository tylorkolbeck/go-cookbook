package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	SigningKey []byte
}

func NewAuthConfig(signingKey []byte) (*AuthConfig, error) {
	if len(signingKey) == 0 {
		return nil, errors.New("signing key cannot be empty")
	}

	return &AuthConfig{SigningKey: signingKey}, nil
}

func SaltPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (ac *AuthConfig) GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(ac.SigningKey)

	return tokenString, err
}
