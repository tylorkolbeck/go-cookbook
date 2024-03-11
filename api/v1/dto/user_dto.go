package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserBase struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	UserBase
}

type UpdateUserRequest struct {
	UserBase
}

type LoginRequest struct {
	UserBase
}

func (ub *UserBase) Validate() error {
	validate := validator.New()
	return validate.Struct(ub)
}

func (r CreateUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r UpdateUserRequest) Validate() error {
	validate := validator.New()
	// Custom validation logic for UpdateUserRequest
	return validate.Struct(r)
}
