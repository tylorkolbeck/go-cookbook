package service

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/google/uuid"
	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
	"golang.org/x/crypto/bcrypt"

	"github.com/tylorkolbeck/go-cookbook/auth"
)

type UserService struct {
	repo       repository.UserRepository
	authConfig auth.AuthConfig
}

func NewUserService(repo repository.UserRepository, authConfig auth.AuthConfig) *UserService {
	return &UserService{repo: repo, authConfig: authConfig}
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	if !ValidateEmail(user.Email) {
		return model.User{}, errors.New("Invalid email")
	}

	if !ValidatePassword(user.Password) {
		return model.User{}, errors.New("Invalid password: Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, and one number.")
	}

	_, err := s.repo.FindUserByEmail(user.Email)

	if err != nil {
		return model.User{}, errors.New("User with email already exists")
	}

	hashedPassword, err := auth.SaltPassword(user.Password)

	user.EmailVerified = false
	user.VerificationToken = uuid.New().String()

	if err != nil {
		return model.User{}, err
	}

	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func (s *UserService) Login(email string, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	// If email is not verified return an error
	if !user.EmailVerified {
		return "", errors.New("Email not verified")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	return s.authConfig.GenerateToken(user.ID.String())
}

func (s *UserService) VerifyEmail(token string) (bool, error) {
	user, err := s.repo.FindUserByVerificationToken(token)

	if err != nil {
		return false, err
	}

	user.EmailVerified = true
	user.VerificationToken = "" // Clear the token after verification

	return s.repo.SetUserEmailVerified(user.ID.String())
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(email string) (model.User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *UserService) ListUsers() ([]model.User, error) {
	return s.repo.ListUsers()
}

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	return regex.MatchString(email)
}

func ValidatePassword(password string) bool {
	var hasUpper, hasLower, hasNumber bool
	const minLength = 8

	if len(password) < minLength {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	// Modify the condition based on your password policy requirements
	return hasUpper && hasLower && hasNumber
}

func (s *UserService) DeleteUser(id string) (string, error) {
	return s.repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(id string, user dto.UpdateUserRequest) (model.User, error) {
	existingUser, err := s.repo.GetUserByID(id)

	if err != nil {
		return model.User{}, err
	}

	// Update fields based on the provided request. Check for non-nil before updating.
	if user.Email != nil {
		existingUser.Email = *user.Email
	}
	if user.Name != nil {
		existingUser.Name = *user.Name
	}

	if user.Name == nil && user.Email == nil {
		return model.User{}, errors.New("No fields to update")
	}

	return s.repo.UpdateUser(id, existingUser)
}
