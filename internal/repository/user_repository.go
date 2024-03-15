package repository

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.SafeUser, error)
	FindUserByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	FindUserByVerificationToken(token string) (model.User, error)
	SetUserEmailVerified(id string) (bool, error)
	ListUsers() ([]model.SafeUser, error)
	DeleteUser(id string) (string, error)
	UpdateUser(id string, user model.User) (model.SafeUser, error)
}
