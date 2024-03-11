package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

var (
	ErrUserNotFound = errors.New("User not found")
	ErrorUserExists = errors.New("User already exists")
	ErrorUserUpdate = errors.New("User update failed")
	ErrorUserDelete = errors.New("User delete failed")
	ErrorUserCreate = errors.New("User create failed")
)

type InMemoryUserRepository struct {
	users map[string]model.User
	mutex sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]model.User),
	}
}

func (r *InMemoryUserRepository) FindUserByVerificationToken(token string) (model.User, error) {
	for _, user := range r.users {
		if user.VerificationToken == token {
			return user, nil
		}
	}

	return model.User{}, ErrUserNotFound
}

func (r *InMemoryUserRepository) ListUsers() ([]model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepository) SetUserEmailVerified(id string) (bool, error) {
	user, ok := r.users[id]

	if !ok {
		return false, ErrUserNotFound
	}

	user.EmailVerified = true
	r.users[id] = user
	return true, nil
}

func (r *InMemoryUserRepository) CreateUser(user model.User) (model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return model.User{}, ErrorUserExists
	}

	user.ID = uuid.New().String()

	r.users[user.ID] = user
	return user, nil
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return model.User{}, ErrUserNotFound
}

func (r *InMemoryUserRepository) FindUserByID(id string) (model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepository) UpdateUser(id string, user model.User) (model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return model.User{}, ErrUserNotFound
	}

	r.users[id] = user
	return user, nil
}

func (r *InMemoryUserRepository) DeleteUser(id string) (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return "", ErrUserNotFound
	}

	delete(r.users, id)
	return id, nil
}

func (r *InMemoryUserRepository) Get() []model.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *InMemoryUserRepository) GetByID(id string) (model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}
