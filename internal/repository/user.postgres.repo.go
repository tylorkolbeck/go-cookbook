package repository

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *PostgresUserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *PostgresUserRepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error

	return user, err
}

func (r *PostgresUserRepository) FindUserByVerificationToken(token string) (model.User, error) {
	var user model.User
	err := r.db.Where("verification_token = ?", token).First(&user).Error

	return user, err
}

func (r *PostgresUserRepository) SetUserEmailVerified(id string) (bool, error) {
	err := r.db.Model(&model.User{}).Where("id = ?", id).Update("email_verified", true).Error

	return err == nil, err
}

func (r *PostgresUserRepository) ListUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error

	safeUsers := make([]model.User, len(users))

	for i, user := range users {
		safeUsers[i] = user
	}

	return safeUsers, err
}

func (r *PostgresUserRepository) DeleteUser(id string) (string, error) {
	err := r.db.Delete(&model.User{}, id).Error

	return id, err
}

func (r *PostgresUserRepository) UpdateUser(id string, user model.User) (model.User, error) {
	var existingUser model.User
	err := r.db.First(&existingUser, id).Error

	if err != nil {
		return model.User{}, err
	}

	err = r.db.Model(&existingUser).Updates(user).Error

	return existingUser, err
}
