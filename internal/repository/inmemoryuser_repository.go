package repository

import (
	"errors"
	"sync"
	"time"

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
		users: map[string]model.User{
			"1e8d1d67-f61a-436b-8f43-4e0a094a5568": {
				ID:            "1e8d1d67-f61a-436b-8f43-4e0a094a5568",
				Email:         "test@test.com",
				Password:      "", // Assuming Password is required but omitted here
				Name:          "Tylor",
				Created_at:    time.Now(),
				Updated_at:    time.Now(),
				EmailVerified: false,
				Role:          "admin",
				// VerificationToken field is omitted; add if necessary
			},
		},
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

func (r *InMemoryUserRepository) ListUsers() ([]model.SafeUser, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]model.SafeUser, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, model.SafeUser{
			ID:            user.ID,
			Email:         user.Email,
			Name:          user.Name,
			Created_at:    user.Created_at.String(),
			Updated_at:    user.Updated_at.String(),
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
		})
	}

	return users, nil
}

func (r *InMemoryUserRepository) GetUserByID(id string) (model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
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

func (r *InMemoryUserRepository) CreateUser(user model.User) (model.SafeUser, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return model.SafeUser{}, ErrorUserExists
	}

	user.ID = uuid.New().String()

	// Turn user into safe user
	safeUser := model.SafeUser{
		ID:         user.ID,
		Email:      user.Email,
		Created_at: user.Created_at.String(),
		Updated_at: user.Updated_at.String(),
	}

	r.users[user.ID] = user
	return safeUser, nil
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

func (r *InMemoryUserRepository) DeleteUser(id string) (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return id, ErrUserNotFound
	}

	delete(r.users, id)
	return id, nil
}

func (r *InMemoryUserRepository) UpdateUser(id string, user model.User) (model.SafeUser, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return model.SafeUser{}, ErrorUserUpdate
	}

	println(user.Name)
	println(user.Email)

	r.users[id] = user

	// Turn user into safe user
	safeUser := model.SafeUser{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Created_at: user.Created_at.String(),
		Updated_at: user.Updated_at.String(),
	}

	return safeUser, nil
}
