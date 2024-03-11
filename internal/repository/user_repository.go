package repository

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.SafeUser, error)
	FindUserByEmail(email string) (model.SafeUser, error)
	FindUserByVerificationToken(token string) (model.SafeUser, error)
	SetUserEmailVerified(id string) (bool, error)
	ListUsers() ([]model.SafeUser, error)
	// FindUserByID(id string) (model.User, error)
	// UpdateUser(id string, user model.User) (model.User, error)
	// DeleteUser(id string) (string, error)
}
